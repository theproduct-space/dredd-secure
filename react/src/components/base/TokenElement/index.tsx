// dredd-secure-client-ts Imports
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";

export interface TokenElementProps {
  onClick?: () => void;
  selectedToken?: Coin;
  baseButton?: JSX.Element;
}

const TokenElement = (props: TokenElementProps) => {
  const { onClick, selectedToken, baseButton } = props;
  const handleOnClick = () => {
    if (onClick) onClick();
    console.log("clicked");
  };

  return (
    <>
      {selectedToken ? (
        <div className="token-display">
          <div className="token-amount">{selectedToken.amount}</div>
          <button className="token-info" onClick={() => handleOnClick()}>
            <div className="token-img"></div>
            <div className="token">
              <div className="token-name">{selectedToken.denom}</div>
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
