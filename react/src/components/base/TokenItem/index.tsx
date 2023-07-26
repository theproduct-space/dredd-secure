import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";

interface TokenItemProps {
  token: IToken;
  onClick?: (token: IToken) => void;
  showAmount?: boolean;
  selected: boolean;
  className?: string;
}

const TokenItem = (props: TokenItemProps) => {
  const { token, onClick, showAmount, selected, className } = props;
  const logoUrl = token.logos ? token.logos.svg ?? token.logos.png : undefined;

  return (
    <button
      className={`${`w-full p-2 rounded hover:bg-white-200 ${className}`}`}
      onClick={() => onClick && onClick(token)}
    >
      {selected ? (
        <div className="flex justify-between items-center">
          <div className="flex gap-2 items-center">
            <div className="w-6">
              {logoUrl && <img src={logoUrl} alt="token" />}
            </div>
            <div className="flex flex-col align-start">
              <Typography variant="body" className="uppercase text-left">
                {token.display}
              </Typography>
              <Typography
                variant="body-small"
                className="text-white-500 capitalize text-left"
              >
                {token.name}
              </Typography>
            </div>
          </div>
          {showAmount && (
            <div className="text-white-1000">{token.amount ?? 0}</div>
          )}
        </div>
      ) : (
        <div className="flex justify-between items-center">
          <div className="flex gap-2 items-center">
            <div className="w-6">
              {logoUrl && <img src={logoUrl} alt="token" />}
            </div>
            <div className="flex flex-col align-start">
              <Typography variant="body" className="uppercase text-left">
                {token.display}
              </Typography>
              <Typography
                variant="body-small"
                className="text-white-500 capitalize text-left"
              >
                {token.name}
              </Typography>
            </div>
          </div>
          {showAmount && (
            <div className="text-white-1000">{token.amount ?? 0}</div>
          )}
        </div>
      )}
    </button>
  );
};

export default TokenItem;
