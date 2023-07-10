import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { FilterDropDown, FilterDropDownProps } from "../components/base/DropDown/FilterDropDown";
import TableView, { TableData, TableHeader, TableViewProps } from "../components/base/TableView/TableView";
import { queryClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";

function Dashboard() {
    const [selectedStatus, setSelectedStatus] = useState<string>("");
    const [onlyOwnDisplayed, setOnlyOwnDisplayed] = useState(false);
    const [escrows, setEscrows] = useState<EscrowEscrow[]>([]);
    const [address, setAddress] = useState("c"); // For testing purposes only, // TODO: Get address from keplr or other wallet manager
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
                filterValue: ""
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
    const formatedData: TableData[] = ([...escrows].map((escrow) => {
        const creatorCoin = escrow?.initiatorCoins?.[0]?.denom ?? "";
        const fulfillerCoin = escrow?.fulfillerCoins?.[0]?.denom ?? "";

        if (onlyOwnDisplayed && escrow.fulfiller != address && escrow.initiator != address)
            return ;

        return {
            id: escrow.id ?? "",
            deadline: escrow.endDate ?? "",
            assetsInvolved: `${creatorCoin} <-> ${fulfillerCoin}`,
            status: escrow.status ?? ""
        };
    })).filter(Boolean) as TableData[];

    useEffect(() => {
        // Function to fetch and update data array
        const fetchData = async () => {
            let escrows = (await queryClient().queryEscrowAll()).data.Escrow ?? [];

            setEscrows(escrows);

            // For testing purposes only
            // TODO: Remove this when Xzan has finished with the design
            if (escrows.length == 0) {
                setEscrows([
                    {
                        id: "012345",
                        initiatorCoins: [{denom: "ETH", amount: "2"}],
                        fulfillerCoins: [{denom: "ATOM", amount: "2"}],
                        status: "open",
                        startDate: "1689006043",
                        endDate: "1789006043",
                    },
                    {
                        id: "123456",
                        initiatorCoins: [{denom: "ATOM", amount: "2"}],
                        fulfillerCoins: [{denom: "TOKEN", amount: "2"}],
                        status: "open",
                        startDate: "1689001053",
                        endDate: "1779006043",
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
                    You can view the status of all contracts that you have sent and/or received from here
                </div>
            </div>
            <div>
                <Link to={"#"}>Create Contract</Link>
            </div>
            <hr></hr> {/* Just to let me see a clear separation between the table and other thing */}
            <div>
                <div>
                    <div className="font-bold">Filter by status</div>
                    <FilterDropDown {...statusFilterDropDownProps} />
                </div>
                <div>
                    <div className="font-bold">Show my relevant contracts only</div>
                    <input checked={onlyOwnDisplayed ?? undefined} type="checkbox" onChange={e => setOnlyOwnDisplayed(e.target.checked)}></input>
                </div>
                <TableView  headers={tableHeaders} 
                            data={formatedData} 
                            filterOptions={[{ prop: "status", value: selectedStatus }]}
                            />
            </div>
        </>
    );
}

export default Dashboard;
