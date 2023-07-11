// React Imports
import React, { useState } from "react";

// dredd-secure-client-ts Imports
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";

// Custom Imports
import TokenPreview from "~baseComponents/TokenPreview";
import TokenSelector from "~baseComponents/TokenSelector";
import { ICondition, IContract } from "~sections/CreateContract";
import Tips from "~sections/Tips";

// Hooks Imports
import useKeplr from "~def-hooks/useKeplr";

interface ReviewContractSectionProps {
  contract: EscrowEscrow | undefined;
}

export const ConditionTypes: ICondition[] = [
  {
    type: "Starting Date",
    prop: "startDate",
  },
  {
    type: "Deadline",
    prop: "endDate",
  },
];

function ReviewContractSection(props: ReviewContractSectionProps) {
  const { contract } = props;
  const [modalOpened, setModalOpened] = useState<boolean>(false);
  const [selectedTips, setSelectedTips] = useState<Coin | undefined>();

  const [address, setAddress] = useState("c"); // For testing purposes only, // TODO: Get address from keplr or other wallet manager

  return (
    <div>
      <div className="card">
        <button onClick={() => setAddress("cosmosAddresss")}>
          Connect with testing wallet
        </button>{" "}
        {/* For testing purposes only */}
        <div className="card-subtitle">What the owner wants</div>
        <TokenPreview token={contract?.fulfillerCoins?.[0]} />
        <div className="card-subtitle">What they are offering</div>
        <TokenPreview token={contract?.initiatorCoins?.[0]} />
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
          />
        )}
      </div>
      {contract?.status != "closed" && address && address != "" && (
        <div className="card">
          <div className="card-title">Confirm</div>
          <div className="bold">Transaction cost</div>
          <div className="text">FREE</div>
          <div className="bold">What you're offering</div>
          <TokenPreview token={contract?.fulfillerCoins?.[0]} />
          <label>
            <input type="checkbox"></input>
            by checking this box ......
          </label>
          <button
            onClick={() => {
              console.log("confirm exchange");
            }}
          >
            Confirm Exchange
          </button>{" "}
          {/* TODO: Confirmation logic */}
        </div>
      )}
    </div>
  );
}

export default ReviewContractSection;
