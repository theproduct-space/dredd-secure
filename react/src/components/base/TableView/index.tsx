// React Imports
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

// dredd-secure-client-ts Imports
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import {
  OfflineAminoSigner,
  OfflineDirectSigner,
} from "dredd-secure-client-ts/node_modules/@keplr-wallet/types";

// Styles Imports
import "./TableView.css";

export interface IWallet {
  address: string;
  offlineSigner: OfflineAminoSigner & OfflineDirectSigner;
}

export interface TableHeader {
  label: string;
  dataProp: string;
}

export interface TableData {
  id: number;
  deadline: string;
  status: string;
  assetsInvolved: string;
  initiator: string;
}

export interface TableViewProps {
  headers: TableHeader[];
  data: TableData[];
  wallet: IWallet;
  filterOptions: {
    prop: string;
    value: string | undefined;
  }[];
}

const TableView = (props: TableViewProps) => {
  const navigate = useNavigate();
  const { data, headers, filterOptions, wallet } = props;

  const [sortKey, setSortKey] = useState(headers[0].dataProp);
  const [sortAscending, setSortAscending] = useState(false);

  const messageClient = txClient({
    signer: wallet.offlineSigner,
    prefix: "cosmos",
    addr: "http://localhost:26657",
  });

  const handleSortingChange = (param: string) => {
    if (sortKey == param) {
      setSortAscending(!sortAscending);
    } else {
      setSortKey(param);
      setSortAscending(false);
    }
  };

  const sortedData = [...data].sort((a, b) => {
    const valueA = a[sortKey];
    const valueB = b[sortKey];

    if (valueA < valueB) return sortAscending ? -1 : 1;
    if (valueA > valueB) return sortAscending ? 1 : -1;
    return 0;
  });

  const handleCancelEscrow = (id: number) => {
    // Creator here is for testing only.
    // With a wallet connector, we will put the offline signer into the txClient above.
    console.log(id);
    messageClient.sendMsgCancelEscrow({
      value: { creator: wallet.address, id: id },
    });
    console.log("clicked");
  };

  const handleOnClickRow = (id: number) => {
    navigate(`/escrow/${id}`);
  };
  console.log(sortedData);
  console.log(wallet);
  return (
    <div className="table">
      <div className="table-row header-row">
        {headers.map((header) => {
          return (
            <button
              key={"header-" + header.label}
              className="table-cell header-cell"
              onClick={() => handleSortingChange(header.dataProp)}
            >
              {header.label}
              {
                /* TODO (Design): Replace A & D by an actual icon */
                sortKey === header.dataProp && sortAscending ? (
                  <span className="icon-ascending">
                    <b>A</b>
                  </span>
                ) : sortKey === header.dataProp && !sortAscending ? (
                  <span className="icon-descending">
                    <b>D</b>
                  </span>
                ) : null
              }
            </button>
          );
        })}
      </div>
      {sortedData.map((element, index) => {
        for (const filter of filterOptions) {
          if (filter.value != "" && element[filter.prop] != filter.value)
            return;
        }

        return (
          <React.Fragment key={`data-${index}`}>
            <button
              className="table-row"
              onClick={() => handleOnClickRow(element.id)}
            >
              {headers.map((header) => {
                return (
                  <div
                    key={`data-${index}-${header.dataProp}`}
                    className="table-cell"
                  >
                    {element[header.dataProp]}
                  </div>
                );
              })}
            </button>
            {element.initiator === wallet.address && (
              <span key={`initiator-${index}`} className="table-cell">
                <button onClick={() => handleCancelEscrow(element.id)}>
                  Cancel
                </button>
              </span>
            )}
          </React.Fragment>
        );
      })}
    </div>
  );
};

export default TableView;
