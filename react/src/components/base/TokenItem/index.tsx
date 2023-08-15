/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
import React, { useEffect, useState } from "react";
import { TextField } from "@mui/material";
import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";

interface TokenItemProps {
  token: IToken;
  onClick?: (token: IToken) => void;
  showAmount?: boolean;
  selected: boolean;
  input?: boolean;
  className?: string;
  selectedAmount?: number;
  setSelectedAmount?: (amount: number) => void;
  noMax?: boolean;
}

type PropsWithSelectedAmount = Required<TokenItemProps>;

const TokenItem = (props: TokenItemProps) => {
  const {
    token,
    onClick,
    showAmount,
    selected,
    input,
    selectedAmount,
    setSelectedAmount,
    noMax,
    className,
  } = props as PropsWithSelectedAmount;
  const logoUrl = token.logos ? token.logos.svg ?? token.logos.png : undefined;
  const [inputValue, setInputValue] = useState<string>(
    selectedAmount?.toString() ?? "",
  );

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = e.target.value;
    const isValidInput = /^\d+(\.\d{0,4})?$/.test(newValue);
    if (isValidInput) {
      const enteredAmount = parseFloat(newValue);
      const maxAmount = handleMaxAmount();
      if (props.noMax || enteredAmount <= parseFloat(maxAmount)) {
        setInputValue(newValue);
        setSelectedAmount && setSelectedAmount(enteredAmount);
      } else {
        setInputValue(maxAmount);
        setSelectedAmount && setSelectedAmount(parseFloat(maxAmount));
      }
    } else if (newValue === "") {
      setInputValue("");
      setSelectedAmount && setSelectedAmount(0);
    }
  };
  const handleMaxAmount = () => {
    // no max amount if nomax is true
    if (props.noMax) {
      return "";
    } else {
      return token.amount ? token.amount.toString() : "0";
    }
  };

  return (
    <>
      {selected &&
        !input && ( //show this when selected and not input
          <button
            className={`${`w-full p-2 rounded border-[1px] border-white-200 bg-buttonBg ${className}`}`}
            onClick={() => onClick && onClick(token)}
          >
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
          </button>
        )}
      {!selected &&
        !input && ( //show this when not selected and not input
          <button
            className={`${`w-full p-2 rounded hover:bg-white-200 ${className}`}`}
            onClick={() => onClick && onClick(token)}
          >
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
          </button>
        )}
      {input &&
        selected && ( //show this when selected and input
          <TextField
            fullWidth
            variant="outlined"
            type="number"
            value={inputValue}
            onChange={handleInputChange}
            inputProps={{
              min: 0,
              max: handleMaxAmount,
            }}
            sx={{
              "& input::-webkit-outer-spin-button, & input::-webkit-inner-spin-button":
                {
                  "-webkit-appearance": "none",
                  margin: 0,
                },
            }}
            inputMode="decimal"
            InputProps={{
              endAdornment: (
                <button onClick={() => onClick && onClick(token)}>
                  <div className="flex gap-2 items-center">
                    <div className="w-6">
                      {logoUrl && <img src={logoUrl} alt="token" />}
                    </div>
                    <div className="flex flex-col align-start">
                      <Typography
                        variant="body"
                        className="uppercase text-left"
                      >
                        {token.display}
                      </Typography>
                    </div>
                  </div>
                </button>
              ),
            }}
          />
        )}
    </>
  );
};

export default TokenItem;
