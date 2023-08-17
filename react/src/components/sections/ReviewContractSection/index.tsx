// React Imports
import { useState } from "react";

// dredd-secure-client-ts Imports
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";

// Custom Imports
import TokenPreview from "~baseComponents/TokenPreview";
import TokenSelector, { IToken } from "~baseComponents/TokenSelector";
import Tips from "~sections/Tips";
import { V1Beta1Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/rest";
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import useWallet from "../../utils/useWallet";
import assets from "~src/tokens.json";
import { env } from "~src/env";
import { ConditionTypes } from "~sections/CreateContract/AddConditions";
import { toast } from "react-toastify";
import Typography from "~baseComponents/Typography";
import { Link } from "react-router-dom";
import { IContract } from "~sections/CreateContract";

//Assets Imports
import randomCubes from "~assets/random-cubes.webp";
import Card from "~baseComponents/Card";
import Transaction from "~icons/Transaction";
import SideCard from "~baseComponents/SideCard";
import ContentContainer from "~layouts/ContentContainer";
import BaseModal from "~baseComponents/BaseModal/Index";

// Hooks Imports

interface ReviewContractSectionProps {
  contract: IContract;
  onSuccess: () => void;
}

function ReviewContractSection(props: ReviewContractSectionProps) {
  const { contract, onSuccess } = props;
  enum Modals {
    Tips,
  }
  const { address, offlineSigner } = useWallet();
  const [modalToOpen, setModalToOpen] = useState<Modals | undefined>();
  const [selectedTips, setSelectedTips] = useState<IToken | undefined>();
  const [selectedTipAmount, setSelectedTipAmount] = useState<number>(0);
  const [selectedTokenTips, setSelectedTokenTips] = useState<
    IToken | undefined
  >(contract?.tips);
  const messageClient = txClient({
    signer: offlineSigner,
    prefix: "cosmos",
    addr: env.rpcURL,
  });
  const handleSaving = (t: IToken | undefined) => {
    switch (modalToOpen) {
      case Modals.Tips:
        setSelectedTokenTips(t);
        break;
      default:
        break;
    }
    setModalToOpen(undefined);
  };

  const CoinToIToken = (c: V1Beta1Coin | undefined): IToken | undefined => {
    if (c) {
      const token = assets.tokens.find((t) => t.denom === c.denom);
      if (token) {
        console.log("token", token);
        return {
          name: token.name,
          denom: token.denom,
          amount: Number(c.amount),
          logos: token.logos,
          display: token.display,
          chain_name: token.chain_name,
        };
      }
    }

    return;
  };

  const handleConfirmation = async () => {
    if (!contract) return;

    const id = contract.id;
    const request = messageClient.sendMsgFulfillEscrow({
      value: {
        creator: address,
        id: Number(id),
      },
    });

    const response = await toast.promise(request, {
      pending: `Fulfilling Escrow #${id} in-progress`,
      success: `Successfully fulfilled Escrow #${id}!`,
      error: `An error happened while fulfilling Escrow #${id}!`,
    });

    if (response.code == 0) {
      onSuccess();
    }
  };
  const displayModal = () => {
    let modal;
    let showOwnedToken = false;
    switch (modalToOpen) {
      case Modals.Tips:
        modal = selectedTokenTips;
        showOwnedToken = true;
        break;
      default:
        modal = null;
        break;
    }
    return (
      <TokenSelector
        selectedToken={modal}
        onSave={handleSaving}
        ownedToken={showOwnedToken}
        handleClose={() => setModalToOpen(undefined)}
      />
    );
  };
  const formatDate = (timestamp: string): string => {
    console.log("timestamp", timestamp);
    const date = new Date(Number(timestamp) * 1000);
    console.log("date", date);
    return `${String(date.getMonth() + 1).padStart(2, "0")}/${String(
      date.getDate(),
    ).padStart(2, "0")}/${date.getFullYear()}`;
  };

  console.log("ConditionTypes", ConditionTypes);
  console.log("contract", contract);

  return (
    <div>
      <img
        src={randomCubes}
        alt="Dredd-Secure"
        className="object-cover absolute z-0 top-32 right-32 drop-shadow-lightOrange opacity-80 "
        loading="lazy"
      />
      <div className="relative min-h-screen w-full pt-32">
        <Link to="/dashboard">
          <Typography
            variant="body"
            className="font-revalia text-orange px-4 md:px-8 xl:px-16"
          >
            {"< GO BACK"}
          </Typography>
        </Link>
        <div className="relative mx-auto pt-32 max-w-6xl px-4 md:px-8 xl:px-16">
          <div className="title-2">
            {/* <div className="messages">messages here</div> */}
            <Typography variant="h5" className="font-revalia pb-4">
              Escrow Contract #{contract?.id}
            </Typography>
          </div>
        </div>
        <ContentContainer className="max-w-6xl flex gap-4">
          <div className="w-7/12">
            <Card>
              <div className="p-4 md:p-8">
                <div className="flex flex-col gap-6">
                  <div>
                    <Typography variant="h6" className="pb-4">
                      Asset Offered
                    </Typography>
                    <TokenPreview
                      token={CoinToIToken(contract?.initiatorCoins?.[0])}
                      tokenType="initiator"
                      text="Offering"
                    />
                  </div>
                  <div>
                    <Typography variant="h6" className="pb-4">
                      Asset Wanted
                    </Typography>
                    <TokenPreview
                      token={CoinToIToken(contract?.fulfillerCoins?.[0])}
                      tokenType="fulfiller"
                      text="Wanted"
                    />
                  </div>
                </div>
                <Typography variant="h6" className="pt-8">
                  Conditions
                </Typography>
                <div className="py-4 md:py-8">
                  {contract?.conditions?.map((condition, index) => {
                    return (
                      <div key={`condition-${index}`}>
                        <div className="pb-4">
                          <Typography
                            variant="body-small"
                            className="text-white-500"
                          >
                            {condition.label}
                          </Typography>
                          <Typography variant="h6" className="condition-value">
                            {condition.name === "startDate" ||
                            condition.name === "endDate"
                              ? formatDate(condition.value as string)
                              : condition.value}
                          </Typography>
                        </div>
                        {condition.type === "apiCondition" && (
                          <div className="pb-8">
                            <Typography
                              variant="body-small"
                              className="text-white-500"
                            >
                              Token of Interest
                            </Typography>
                            <Typography
                              variant="h6"
                              className="condition-value"
                            >
                              {condition.tokenOfInterest?.name} (
                              {condition.tokenOfInterest?.symbol})
                            </Typography>
                            {condition.subConditions &&
                              condition.subConditions.length > 0 && (
                                <div>
                                  {condition.subConditions.map(
                                    (subCondition, subIndex) => {
                                      let conditionSymbol;
                                      switch (subCondition.conditionType) {
                                        case "eq":
                                          conditionSymbol = "=";
                                          break;
                                        case "gt":
                                          conditionSymbol = ">";
                                          break;
                                        case "lt":
                                          conditionSymbol = "<";
                                          break;
                                        default:
                                          conditionSymbol = "";
                                          break;
                                      }
                                      return (
                                        <div key={`sub-condition-${subIndex}`}>
                                          <Typography
                                            variant="body-small"
                                            className="text-white-500"
                                          >
                                            {subCondition.label}{" "}
                                            {conditionSymbol}{" "}
                                            {subCondition.value}
                                          </Typography>
                                        </div>
                                      );
                                    },
                                  )}
                                </div>
                              )}
                          </div>
                        )}
                      </div>
                    );
                  })}
                </div>
              </div>
              {contract?.status != "closed" && !modalToOpen && (
                <Tips
                  token={selectedTips}
                  onClick={() => setModalToOpen(Modals.Tips)}
                  selectedAmount={selectedTipAmount}
                  setSelectedAmount={setSelectedTipAmount}
                />
              )}
              {contract?.status != "closed" && modalToOpen && (
                <TokenSelector
                  selectedToken={selectedTips}
                  onSave={setSelectedTips}
                  ownedToken={true}
                  handleClose={() => setModalToOpen(undefined)}
                />
              )}
            </Card>
          </div>
          {contract && contract?.status != "closed" && address != "" && (
            <SideCard
              handleConfirmExchange={handleConfirmation}
              contract={contract}
              token={CoinToIToken(contract?.fulfillerCoins?.[0])}
            />
          )}
        </ContentContainer>
      </div>
      <BaseModal
        open={modalToOpen !== undefined}
        handleClose={() => setModalToOpen(undefined)}
      >
        {displayModal()}
      </BaseModal>
    </div>
  );
}

export default ReviewContractSection;
