// dredd-secure-client-ts Imports

// React import
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";
import Button from "~baseComponents/Button";
import Card from "~baseComponents/Card";
import TokenPreview from "~baseComponents/TokenPreview";
import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";
import { IContract } from "~sections/CreateContract";
import { CoinToIToken } from "~utils/tokenTransformer";

export interface SideCardProps {
  handleConfirmExchange: () => void;
  contract: EscrowEscrow;
  paymentInterface?: boolean;
  token?: IToken;
}
const SideCard = (props: SideCardProps) => {
  const { handleConfirmExchange, contract, paymentInterface, token } = props;
  const isPaymentInterface = "tips" in contract && paymentInterface;
  return (
    <>
      <Card className="w-4/12 h-fit">
        <div className="p-4 md:p-8 flex flex-col justify-between">
          {isPaymentInterface ? (
            <div className="flex flex-col gap-4 pb-8">
              <Typography variant="h6" className="font-revalia pb-4">
                Confirm
              </Typography>
              <div>
                <Typography
                  variant="body-small"
                  className="text-white-500 uppercase"
                >
                  Transaction cost
                </Typography>
                <Typography variant="h6">FREE</Typography>
              </div>
              <div>
                {contract?.tips ? (
                  <>
                    <Typography
                      variant="body-small"
                      className="text-white-500 uppercase py-4"
                    >
                      Donation to dreddsecure
                    </Typography>
                    <TokenPreview token={CoinToIToken(contract.tips[0])} text="" />
                  </>
                ) : (
                  <>
                    <div className="flex justify-between">
                      <Typography
                        variant="body-small"
                        className="text-white-500 uppercase"
                      >
                        Donation to dreddsecure
                      </Typography>
                      <button>
                        <Typography
                          variant="body-small"
                          className="text-orange uppercase"
                        >
                          +Add
                        </Typography>
                      </button>
                    </div>
                    <Typography variant="h6">0.00</Typography>
                  </>
                )}
              </div>
            </div>
          ) : (
            <div className="flex flex-col gap-4 pb-8">
              <Typography variant="h6" className="font-revalia pb-4">
                Confirm
              </Typography>
              <div>
                <Typography
                  variant="body-small"
                  className="text-white-500 uppercase"
                >
                  Transaction cost
                </Typography>
                <Typography variant="h6">FREE</Typography>
              </div>
              <div>
                <Typography
                  variant="body-small"
                  className="text-white-500 uppercase"
                >
                  What You're Offering
                </Typography>
                {/*todo add token preview*/}
                {token && (
                  <TokenPreview token={token} tokenType="fulfiller" text="" />
                )}
              </div>
              {/*todo add agreement checkbox*/}
            </div>
          )}
          {isPaymentInterface ? (
            <Button
              text="Deploy Contract"
              className="w-full"
              onClick={handleConfirmExchange}
            />
          ) : (
            <Button
              text="Confirm Exchange"
              className="w-full"
              //todo pass handle confirmation instead
              onClick={handleConfirmExchange}
            />
          )}
        </div>
      </Card>
    </>
  );
};

export default SideCard;
