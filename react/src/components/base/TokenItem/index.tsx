import { IToken } from "~baseComponents/TokenSelector";

interface TokenItemProps {
  token: IToken;
  onClick: (token: IToken) => void;
  showAmount?: boolean;
  selected: boolean;
}

const TokenItem = (props: TokenItemProps) => {
  const { token, onClick, showAmount, selected } = props;

  return (
    <button className="single-token" onClick={() => onClick(token)}>
      <div className="token-info">
        <div className="token-img">IMG</div>
        <div className="token">
          <div className="token-name">{token.name}</div>
          <div className="token-denom">{token.denom}</div>
        </div>

        {showAmount && <div className="token-amount">{token.amount ?? 0}</div>}
      </div>
    </button>
  );
};

export default TokenItem;
