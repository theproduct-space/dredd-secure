import React, { useState } from 'react';
import './TableView.css';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import { MsgCancelEscrowResponse } from 'dredd-secure-client-ts/dreddsecure.escrow/types/dreddsecure/escrow/tx';

export interface TableHeader {
    label: string;
    dataProp: string;
}

export interface TableData {
    id: string;
    deadline: string;
    status: string;
    assetsInvolved: string;
}

export interface TableViewProps {
    headers: TableHeader[];
    data: TableData[];
    filterOptions: {
        prop: string;
        value: string | undefined;
    }[];
}

export default function TableView(props: TableViewProps) {
    const navigate = useNavigate();
    const { data, headers, filterOptions } = props;

    const [sortKey, setSortKey] = useState(headers[0].dataProp);
    const [sortAscending, setSortAscending] = useState(false);

    const handleSortingChange = (param: string) => {
        if (sortKey == param) {
            setSortAscending(!sortAscending);
        }
        else {
            setSortKey(param);
            setSortAscending(false);
        }
    }

    const sortedData = [...data].sort((a, b) => {
        const valueA = a[sortKey];
        const valueB = b[sortKey];

        if (valueA < valueB) return sortAscending ? -1 : 1;
        if (valueA > valueB) return sortAscending ? 1 : -1;
        return 0;
    });

    const handleOnClickRow = (id: string) => {
        navigate(`/escrow/${id}`);
    }

    return (
        <>
            <div className="table">
                <div className="table-row header-row">
                    {
                        headers.map((header) => {
                            return (
                                <div key={"header-" + header.label} className="table-cell header-cell" onClick={() => handleSortingChange(header.dataProp)}>
                                    {header.label}
                                    {
                                        /* TODO (Design): Replace A & D by an actual icon */
                                        sortKey === header.dataProp && sortAscending ?
                                            <span className="icon-ascending"><b>A</b></span>
                                            :
                                            sortKey === header.dataProp && !sortAscending ?
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
                    sortedData.map((element, index) => {
                        for (const filter of filterOptions) {
                            if (filter.value != "" && element[filter.prop] != filter.value) return;
                        }

                        return (
                            <div key={`data-${index}`} className="table-row" onClick={() => handleOnClickRow(element.id)}>
                                {
                                    headers.map((header) => {
                                        return <div key={`data-${index}-${header.dataProp}`} className="table-cell">{element[header.dataProp]}</div>
                                    })
                                }
                            </div>
                        )
                    })}
            </div>
        </>
    )
}