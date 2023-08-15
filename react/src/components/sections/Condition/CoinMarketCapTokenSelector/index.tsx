import axios from "axios";
import { useEffect, useState } from "react";
import BaseModal from "~baseComponents/BaseModal/Index";
import SearchBar from "~baseComponents/SearchBar";
import SecondaryButton from "~baseComponents/SecondaryButton";
import Typography from "~baseComponents/Typography";
import Close from "~icons/Close";

interface CoinMarketCapTokenInfoProps {
  className?: string;
  selectedToken?: CoinMarketCapTokenI;
  setSelectedToken: (token: CoinMarketCapTokenI) => void;
}

export interface CoinMarketCapTokenI {
  id: number;
  name: string;
  symbol: string;
}

const CoinMarketCapTokenSelector = ({
  className = "",
  selectedToken,
  setSelectedToken,
}: CoinMarketCapTokenInfoProps) => {
  // const [searchValue, setSearchValue] = useState<string>("");
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  const [tokenList, setTokenList] = useState<CoinMarketCapTokenI[]>([
    {
      id: 1,
      name: "Bitcoin",
      symbol: "BTC",
    },
    {
      id: 2,
      name: "Litecoin",
      symbol: "LTC",
    },
    {
      id: 3,
      name: "Namecoin",
      symbol: "NMC",
    },
    {
      id: 4,
      name: "Terracoin",
      symbol: "TRC",
    },
    {
      id: 5,
      name: "Peercoin",
      symbol: "PPC",
    },
    {
      id: 6,
      name: "Novacoin",
      symbol: "NVC",
    },
  ]); // TODO: initilise to empty array and make the api call in useEffect
  const [filteredTokenList, setFilteredTokenList] =
    useState<CoinMarketCapTokenI[]>(tokenList);
  useEffect(() => {
    // fetchTokenList(); // TODO reactivate the api call -- might need a proxy server
  }, []);

  const fetchTokenList = async () => {
    try {
      axios.defaults.headers.post["Access-Control-Allow-Origin"] = "*";
      const res = await axios.get(
        "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map",
        {
          headers: {
            "X-CMC_PRO_API_KEY": "0c27f13f-c6fe-45f8-8829-2d82404d7ef9", // use proxy server to hide api key
          },
        },
      );
      console.log("res cmc token list", res);
      const formattedTokenList = res.data.map((token) => ({
        id: token.id,
        symbol: token.symbol,
        name: token.name,
      }));
      console.log("res cmc formattedTokenList", formattedTokenList);

      setTokenList(formattedTokenList);
      setFilteredTokenList(res.data);
    } catch (error) {
      console.log("error", error);
    }
  };

  const handleFiltering = (searchValue) => {
    const result = tokenList.filter((token) => {
      if (searchValue === "") return tokenList;

      return (
        token.name.toLowerCase().includes(searchValue.toLowerCase()) ||
        token.symbol.toLowerCase().includes(searchValue.toLowerCase())
      );
    });

    setFilteredTokenList(result);
  };

  const displayToken = () => {
    return filteredTokenList.map((token) => (
      <button
        className="flex gap-3 justify-between py-3 px-3 rounded hover:bg-white-200"
        key={`coinMarketCap-token-selector-${token.id}`}
        onClick={() => {
          setSelectedToken(token);
          setIsModalOpen(false);
        }}
      >
        <Typography variant="body-small" className="truncate max-w-[40%]">
          {token.name}
        </Typography>
        <Typography variant="body-small" className="truncate max-w-[40%]">
          {token.symbol.toUpperCase()}
        </Typography>
      </button>
    ));
  };

  return (
    <>
      <Typography variant="body-small" className="text-white-500 p-1">
        Token of Interest
      </Typography>
      {selectedToken ? (
        <button
          className="flex gap-5 justify-between border border-white-200 p-4 bg-buttonBg rounded-lg"
          onClick={() => setIsModalOpen(true)}
        >
          <Typography variant="body-small">{selectedToken?.name}</Typography>
          <Typography variant="body-small">
            {selectedToken?.symbol?.toUpperCase()}
          </Typography>
        </button>
      ) : (
        <div className={`${className}`}>
          <SecondaryButton
            text="Select Token"
            orangeText
            onClick={() => setIsModalOpen(true)}
          />
          {/* <Typography variant="body-small" className="text-white-500 p-1">
          Select Token
        </Typography>
        <TextField
          type="text"
          variant="outlined"
          placeholder="Value"
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
        /> */}
        </div>
      )}

      <BaseModal open={isModalOpen} handleClose={() => setIsModalOpen(false)}>
        <div className="relative">
          <div className="sticky top-0 w-full border-b-[1px] border-white-200 p-4 bg-gray">
            <div className="flex justify-between items-center pb-4">
              <Typography variant="body">Select a token</Typography>
              <Close
                onClick={() => setIsModalOpen(false)}
                className="hover: cursor-pointer w-6 h-6"
              />
            </div>
            <SearchBar
              placeholder="Search Tokens"
              onChange={(searchValue) => handleFiltering(searchValue)}
            />
          </div>
          <div className="flex flex-col gap-3 p-4 max-w-full">
            {displayToken()}
          </div>
        </div>
      </BaseModal>
    </>
  );
};

export default CoinMarketCapTokenSelector;
