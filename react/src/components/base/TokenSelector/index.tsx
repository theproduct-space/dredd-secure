// react Imports
import { useEffect, useState } from "react";

// dredd-secure-client-ts Imports

// Custom Imports
import TokenItem from "~baseComponents/TokenItem";
import { useClient } from "~hooks/useClient";
import assets from "~src/tokens.json";
import useWallet from "../../utils/useWallet";

export interface IToken {
  name: string;
  display: string;
  amount?: number;
  denom: string;
  chain_name: string;
  logos?:
    | { svg: string; png: string }
    | { png: string; svg?: undefined }
    | { svg: string; png?: undefined }
    | undefined;
}

export interface TokenSelectorProps {
  onSave: (token: IToken | undefined) => void;
  selectedToken?: IToken;
  ownedToken?: boolean;
}

const TokenSelector = (props: TokenSelectorProps) => {
  const [selectedToken, setSelectedToken] = useState(props.selectedToken);
  const { onSave, ownedToken } = props;
  const { address } = useWallet();
  const [searchQuery, setSearchQuery] = useState("");
  const [tokens, setTokens] = useState<IToken[]>([]);

  const displayTokens = () => {
    // TODO: filter tokens by searchQuery
    let tokenList: IToken[] = tokens;

    if (searchQuery != "") {
      const properties = ["display", "name", "denom", "chain_name"];
      const filteredList: IToken[] = [];

      tokenList.forEach((token) => {
        let isValid = false;
        for (let i = 0; i < properties.length && !isValid; i++) {
          const value = token[properties[i]].toString();
          if (value.includes(searchQuery)) {
            isValid = true;
          }
        }

        if (isValid) filteredList.push(token);
      });

      tokenList = filteredList;
    }

    return tokenList.map((token, index) => {
      return (
        <TokenItem
          key={`token-selector-${index}`}
          token={token}
          onClick={(t) => onSave(t)}
          showAmount={true}
          selected={selectedToken?.denom == token.denom}
        />
      );
    });
  };

  const fetchOwnedToken = () => {
    if (address != "") {
      useClient()
        .CosmosBankV1Beta1.query.queryAllBalances(address)
        .then((response) => {
          const tokens: IToken[] = [];
          response.data.balances?.forEach((token) => {
            const t = assets.tokens.find((t) => t.denom === token.denom);
            if (t) {
              tokens.push({
                denom: t.denom,
                display: t.display,
                name: t.name,
                amount: Number(token.amount),
                chain_name: t.chain_name,
                logos: t.logos,
              });
            }
          });
          setTokens(tokens);
        });
    }
  };

  useEffect(() => {
    if (ownedToken) fetchOwnedToken();
    else setTokens(assets.tokens);
  }, []);

  return (
    <div className="modal">
      <div className="card">
        <div className="card-headers">
          Select a token
          <input
            className="search-bar"
            type="search"
            placeholder="Search Tokens"
            onChange={(e) => setSearchQuery(e.target.value)}
          />
        </div>
        <div className="card-body">{displayTokens()}</div>
      </div>
    </div>
  );
};

export default TokenSelector;
