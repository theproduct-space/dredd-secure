// React Imports
import React, { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

// dredd-secure-client-ts Imports
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";

// Hooks Imports
import useWallet from "../../utils/useWallet";

// Custom Components Imports
import ChevronDownIcon from "~icons/ChevronDownIcon";
import Typography from "~baseComponents/Typography";
import StatusPill from "~baseComponents/StatusPill";
import Garbage from "~icons/Garbage";

export interface TableHeader {
  label: string;
  dataProp: string;
  minWidth: string;
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
  filterOptions: {
    prop: string;
    value: string | undefined;
  }[];
}

const TableView = (props: TableViewProps) => {
  const chainId = "dreddsecure";
  const navigate = useNavigate();

  const { data, headers, filterOptions } = props;
  const { address, offlineSigner } = useWallet();
  const [sortKey, setSortKey] = useState(headers[0].dataProp);
  const [sortAscending, setSortAscending] = useState(false);

  const messageClient = txClient({
    signer: offlineSigner,
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

    messageClient.sendMsgCancelEscrow({
      value: { creator: address, id: id },
    });
  };

  const handleOnClickRow = (id: number) => {
    navigate(`/escrow/${id}`);
  };

  return (
    <div className="overflow-x-auto pb-4">
      <table className="table-auto w-full text-white-1000 border-separate border-spacing-y-2">
        <thead className="text-left">
          <tr>
            {headers.map((header) => (
              <th key={"header-" + header.label}>
                <button
                  className="flex items-center gap-1 mb-4"
                  onClick={() => handleSortingChange(header.dataProp)}
                >
                  <Typography
                    variant={"small"}
                    className="text-white-700 font-light"
                  >
                    {header.label.toUpperCase()}
                  </Typography>
                  {
                    /* TODO (Design): Replace A & D by an actual icon */
                    sortKey === header.dataProp && sortAscending ? (
                      <ChevronDownIcon className="rotate-180" />
                    ) : sortKey === header.dataProp && !sortAscending ? (
                      <ChevronDownIcon />
                    ) : null
                  }
                </button>
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {sortedData.map((element, index) => {
            for (const filter of filterOptions) {
              if (filter.value != "" && element[filter.prop] != filter.value)
                return;
            }

            return (
              <tr
                key={`data-${index}`}
                onClick={() => handleOnClickRow(element.id)}
              >
                {headers.map((header) => (
                  <td
                    key={`data-${index}-${header.dataProp}`}
                    style={{ minWidth: header.minWidth }}
                  >
                    {header.dataProp === "status" ? (
                      <StatusPill status={element[header.dataProp]} />
                    ) : (
                      element[header.dataProp]
                    )}
                  </td>
                ))}
                {/* // TODO TEST THIS*/}
                {element.initiator ===
                  "cosmos1peqnw9w3tsana5ycwcgp88mvnfqyl0e3clpym4" && (
                  <td key={`initiator-${index}`}>
                    <button
                      onClick={(e) => {
                        e.stopPropagation();
                        handleCancelEscrow(element.id);
                      }}
                      className={"flex"}
                    >
                      <Garbage height={20} width={20} />
                    </button>
                  </td>
                )}
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default TableView;
