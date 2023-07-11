// dredd-secure-client-ts Imports
import { V1Beta1Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/rest";

interface TokenPreviewProps {
  token: V1Beta1Coin | undefined;
}

const TokenPreview = (props: TokenPreviewProps) => {
  const { token } = props;

  return (
    <div className="token-preview">
      <div className="token-img">IMG</div>
      <div className="token-amount">{token?.amount}</div>
      <div className="token-denom">{token?.denom}</div>
    </div>
  );
};

export default TokenPreview;
