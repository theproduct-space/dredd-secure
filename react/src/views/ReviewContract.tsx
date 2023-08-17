// React Imports
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

// Redredd-secure-client-tsact Imports
import { queryClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";

// Custom Imports
import { toast } from "react-toastify";
import { useClient } from "~hooks/useClient";
import Loading from "~sections/Loading";
import ReviewContractSection from "~sections/ReviewContractSection";
import useWallet from "../components/utils/useWallet";
import Failure from "./Failure";
import Success from "./Success";

const ReviewContract = () => {
  enum SectionState {
    NOT_LOGGED_IN,
    LOADING,
    WALLET_FAILURE,
    CONFIRMATION,
    ESCROW_FULFILLED,
  }

  const { address } = useWallet();

  const [section, setSection] = useState(SectionState.NOT_LOGGED_IN);
  const { id } = useParams<{ id: string }>();
  const [contract, setContract] = useState<EscrowEscrow>();

  const checkIfConnected = () => {
    if (section == SectionState.NOT_LOGGED_IN && address != "") {
      setSection(SectionState.LOADING);
    }
  };
  //TODO use contract instead of EscrowEscrow or make sure Escrow has all the contract info
  const fetchEscrow = async () => {
    try {
      const request = queryClient().queryEscrow(id ?? "");

      const response = await toast.promise(request, {
        pending: `Fetching Escrow #${id} in-progress`,
        success: `Successfully fetched Escrow #${id}!`,
        error: `An error happened while fetching Escrow #${id}!`,
      });

      setContract(response.data.Escrow);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchEscrow();
  }, []);

  useEffect(() => {
    checkIfConnected();

    if (contract) {
      if (contract.status === "closed") setSection(SectionState.CONFIRMATION);
      else if (address != "") handleLoading();
    }
  }, [address, contract]);

  const handleLoading = async () => {
    setTimeout(() => {
      const coin = contract?.fulfillerCoins?.[0];
      useClient()
        .CosmosBankV1Beta1.query.queryBalance(address, { denom: coin?.denom })
        .then((response) => {
          const amount = response.data.balance?.amount;

          if (amount && Number(amount) >= Number(coin?.amount)) {
            setSection(SectionState.CONFIRMATION);
          } else {
            setSection(SectionState.WALLET_FAILURE);
          }
        });
    }, 1200);
  };

  const handleContinueFailureButton = () => {
    setSection(SectionState.LOADING);
  };

  return (
    <>
      {section == SectionState.NOT_LOGGED_IN ||
      section == SectionState.CONFIRMATION ? (
        <ReviewContractSection
          contract={contract}
          onSuccess={() => setSection(SectionState.ESCROW_FULFILLED)}
        />
      ) : section == SectionState.LOADING ? (
        <Loading />
      ) : section == SectionState.WALLET_FAILURE ? (
        <Failure
          errorTitle={
            "The assets needed for this escrow contract were not found in your wallet"
          }
          errorBody={"Please connect a different wallet with these assets"}
          continueButton={
            <button onClick={handleContinueFailureButton}>Try again</button>
          }
        />
      ) : section == SectionState.ESCROW_FULFILLED ? (
        <Success
          successTitle={"Your assets have been exchanged"}
          successBody={"Thank you for interacting with our contract."}
        />
      ) : null}
    </>
  );
};

export default ReviewContract;
