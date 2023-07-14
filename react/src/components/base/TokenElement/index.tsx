// dredd-secure-client-ts Imports

// React import
import { useState } from "react";
import { IToken } from "~baseComponents/TokenSelector";

export interface TokenElementProps {
  onClick?: () => void;
  selectedToken?: IToken;
  baseButton?: JSX.Element;
}

const TokenElement = (props: TokenElementProps) => {
  const { onClick, baseButton, selectedToken } = props;

  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (selectedToken) selectedToken.amount = Number(event.target.value);
  };

  const handleOnClick = () => {
    if (onClick) onClick();
  };

  return (
    <>
      {selectedToken ? (
        <div className="token-display">
          <input
            className="token-amount"
            value={
              !selectedToken || selectedToken?.amount == 0
                ? ""
                : selectedToken.amount
            }
            type="number"
            onChange={handleAmountChange}
          ></input>
          <button className="token-info" onClick={() => handleOnClick()}>
            <div className="token-img"></div>
            <div className="token">
              <div className="token-name">{selectedToken.name}</div>
              <div className="token-denom">{selectedToken.denom}</div>
            </div>
          </button>
        </div>
      ) : (
        <button onClick={() => handleOnClick()}>{baseButton}</button>
      )}
    </>
  );
};

export default TokenElement;
