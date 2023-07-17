// dredd-secure-client-ts Imports
import { IToken } from "~baseComponents/TokenSelector";

interface TokenPreviewProps {
  token: IToken | undefined;
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
