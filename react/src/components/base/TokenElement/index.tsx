// dredd-secure-client-ts Imports

// React import
import { TextField } from "@mui/material";
import { useEffect, useState } from "react";
import { IToken } from "~baseComponents/TokenSelector";
import Typography from "~baseComponents/Typography";

export interface TokenElementProps {
  onClick: (token: IToken) => void;
  token: IToken;
  baseButton?: JSX.Element;
  selectedAmount: number;
  setSelectedAmount: (amount: number) => void;
}

const TokenElement = (props: TokenElementProps) => {
  const { onClick, baseButton, token, selectedAmount, setSelectedAmount } =
    props;
  const [inputValue, setInputValue] = useState<string>("");

  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (token) token.amount = Number(event.target.value);
  };
  const logoUrl = token.logos ? token.logos.svg ?? token.logos.png : undefined;
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = e.target.value;
    const isValidInput = /^\d+(\.\d{0,4})?$/.test(newValue);
    if (isValidInput) {
      const enteredAmount = parseFloat(newValue);
      const maxAmount =
        token.amount !== undefined ? token.amount.toString() : "0";
      if (enteredAmount <= parseFloat(maxAmount)) {
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
  useEffect(() => {
    setInputValue(selectedAmount?.toString() ?? "");
  }, [selectedAmount]);

  return (
    <>
      <TextField
        fullWidth
        variant="outlined"
        type="number"
        value={inputValue}
        onChange={handleInputChange}
        inputProps={{
          min: 0,
          max: token.amount ? token.amount.toString() : "",
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
                  <Typography variant="body" className="uppercase text-left">
                    {token.display}
                  </Typography>
                </div>
              </div>
            </button>
          ),
        }}
      />
    </>
  );
};

export default TokenElement;
