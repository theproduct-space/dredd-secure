import React from 'react'
import { Link } from 'react-router-dom';
import { FilterDropDown, FilterDropDownProps } from '~baseComponents/DropDown/FilterDropDown';
import TableView, { TableData, TableHeader, TableViewProps } from '~baseComponents/TableView/TableView'

function Dashboard() {
    const tableSortingFunction = (param: string, orderAsc: boolean) => {
        // TODO: Implement sorting logic here
    }
    const tableProps: TableViewProps = {
        headers: [
            {
                label: "contract-id",
                dataProp: "id"
            },
            {
                label: "assets involved",
                dataProp: "assetsInvolved"
            },
            {
                label: "status",
                dataProp: "status"
            },
            {
                label: "deadline",
                dataProp: "deadline"
            },
        ],
        data: [
            {
                id: "012345",
                assetsInvolved: "ETH",
                status: "Open",
                deadline: "12/04/23"
            },
            {
                id: "123456",
                assetsInvolved: "ATOM",
                status: "Closed",
                deadline: "13/06/23"
            },
        ],
        sortingFunction: tableSortingFunction
    }

    const statusFilterFunction = (status: string) => {
        // TODO: Implement status filtering logic here.
    }
    const statusFilterDropDownProps: FilterDropDownProps = {
        choices: [
            {
                label: "Open",
                filterValue: "Open"
            },
            {
                label: "Pending",
                filterValue: "Pending"
            },
            {
                label: "Closed",
                filterValue: "Closed"
            },
        ],
        filterFunction: statusFilterFunction,
    }

    return (
        <>
            <div>
                <div className="title text-3xl">All contracts</div>
                <div className="subtitle">You can view the status of all contracts that you have sent and/or received from here</div>
            </div>
            <div>
                <Link to={"#"}>Create Contract</Link>
            </div>
            <hr></hr> {/* Just to let me see a clear separation between the table and other thing */}
            <div>
                <div>
                    <div className="font-bold">Filter</div>
                    <FilterDropDown {...statusFilterDropDownProps} />
                </div>
                <TableView {...tableProps} />
            </div>
        </>
    )
}

export default Dashboard