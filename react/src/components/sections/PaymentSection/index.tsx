// React Imports
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";
import { txClient } from "dredd-secure-client-ts/dreddsecure.escrow";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

// Custom Imports
import TokenPreview from "~baseComponents/TokenPreview";
import Account from "~sections/Account";
import { IContract } from "~sections/CreateContract";
import useWallet from "../../utils/useWallet";

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
    const startDate: string = new Date(
      contract.conditions?.find((e) => e.prop == "startDate")?.value ??
        Date.now(),
    )
      .getTime()
      .toString();

    console.log("payment startDate", startDate);
    const endDate: string = new Date(
      contract.conditions?.find((e) => e.prop == "startDate")?.value ??
        new Date("9999-12-31"),
    )
      .getTime()
      .toString();

    const response = await messageClient.sendMsgCreateEscrow({
      value: {
        creator: address,
        initiatorCoins: initiatorCoins,
        fulfillerCoins: fulfillerCoins,
        startDate: startDate,
        endDate: endDate,
      },
    });

    if (response.code == 0) navigate("/dashboard");
    else {
      setErrorMessage(response.rawLog);
    }
  };

  const handleGoBack = () => {
    navigate("/escrow/create", {
      state: contract,
    });
  };

  return (
    <div>
      {errorMessage && <div className="error-message">{errorMessage}</div>}
      <div className="Title">Review and Confirm Exchange</div>
      <button className="back-button" onClick={handleGoBack}>
        Go Back
      </button>
      <div className="card">
        <div className="card-subtitle">Conditions</div>

        {contract?.conditions?.map((condition, index) => {
          return (
            <div key={`condition-${index}`}>
              <div className="condition-name">{condition.type}</div>
              <div className="condition-value">{condition.value}</div>
            </div>
          );
        })}

        <div className="card-subtitle">You are exchanging</div>
        <div>
          <TokenPreview token={contract.initiatorCoins} />
          <div className="exchange-icon"></div>
          <TokenPreview token={contract.fulfillerCoins} />
        </div>
        {contract?.status != "closed" && address != "" && (
          <div className="card">
            <div className="card-title">Confirm</div>
            <div className="bold">Transaction cost</div>
            <div className="text">FREE</div>
            {contract?.tips ? (
              <>
                <div className="donation-review">Donation to dreddsecure</div>
                <TokenPreview token={contract.tips} />
              </>
            ) : (
              <>
                <div className="donation-review">
                  Donation to dreddsecure <button>+Add</button>
                </div>
                <div className="donation-amount">0.00</div>
              </>
            )}

            <button onClick={handleConfirmExchange}>Confirm Exchange</button>
          </div>
        )}
      </div>
      <Account />
    </div>
  );
};

export default PaymentSection;
