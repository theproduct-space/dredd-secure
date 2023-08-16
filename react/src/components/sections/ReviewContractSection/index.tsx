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
import ModalContainer from "~layouts/ModalContainer";
import Typography from "~baseComponents/Typography";
import { Link } from "react-router-dom";

//Assets Imports
import randomCubes from "~assets/random-cubes.webp";
import Card from "~baseComponents/Card";
import Transaction from "~icons/Transaction";
import SideCard from "~baseComponents/SideCard";

// Hooks Imports

interface ReviewContractSectionProps {
  contract: EscrowEscrow | undefined;
  onSuccess: () => void;
}

function ReviewContractSection(props: ReviewContractSectionProps) {
  const { contract, onSuccess } = props;
  const { address, offlineSigner } = useWallet();
  const [modalOpened, setModalOpened] = useState<boolean>(false);
  const [selectedTips, setSelectedTips] = useState<IToken | undefined>();
  const [selectedTipAmount, setSelectedTipAmount] = useState<number>(0);
  const messageClient = txClient({
    signer: offlineSigner,
    prefix: "cosmos",
    addr: env.rpcURL,
  });

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
        <ModalContainer className="max-w-6xl flex gap-4">
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
                  {ConditionTypes.map((condition, index) => {
                    if (!contract?.[condition.name]) return;
                    return (
                      <div className="condition" key={`condition-${index}`}>
                        <Typography
                          variant="body-small"
                          className="text-white-500"
                        >
                          {condition.label}
                        </Typography>
                        <Typography variant="h6" className="condition-value">
                          {condition.name === "startDate" ||
                          condition.name === "endDate"
                            ? formatDate(contract[condition.name] as string)
                            : contract[condition.name]}
                        </Typography>
                      </div>
                    );
                  })}
                </div>
              </div>
              {contract?.status != "closed" && !modalOpened && (
                <Tips
                  token={selectedTips}
                  onClick={() => setModalOpened(true)}
                  selectedAmount={selectedTipAmount}
                  setSelectedAmount={setSelectedTipAmount}
                />
              )}
              {contract?.status != "closed" && modalOpened && (
                <TokenSelector
                  selectedToken={selectedTips}
                  onSave={setSelectedTips}
                  ownedToken={true}
                  handleClose={() => setModalOpened(false)}
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
        </ModalContainer>
      </div>
    </div>
  );
}

export default ReviewContractSection;
