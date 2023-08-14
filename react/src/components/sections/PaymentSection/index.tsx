// React Imports
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

// Custom Imports
import TokenPreview from "~baseComponents/TokenPreview";
import Account from "~sections/Account";
import { IContract } from "~sections/CreateContract";
import useWallet from "../../utils/useWallet";
import { toast } from "react-toastify";

// Assets
import randomCubes from "~assets/random-cubes.webp";
import Typography from "~baseComponents/Typography";
import ModalContainer from "~layouts/ModalContainer";
import Card from "~baseComponents/Card";
import Transaction from "~icons/Transaction";
import Tips from "~sections/Tips";
import Button from "~baseComponents/Button";

interface PaymentSectionProps {
  contract: IContract;
}

const PaymentSection = (props: PaymentSectionProps) => {
  const { contract } = props;
  const { address, offlineSigner } = useWallet();
  const [errorMessage, setErrorMessage] = useState<string | undefined>();
  const navigate = useNavigate();
  const messageClient = txClient({
    signer: offlineSigner,
    prefix: "cosmos",
    addr: "http://localhost:26657",
  });

  const handleConfirmExchange = async () => {
    const initiatorCoins: Coin[] = [
      {
        denom: contract.initiatorCoins.denom,
        amount: contract.initiatorCoins.amount?.toString() ?? "0",
      },
    ];
    const fulfillerCoins: Coin[] = [
      {
        denom: contract.fulfillerCoins.denom,
        amount: contract.fulfillerCoins.amount?.toString() ?? "0",
      },
    ];

    const startDate = String(
      contract.conditions?.find((e) => e.prop == "startDate")?.value ??
        (Date.now() / 1000).toFixed(),
    );

    const endDate = String(
      contract.conditions?.find((e) => e.prop == "endDate")?.value ??
        (new Date("9999-12-31").getTime() / 1000).toFixed(),
    );

    console.log({ startDate });
    console.log({ endDate });

    const request = messageClient.sendMsgCreateEscrow({
      value: {
        creator: address,
        initiatorCoins: initiatorCoins,
        fulfillerCoins: fulfillerCoins,
        startDate: startDate,
        endDate: endDate,
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

  const handleGoBack = () => {
    navigate("/escrow/create", {
      state: contract,
    });
  };

  const formatDate = (timestamp: string): string => {
    const date = new Date(Number(timestamp));
    return `${String(date.getMonth() + 1).padStart(2, "0")}/${String(
      date.getDate(),
    ).padStart(2, "0")}/${date.getFullYear()}`;
  };
  console.log(
    "START",
    contract.conditions?.find((e) => e.prop == "startDate")?.value,
  );
  console.log(
    "END",
    contract.conditions?.find((e) => e.prop == "endDate")?.value,
  );
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
        <Link to="/escrow/create">
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
        <ModalContainer className="max-w-6xl flex gap-4">
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
                        <Typography
                          variant="body-small"
                          className="text-white-500"
                        >
                          {condition.type}
                        </Typography>
                        <Typography variant="h6" className="condition-value">
                          {condition.prop === "startDate" ||
                          condition.prop === "endDate"
                            ? formatDate(condition.value as string)
                            : condition.value}
                        </Typography>
                      </div>
                    );
                  })}{" "}
                </div>
                <Typography variant="h6" className="font-revalia pb-8">
                  You are exchanging
                </Typography>
                <div className="flex justify-between items-center gap-4">
                  <TokenPreview
                    token={contract.initiatorCoins}
                    tokenType="initiator"
                  />
                  <Transaction className="w-12" />
                  <TokenPreview
                    token={contract.fulfillerCoins}
                    tokenType="fulfiller"
                  />
                </div>
              </div>
            </Card>
          </div>
          {/* {contract?.status != "closed" && address != "" && ( */}
          <Card className="w-4/12">
            <div className="p-4 md:p-8 flex flex-col justify-between ">
              <div className="flex flex-col gap-4 pb-8">
                <Typography variant="h6" className="font-revalia pb-4">
                  Confirm
                </Typography>
                <div>
                  <Typography
                    variant="body-small"
                    className="text-white-500 uppercase"
                  >
                    Transaction cost
                  </Typography>
                  <Typography variant="h6">FREE</Typography>
                </div>
                <div>
                  {contract?.tips ? (
                    <>
                      <Typography
                        variant="body-small"
                        className="text-white-500 uppercase py-4"
                      >
                        Donation to dreddsecure
                      </Typography>
                      <TokenPreview token={contract.tips} />
                    </>
                  ) : (
                    <>
                      <div className="flex justify-between">
                        <Typography
                          variant="body-small"
                          className="text-white-500 uppercase"
                        >
                          Donation to dreddsecure
                        </Typography>
                        <button>
                          <Typography
                            variant="body-small"
                            className="text-orange uppercase"
                          >
                            +Add
                          </Typography>
                        </button>
                      </div>
                      <Typography variant="h6">0.00</Typography>
                    </>
                  )}
                </div>
              </div>
              <Button
                text="Deploy Contract"
                className="w-full"
                onClick={handleConfirmExchange}
              />
            </div>
          </Card>
          {/* )} */}
        </ModalContainer>
      </div>
    </div>
  );
};

export default PaymentSection;
