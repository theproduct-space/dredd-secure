import React from 'react';
import './TableView.css';

export interface TableHeader {
    label: string;
    dataProp: string;
}

export interface TableData {
    [key: string]: any;
}

export interface TableViewProps {
    headers: TableHeader[];
    data: TableData[];
    sortingFunction: (param: string, orderAsc: boolean) => void;
}

export default function TableView(props: TableViewProps) {
    const { headers, data, sortingFunction } = props;

    let sortedBy = {
        param: headers[0].dataProp,
        orderedAsc: false
    }

    const handleSortingChange = (param: string) => {
        // TODO: Handle sorting function call here.
        console.log("Change sorting to '%s'", param);
    }

    return (
        <>
            <div className="table">
                <div className="table-row header-row">
                    {
                        headers.map((header) => {
                            return (
                                <div className="table-cell header-cell" onClick={() => handleSortingChange(header.dataProp)}>
                                    {header.label}
                                    {
                                        /* TODO (Design): Replace A & D by an actual icon */
                                        sortedBy.param === header.dataProp && sortedBy.orderedAsc ?
                                            <span className="icon-ascending"><b>A</b></span>
                                            :
                                            sortedBy.param === header.dataProp && !sortedBy.orderedAsc ?
                                                <span className="icon-descending"><b>D</b></span>
                                                :
                                                null
                                    }
                                </div>
                            )
                        })
                    }
                </div>
                {
                    data.map((element) => {
                        return (
                            <div className="table-row">
                                {
                                    headers.map((header) => {
                                        return <div className="table-cell">{element[header.dataProp]}</div>
                                    })
                                }
                            </div>
                        )
                    })}
            </div>
        </>
    )
}