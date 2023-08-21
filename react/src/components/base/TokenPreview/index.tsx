// dredd-secure-client-ts Imports
import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";

interface TokenPreviewProps {
  token: IToken | undefined;
  className?: string;
  tokenType?: "initiator" | "fulfiller";
  text: string;
}

const TokenPreview = (props: TokenPreviewProps) => {
  const { token, className, tokenType } = props;
  const logoUrl = token?.logos ? token.logos.svg ?? token.logos.png : undefined;

  return (
    <>
      {/* {tokenType === "initiator" && ( */}
      <div
        className={`${`w-full p-2 rounded border-[1px] border-white-200 bg-buttonBg ${className}`}`}
      >
        <Typography
          variant="body-small"
          className="uppercase text-left text-white-500"
        >
          {props.text}
        </Typography>
        <div className="flex justify-between items-center">
          <div className="flex gap-2 items-center">
            <div className="w-8">
              {logoUrl && <img src={logoUrl} alt="token" />}
            </div>
            <div className="flex flex-col align-start">
              <Typography variant="body" className="text-white-1000">
                {token?.amount ?? 0}
              </Typography>
              <Typography variant="body" className="uppercase text-left">
                {token?.display}
              </Typography>
            </div>
          </div>
        </div>
      </div>
      {/* )} */}
      {/* {tokenType === "fulfiller" && (
        <div
          className={`${`w-full p-2 rounded border-[1px] border-white-200 bg-buttonBg ${className}`}`}
        >
          <Typography
            variant="body-small"
            className="uppercase text-left text-white-500"
          >
            {props.text}
          </Typography>
          <div className="flex justify-between items-center">
            <div className="flex gap-2 items-center">
              <div className="w-8">
                {logoUrl && <img src={logoUrl} alt="token" />}
              </div>
              <div className="flex flex-col align-start">
                <Typography variant="body" className="uppercase text-left">
                  {token?.display}
                </Typography>
                <Typography
                  variant="body-small"
                  className="text-white-500 capitalize text-left"
                >
                  {token?.name}
                </Typography>
              </div>
            </div>
          </div>
        </div>
      )} */}
    </>
  );
};

export default TokenPreview;
