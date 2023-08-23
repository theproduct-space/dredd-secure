// dredd-secure-client-ts Imports
import { Checkbox } from "@mui/material";
import { EscrowEscrow } from "dredd-secure-client-ts/dreddsecure.escrow/rest";

// React import
import { useEffect, useState } from "react";
import Button from "~baseComponents/Button";
import Card from "~baseComponents/Card";
import TokenPreview from "~baseComponents/TokenPreview";
import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";
import { IContract } from "~sections/CreateContract";

export interface SideCardProps {
  handleConfirmExchange: () => void;
  contract: IContract | EscrowEscrow;
  paymentInterface?: boolean;
  token?: IToken;
  walletFailure?: boolean;
}

const SideCard = (props: SideCardProps) => {
  const {
    handleConfirmExchange,
    contract,
    paymentInterface,
    token,
    walletFailure,
  } = props;
  const isPaymentInterface = "tips" in contract && paymentInterface;
  const [checked, setChecked] = useState<boolean>(false);
  const handleCheck = () => {
    setChecked(!checked);
  };
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
                    <TokenPreview token={contract.tips} text="" />
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
                {token && (
                  <TokenPreview token={token} tokenType="fulfiller" text="" />
                )}
              </div>
            </div>
          )}
          {paymentInterface && (
            <div>
              <div className="flex items-start">
                <Checkbox checked={checked} onChange={handleCheck} />
                <Typography variant="small" className="pb-2">
                  *By checking this box, you are agreeing to this contract and
                  your assets will be exchanged upon review of this submission
                </Typography>
              </div>
              {checked ? (
                <Button
                  text="Deploy Contract"
                  className="w-full"
                  onClick={handleConfirmExchange}
                />
              ) : (
                <Button
                  text="Deploy Contract"
                  className="w-full"
                  onClick={handleConfirmExchange}
                  disabled
                />
              )}
            </div>
          )}
          {walletFailure && !paymentInterface && (
            <div>
              <Typography variant="body-small" className="text-red-500">
                *Not enough funds in your wallet to complete this transaction
              </Typography>
              <Button
                text="Confirm Exchange"
                className="w-full"
                disabled
                onClick={handleConfirmExchange}
              />
            </div>
          )}
          {!paymentInterface && !walletFailure && (
            <Button
              text="Confirm Exchange"
              className="w-full"
              onClick={handleConfirmExchange}
            />
          )}
        </div>
      </Card>
    </>
  );
};

export default SideCard;
