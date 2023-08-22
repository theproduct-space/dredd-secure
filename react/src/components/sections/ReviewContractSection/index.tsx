// React Imports
import { useEffect, useState } from "react";

// dredd-secure-client-ts Imports

// Custom Imports
import { V1Beta1Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/rest";
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { Link } from "react-router-dom";
import { toast } from "react-toastify";
import TokenPreview from "~baseComponents/TokenPreview";
import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";
import { ConditionTypes } from "~sections/CreateContract/AddConditions";
import { env } from "~src/env";
import assets from "~src/tokens.json";
import useWallet from "../../utils/useWallet";

//Assets Imports
import randomCubes from "~assets/random-cubes.webp";
import Card from "~baseComponents/Card";
import SideCard from "~baseComponents/SideCard";
import ContentContainer from "~layouts/ContentContainer";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";
import { CoinToIToken } from "~utils/tokenTransformer";

// Hooks Imports

interface ReviewContractSectionProps {
  contract: EscrowEscrow | undefined;
  onSuccess: () => void;
}
export interface ParsedCondition {
  label: string;
  name: string;
  value: string | number;
  type?: string;
  tokenOfInterest?: {
    name: string;
    symbol: string;
  };
  subConditions?: Array<{
    label: string;
    conditionType: string;
    value: string | number;
  }>;
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
  const [parsedConditions, setParsedConditions] = useState<ParsedCondition[]>(
    [],
  );
  // const [selectedTokenTips, setSelectedTokenTips] = useState<
  //   IToken | undefined
  // >(contract?.tips);
  const messageClient = txClient({
    signer: offlineSigner,
    prefix: "cosmos",
    addr: env.rpcURL,
  });
  // const handleSaving = (t: IToken | undefined) => {
  //   switch (modalToOpen) {
  //     case Modals.Tips:
  //       setSelectedTokenTips(t);
  //       break;
  //     default:
  //       break;
  //   }
  //   setModalToOpen(undefined);
  // };

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
    });

    if (response.code == 0) {
      toast.success(`Successfully fulfilled Escrow #${id}!`);
      onSuccess();
    } else {
      toast.error(`An error happened while fulfilling Escrow #${id}!`);
      toast.error(response.rawLog);
    }
  };
  // const displayModal = () => {
  //   let modal;
  //   let showOwnedToken = false;
  //   switch (modalToOpen) {
  //     case Modals.Tips:
  //       modal = selectedTokenTips;
  //       showOwnedToken = true;
  //       break;
  //     default:
  //       modal = null;
  //       break;
  //   }
  //   return (
  //     <TokenSelector
  //       selectedToken={modal}
  //       onSave={handleSaving}
  //       ownedToken={showOwnedToken}
  //       handleClose={() => setModalToOpen(undefined)}
  //     />
  //   );
  // };
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

  useEffect(() => {
    if (contract && contract.apiConditions) {
      const conditions = JSON.parse(contract.apiConditions);
      setParsedConditions(conditions);
    }
  }, [contract]);
  console.log("parsedConditions", parsedConditions);

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
        <ContentContainer className="max-w-6xl flex gap-4 pb-24">
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
                  {parsedConditions.map((condition, index) => {
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
              {/* {contract?.status != "closed" && !modalToOpen && (
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
              )} */}
            </Card>
          </div>
          {contract &&
            contract?.status != "closed" &&
            address != "" &&
            contract?.initiator !== address && (
              <SideCard
                handleConfirmExchange={handleConfirmation}
                contract={contract}
                token={CoinToIToken(contract?.fulfillerCoins?.[0])}
              />
            )}
        </ContentContainer>
      </div>
      {/* <BaseModal
        open={modalToOpen !== undefined}
        handleClose={() => setModalToOpen(undefined)}
      >
        {displayModal()}
      </BaseModal> */}
    </div>
  );
}

export default ReviewContractSection;
