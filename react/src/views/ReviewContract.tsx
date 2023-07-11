import { queryClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";
import React, { useEffect, useState } from "react";
import { useLocation, useParams } from "react-router-dom";
import { IContract } from "~sections/CreateContract";
import ReviewContractSection from "~sections/ReviewContractSection";

function ReviewContract() {
  const { id } = useParams<{ id: string }>();
  const [contract, setContract] = useState<EscrowEscrow>();

  useEffect(() => {
    const fetchEscrow = async () => {
      try {
        const response = await queryClient().queryEscrow(id ?? "");
        setContract(response.data.Escrow);
      } catch (error) {
        console.error(error);
      }
    };

    fetchEscrow();
  });

  return (
    <div>
      <div className="messages">Success message here</div>
      <div className="title">Review escrow Contract #{contract?.id}</div>
      <ReviewContractSection contract={contract} />
    </div>
  );
}

export default ReviewContract;
