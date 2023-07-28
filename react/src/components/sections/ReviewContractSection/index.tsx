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
  const messageClient = txClient({
    signer: offlineSigner,
    prefix: "cosmos",
    addr: env.rpcURL,
  });

  const CoinToIToken = (c: V1Beta1Coin | undefined): IToken | undefined => {
    if (c) {
      const token = assets.tokens.find((t) => t.denom === c.denom);
      if (token) {
        return {
          name: token.denom, // TODO: To change with the name got from a list of tokens
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
    messageClient
      .sendMsgFulfillEscrow({
        value: {
          creator: address,
          id: Number(contract?.id),
        },
      })
      .then((response) => {
        if (response.code == 0) {
          onSuccess();
        }
      });
  };

  return (
    <div>
      <div className="messages">messages here</div>
      <div className="title">Review escrow Contract #{contract?.id}</div>
      <div className="card">
        {/* For testing purposes only */}
        <div className="card-subtitle">What the owner wants</div>
        <TokenPreview token={CoinToIToken(contract?.fulfillerCoins?.[0])} />
        <div className="card-subtitle">What they are offering</div>
        <TokenPreview token={CoinToIToken(contract?.initiatorCoins?.[0])} />
        <div className="card-subtitle">Conditions</div>
        <div className="conditions">
          {ConditionTypes.map((condition, index) => {
            if (!contract?.[condition.prop]) return;

            return (
              <div className="condition" key={`condition-${index}`}>
                <div className="condition-type bold">{condition.type}</div>
                <div className="condition-value">
                  {contract[condition.prop]}
                </div>
              </div>
            );
          })}
        </div>
        {contract?.status != "closed" && !modalOpened && (
          <Tips
            selectedToken={selectedTips}
            onClick={() => setModalOpened(true)}
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
      </div>
      {contract?.status != "closed" && address != "" && (
        <div className="card">
          <div className="card-title">Confirm</div>
          <div className="bold">Transaction cost</div>
          <div className="text">FREE</div>
          <div className="bold">What you're offering</div>
          <TokenPreview token={CoinToIToken(contract?.fulfillerCoins?.[0])} />
          <button onClick={handleConfirmation}>Confirm Exchange</button>
        </div>
      )}
    </div>
  );
}

export default ReviewContractSection;
