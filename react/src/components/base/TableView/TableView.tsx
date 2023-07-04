import React from 'react';
import styles from './TableView.module.scss';

export function TableView() {
    const escrows = [
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
    ]

    return (
        <div className="table">
            <div className="table-row header-row">
                <div className="table-cell header-cell">contract id#</div>
                <div className="table-cell header-cell">assets involved</div>
                <div className="table-cell header-cell">status</div>
                <div className="table-cell header-cell">deadline</div>
            </div>
            {escrows.map((escrow) => {
                return (
                    <div className="table-row">
                        <div className="table-cell">{escrow.id}</div>
                        <div className="table-cell">{escrow.assetsInvolved}</div>
                        <div className="table-cell">{escrow.status}</div>
                        <div className="table-cell">{escrow.deadline}</div>
                    </div>
                )
            })}
        </div>
    )
}