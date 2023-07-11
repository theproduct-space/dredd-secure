// React Imports
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

// Redredd-secure-client-tsact Imports
import { queryClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";

// Custom Imports
import ReviewContractSection from "~sections/ReviewContractSection";

const ReviewContract = () => {
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
};

export default ReviewContract;
