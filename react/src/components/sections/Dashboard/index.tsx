// React Imports
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

// dredd-secure-client-ts Imports
import { queryClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";
// Hooks Imports

// Custom Imports
import Card from "~baseComponents/Card";
import Dropdown, { DropdownProps } from "~baseComponents/Dropdown";
import TableView, { TableData } from "~baseComponents/TableView";
import Typography from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import useWallet from "../../utils/useWallet";

// Assets Imports
import bgImage from "~assets/3d-logoNoBg.webp";
import Button from "~baseComponents/Button";

const Dashboard = () => {
  const [selectedFilter, setSelectedFilter] = useState({
    label: "All",
    value: "",
  });
  const [onlyOwnDisplayed, setOnlyOwnDisplayed] = useState(false);
  const [escrows, setEscrows] = useState<EscrowEscrow[]>([]);
  // TODO: Get address from keplr or other wallet manager
  const { address } = useWallet();

  const formatDate = (timestamp: string): string => {
    const date = new Date(Number(timestamp) * 1000);
    return `${String(date.getMonth() + 1).padStart(2, "0")}/${String(
      date.getDate(),
    ).padStart(2, "0")}/${date.getFullYear()}`;
  };

  const tableHeaders = [
    {
      label: "contract id#",
      dataProp: "id",
      minWidth: "150px",
    },
    {
      label: "assets involved",
      dataProp: "assetsInvolved",
      minWidth: "180px",
    },
    {
      label: "status",
      dataProp: "status",
      minWidth: "100px",
    },
    {
      label: "deadline",
      dataProp: "deadline",
      minWidth: "100px",
    },
  ];

  const filterDropdownProps: DropdownProps = {
    choices: [
      {
        label: "All",
        value: "",
      },
      {
        label: "Open",
        value: "open",
      },
      {
        label: "Pending",
        value: "pending",
      },
      {
        label: "Closed",
        value: "closed",
      },
    ],
    selectedOption: selectedFilter,
    setSelectedOption: setSelectedFilter,
  };

  const formatedData: TableData[] = [...escrows]
    .map((escrow) => {
      const creatorCoin = escrow?.initiatorCoins?.[0]?.denom ?? "";
      const fulfillerCoin = escrow?.fulfillerCoins?.[0]?.denom ?? "";

      if (
        onlyOwnDisplayed &&
        escrow.fulfiller != address &&
        escrow.initiator != address
      )
        return;

      return {
        id: escrow.id ?? -1,
        deadline: formatDate(escrow.endDate as string) ?? "",
        assetsInvolved: `${creatorCoin} <-> ${fulfillerCoin}`,
        status: escrow.status ?? "",
        initiator: escrow.initiator ?? "",
        fulfiller: escrow.fulfiller ?? "",
      };
    })
    .filter(Boolean) as TableData[];

  useEffect(() => {
    // Function to fetch and update data array
    const fetchData = async () => {
      const escrows = (await queryClient().queryEscrowAll()).data.Escrow ?? [];

      // For testing purposes only
      // TODO: Remove this when Xzan has finished with the design
      if (escrows.length == 0) {
        setEscrows([
          {
            id: "012345",
            initiatorCoins: [{ denom: "ETH", amount: "2" }],
            fulfillerCoins: [{ denom: "ATOM", amount: "2" }],
            status: "open",
            startDate: "1689006043",
            endDate: "1789006043",
            initiator: "cosmos19r8syyde5naz8ysq0ul87qv3s56zgxw0829hag",
          },
          {
            id: "123456",
            initiatorCoins: [{ denom: "ATOM", amount: "2" }],
            fulfillerCoins: [{ denom: "TOKEN", amount: "2" }],
            status: "open",
            startDate: "1689001053",
            endDate: "1779006043",
            initiator: "cosmos1rljg6ldskneppq0j39mngv57avsvnjxjlw8z2q",
          },
        ]);
      } else setEscrows(escrows);
    };
    // Fetch data initially
    fetchData();

    // Fetch data every 2 seconds
    const interval = setInterval(fetchData, 2000);

    // Clean up the interval when the component is unmounted
    return () => clearInterval(interval);
  }, []); // Empty dependency array to run the effect only once on component mount

  return (
    <ContentContainer>
      <div className="relative min-h-screen w-full pt-32 max-w-4xl mx-auto">
        <img
          src={bgImage}
          alt="Dredd-Secure"
          className="object-cover absolute -right-20 bottom-0 z-0 drop-shadow-lightOrangeWide opacity-80 max-w-[500px]"
        />
        <div className="flex justify-between items-center gap-3 mb-3">
          <div>
            <Typography variant="h5">All Contracts</Typography>
            <Typography variant="small">
              You can view the status of all contracts that you have sent and/or
              received from here
            </Typography>
          </div>
          <div>
            {address && address !== "" && (
              <Link to={"/escrow/create"}>
                <Button text="Create Contract" className="capitalize" />
              </Link>
            )}
          </div>
        </div>

        <Card progress={"90"}>
          {/* // TODO REMOVE MIN-H */}
          <div className="p-4 md:p-10 min-h-[500px] ">
            <div className="flex justify-between mb-5">
              <div className="flex flex-col gap-2 max-w-[100px]">
                <Typography variant="small" className="font-revalia">
                  Filter
                </Typography>
                <Dropdown {...filterDropdownProps} />
              </div>

              {address != "" && (
                <div className="flex items-center gap-2">
                  <input
                    id="involved-escrow"
                    checked={onlyOwnDisplayed ?? undefined}
                    type="checkbox"
                    onChange={(e) => setOnlyOwnDisplayed(e.target.checked)}
                    className="mt-[3px]"
                  />
                  <label htmlFor="involved-escrow">
                    <Typography variant="small" className="font-revalia">
                      My escrows
                    </Typography>
                  </label>
                </div>
              )}
            </div>

            <div className="relative">
              <div className="absolute left-0 top-10 w-full border-b border-orange scale-150" />
              <TableView
                headers={tableHeaders}
                data={formatedData}
                filterOptions={[
                  { prop: "status", value: selectedFilter.value },
                ]}
              />
            </div>
          </div>
        </Card>
      </div>
    </ContentContainer>
  );
};

export default Dashboard;
