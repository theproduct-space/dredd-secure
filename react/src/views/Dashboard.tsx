// React Imports
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

// dredd-secure-client-ts Imports
import { queryClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";
// Hooks Imports

// Custom Imports
import FilterDropDown, {
  FilterDropDownProps,
} from "~baseComponents/FilterDropDown";
import TableView, { TableData } from "~baseComponents/TableView";
import Account from "~sections/Account";
import useWallet from "../components/utils/useWallet";

const Dashboard = () => {
  const [selectedStatus, setSelectedStatus] = useState<string>("");
  const [onlyOwnDisplayed, setOnlyOwnDisplayed] = useState(false);
  const [escrows, setEscrows] = useState<EscrowEscrow[]>([]);
  // TODO: Get address from keplr or other wallet manager
  const { address } = useWallet();

  const tableHeaders = [
    {
      label: "contract-id",
      dataProp: "id",
    },
    {
      label: "assets involved",
      dataProp: "assetsInvolved",
    },
    {
      label: "status",
      dataProp: "status",
    },
    {
      label: "deadline",
      dataProp: "deadline",
    },
  ];

  const statusFilterFunction = (status: string) => {
    setSelectedStatus(status);
  };
  const statusFilterDropDownProps: FilterDropDownProps = {
    choices: [
      {
        label: "All",
        filterValue: "",
      },
      {
        label: "Open",
        filterValue: "open",
      },
      {
        label: "Pending",
        filterValue: "pending",
      },
      {
        label: "Closed",
        filterValue: "closed",
      },
    ],
    filterFunction: statusFilterFunction,
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
        deadline: escrow.endDate ?? "",
        assetsInvolved: `${creatorCoin} <-> ${fulfillerCoin}`,
        status: escrow.status ?? "",
        initiator: escrow.initiator ?? "",
      };
    })
    .filter(Boolean) as TableData[];

  useEffect(() => {
    // Function to fetch and update data array
    const fetchData = async () => {
      const escrows = (await queryClient().queryEscrowAll()).data.Escrow ?? [];
      setEscrows(escrows);

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
      }
    };
    // Fetch data initially
    fetchData();

    // Fetch data every 2 seconds
    const interval = setInterval(fetchData, 2000);

    // Clean up the interval when the component is unmounted
    return () => clearInterval(interval);
  }, []); // Empty dependency array to run the effect only once on component mount

  return (
    <>
      <div>
        <div className="title text-3xl">All contracts</div>
        <div className="subtitle">
          You can view the status of all contracts that you have sent and/or
          received from here
        </div>
      </div>
      <div>
        <Link to={"/escrow/create"}>Create Contract</Link>
      </div>
      <hr></hr>{" "}
      {/* Just to let me see a clear separation between the table and other thing */}
      <div>
        <div>
          <div className="font-bold">Filter by status</div>
          <FilterDropDown {...statusFilterDropDownProps} />
        </div>
        <div>
          <div className="font-bold">Show my relevant contracts only</div>
          <input
            checked={onlyOwnDisplayed ?? undefined}
            type="checkbox"
            onChange={(e) => setOnlyOwnDisplayed(e.target.checked)}
          ></input>
        </div>
        <TableView
          headers={tableHeaders}
          data={formatedData}
          filterOptions={[{ prop: "status", value: selectedStatus }]}
        />
      </div>
      <Account />
    </>
  );
};

export default Dashboard;
