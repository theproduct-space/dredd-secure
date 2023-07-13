// react Imports
import { useState } from "react";

// dredd-secure-client-ts Imports
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";

// Custom Imports
import TokenItem from "~baseComponents/TokenItem";

export interface IToken {
  name: string;
  amount?: number;
  denom: string;
}

export interface TokenSelectorProps {
  onSave: (token: IToken | undefined) => void;
  address?: string;
  selectedToken?: IToken;
}

const TokenSelector = (props: TokenSelectorProps) => {
  const [selectedToken, setSelectedToken] = useState(props.selectedToken);
  const { onSave, address } = props;

  const displayOwnedToken = () => {
    // TODO: When the wallet connector will be implemented, get all tokens from wallet logic here.
  };

  const displayAllToken = () => {
    const tokens: IToken[] = [
      {
        name: "token",
        denom: "utok",
      },
      {
        name: "ATOM",
        denom: "uatom",
      },
    ];

    return tokens.map((token) => {
      return (
        <TokenItem
          token={token}
          onClick={(t) => onSave(t)}
          showAmount={true}
          selected={selectedToken?.denom == token.denom}
        />
      );
    });
  };

  return (
    <div className="modal">
      <div className="card">
        <div className="card-headers">
          Select a token
          <div className="search-bar">Search token</div>
        </div>
        <div className="card-body">{displayAllToken()}</div>
      </div>
    </div>
  );
};

export default TokenSelector;
