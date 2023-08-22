// React Imports
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

// Custom Imports
import { toast } from "react-toastify";
import TokenPreview from "~baseComponents/TokenPreview";
import { IContract } from "~sections/CreateContract";
import { env } from "~src/env";
import useWallet from "../../utils/useWallet";

// Assets
import randomCubes from "~assets/random-cubes.webp";
import Card from "~baseComponents/Card";
import SideCard from "~baseComponents/SideCard";
import Typography from "~baseComponents/Typography";
import Transaction from "~icons/Transaction";
import ContentContainer from "~layouts/ContentContainer";
import { ContractToEscrow } from "~utils/tokenTransformer";

interface PaymentSectionProps {
  contract: IContract;
}

const PaymentSection = (props: PaymentSectionProps) => {
  const { contract } = props;
  console.log("contract", contract)
  const { address, offlineSigner } = useWallet();
  const [errorMessage, setErrorMessage] = useState<string | undefined>();
  const navigate = useNavigate();
  const messageClient = txClient({
    signer: offlineSigner,
    prefix: "cosmos",
    addr: env.rpcURL,
  });

  const handleConfirmExchange = async () => {
    const c = ContractToEscrow(contract);

    const request = messageClient.sendMsgCreateEscrow({
      value: {
        creator: address,
        ...c
      },
    });

    const response = await toast.promise(request, {
      pending: `Creating the Escrow in-progress`,
      success: `Successfully created Escrow!`,
      error: `An error happened while creating the Escrow.`,
    });

    if (response.code == 0) navigate("/dashboard");
    else setErrorMessage(response.rawLog);
  };

  const formatDate = (timestamp: string): string => {
    console.log("timestamp", timestamp);
    const date = new Date(Number(timestamp) * 1000);
    console.log("date", date);
    return `${String(date.getMonth() + 1).padStart(2, "0")}/${String(
      date.getDate(),
    ).padStart(2, "0")}/${date.getFullYear()}`;
  };

  return (
    <div>
      <img
        src={randomCubes}
        alt="Dredd-Secure"
        className="object-cover absolute z-0 top-32 right-32 drop-shadow-lightOrange opacity-80 "
        loading="lazy"
      />
      {errorMessage && <div className="error-message">{errorMessage}</div>}
      <div className="relative min-h-screen w-full pt-32">
        <Link to="/escrow/create" state={contract}>
          <Typography
            variant="body"
            className="font-revalia text-orange px-4 md:px-8 xl:px-16"
          >
            {"< GO BACK"}
          </Typography>
        </Link>
        <div className="relative mx-auto pt-32 max-w-6xl px-4 md:px-8 xl:px-16">
          <Typography variant="small" className="text-white-500">
            STEP 2
          </Typography>
          <div className="title-2">
            <Typography variant="h5" className="font-revalia pb-4">
              Review and Confirm Exchange
            </Typography>
          </div>
        </div>
        <ContentContainer className="max-w-6xl flex gap-4">
          <div className="w-7/12">
            <Card progress="75">
              <div className="p-4 md:p-8">
                <Typography variant="h6" className="font-revalia">
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
                <Typography variant="h6" className="font-revalia pb-8">
                  You are exchanging
                </Typography>
                <div className="flex justify-between items-center gap-4">
                  <TokenPreview
                    token={contract.initiatorCoins}
                    tokenType="initiator"
                    text="Your Asset"
                  />
                  <Transaction className="w-12" />
                  <TokenPreview
                    token={contract.fulfillerCoins}
                    tokenType="fulfiller"
                    text="Asset You Want"
                  />
                </div>
              </div>
            </Card>
          </div>
          <SideCard
            handleConfirmExchange={handleConfirmExchange}
            contract={ContractToEscrow(contract)}
            paymentInterface
          />
        </ContentContainer>
      </div>
    </div>
  );
};

export default PaymentSection;
