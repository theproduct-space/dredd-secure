import axios from "axios";
import { useEffect, useState } from "react";
import BaseModal from "~baseComponents/BaseModal/Index";
import SearchBar from "~baseComponents/SearchBar";
import SecondaryButton from "~baseComponents/SecondaryButton";
import Typography from "~baseComponents/Typography";
import Close from "~icons/Close";

interface CoinGeckoTokenInfoProps {
  className?: string;
  selectedToken?: CoinGeckoToken;
  setSelectedToken: (token: CoinGeckoToken) => void;
}

interface CoinGeckoToken {
  id: string;
  symbol: string;
  name: string;
}

const CoinGeckoTokenSelector = ({
  className = "",
  selectedToken,
  setSelectedToken,
}: CoinGeckoTokenInfoProps) => {
  // const [searchValue, setSearchValue] = useState<string>("");
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  const [tokenList, setTokenList] = useState<CoinGeckoToken[]>([
    {
      id: "01coin",
      symbol: "zoc",
      name: "01coin",
    },
    {
      id: "0chain",
      symbol: "zcn",
      name: "Zus",
    },
    {
      id: "0vix-protocol",
      symbol: "vix",
      name: "0VIX Protocol",
    },
    {
      id: "0x",
      symbol: "zrx",
      name: "0x Protocol",
    },
    {
      id: "0x0-ai-ai-smart-contract",
      symbol: "0x0",
      name: "0x0.ai: AI Smart Contract",
    },
    {
      id: "0x1-tools-ai-multi-tool",
      symbol: "0x1",
      name: "0x1.tools: AI Multi-tool",
    },
    {
      id: "0xauto-io-contract-auto-deployer",
      symbol: "0xa",
      name: "0xAuto.io : Contract Auto Deployer",
    },
    {
      id: "0xcoco",
      symbol: "coco",
      name: "0xCoco",
    },
    {
      id: "0xdao",
      symbol: "oxd",
      name: "0xDAO",
    },
    {
      id: "0xdao-v2",
      symbol: "oxd v2",
      name: "0xDAO V2",
    },
    {
      id: "0xmonero",
      symbol: "0xmr",
      name: "0xMonero",
    },
    {
      id: "0xshield",
      symbol: "shield",
      name: "0xShield",
    },
    {
      id: "0xsniper",
      symbol: "0xs",
      name: "0xSniper",
    },
    {
      id: "12ships",
      symbol: "tshp",
      name: "12Ships",
    },
    {
      id: "14066-santa-rosa",
      symbol: "realt-s-14066-santa-rosa-dr-detroit-mi",
      name: "RealT - 14066 Santa Rosa Dr, Detroit, MI 48238",
    },
    {
      id: "1617-s-avers",
      symbol: "realt-s-1617-s.avers-ave-chicago-il",
      name: "RealT - 1617 S Avers Ave, Chicago, IL 60623",
    },
    {
      id: "1art",
      symbol: "1art",
      name: "OneArt",
    },
    {
      id: "1bch",
      symbol: "1bch",
      name: "1BCH",
    },
    {
      id: "1eco",
      symbol: "1eco",
      name: "1eco",
    },
    {
      id: "1hive-water",
      symbol: "water",
      name: "1Hive Water",
    },
    {
      id: "1inch",
      symbol: "1inch",
      name: "1inch",
    },
    {
      id: "1inch-yvault",
      symbol: "yv1inch",
      name: "1INCH yVault",
    },
    {
      id: "1million-nfts",
      symbol: "1mil",
      name: "1MillionNFTs",
    },
    {
      id: "1minbet",
      symbol: "1mb",
      name: "1minBET",
    },
    {
      id: "1move token",
      symbol: "1mt",
      name: "1Move Token",
    },
    {
      id: "1peco",
      symbol: "1peco",
      name: "1peco",
    },
    {
      id: "1reward-token",
      symbol: "1rt",
      name: "1Reward Token",
    },
    {
      id: "1safu",
      symbol: "safu",
      name: "1SAFU",
    },
    {
      id: "1sol",
      symbol: "1sol",
      name: "1Sol",
    },
    {
      id: "1sol-io-wormhole",
      symbol: "1sol",
      name: "1sol.io (Wormhole)",
    },
    {
      id: "20weth-80bal",
      symbol: "20weth-80bal",
      name: "20WETH-80BAL",
    },
    {
      id: "28vck",
      symbol: "vck",
      name: "28VCK",
    },
    {
      id: "2acoin",
      symbol: "arms",
      name: "2ACoin",
    },
    {
      id: "2crazynft",
      symbol: "2crz",
      name: "2crazyNFT",
    },
    {
      id: "2dai-io",
      symbol: "2dai",
      name: "2DAI.io",
    },
    {
      id: "2g-carbon-coin",
      symbol: "2gcc",
      name: "2G Carbon Coin",
    },
    {
      id: "2omb-finance",
      symbol: "2omb",
      name: "2omb",
    },
    {
      id: "2share",
      symbol: "2shares",
      name: "2SHARE",
    },
    {
      id: "300fit",
      symbol: "fit",
      name: "300FIT",
    },
    {
      id: "3d3d",
      symbol: "3d3d",
      name: "3d3d",
    },
    {
      id: "3-kingdoms-multiverse",
      symbol: "3km",
      name: "3 Kingdoms Multiverse",
    },
    {
      id: "3shares",
      symbol: "3share",
      name: "3Share",
    },
    {
      id: "3xcalibur",
      symbol: "xcal",
      name: "3xcalibur Ecosystem Token",
    },
    {
      id: "42-coin",
      symbol: "42",
      name: "42-coin",
    },
    {
      id: "4852-4854-w-cortez",
      symbol: "realt-s-4852-4854-w.cortez-st-chicago-il",
      name: "RealT - 4852-4854 W Cortez St, Chicago, IL 60651",
    },
    {
      id: "4artechnologies",
      symbol: "4art",
      name: "4ART Coin",
    },
    {
      id: "zfmcoin",
      symbol: "zfm",
      name: "ZFMCOIN",
    },
    {
      id: "zhc-zero-hour-cash",
      symbol: "zhc",
      name: "ZHC : Zero Hour Cash",
    },
    {
      id: "zibu",
      symbol: "zibu",
      name: "Zibu",
    },
    {
      id: "ziesha",
      symbol: "zsh",
      name: "Ziesha",
    },
    {
      id: "zignaly",
      symbol: "zig",
      name: "Zignaly",
    },
    {
      id: "zigzag-2",
      symbol: "zz",
      name: "ZigZag",
    },
    {
      id: "zik-token",
      symbol: "zik",
      name: "Ziktalk",
    },
    {
      id: "zillion-aakar-xo",
      symbol: "zillionxo",
      name: "Zillion Aakar XO",
    },
    {
      id: "zilliqa",
      symbol: "zil",
      name: "Zilliqa",
    },
    {
      id: "zilpay-wallet",
      symbol: "zlp",
      name: "ZilPay Wallet",
    },
    {
      id: "zilpepe",
      symbol: "zilpepe",
      name: "ZilPepe",
    },
    {
      id: "zilstream",
      symbol: "stream",
      name: "ZilStream",
    },
    {
      id: "zilswap",
      symbol: "zwap",
      name: "ZilSwap",
    },
    {
      id: "zimbocash",
      symbol: "zash",
      name: "ZIMBOCASH",
    },
    {
      id: "zin",
      symbol: "zin",
      name: "Zin",
    },
    {
      id: "zion",
      symbol: "zion",
      name: "Zion",
    },
    {
      id: "zion-token",
      symbol: "zion",
      name: "Zion Token",
    },
    {
      id: "ziontopia",
      symbol: "zion",
      name: "ZionTopia",
    },
    {
      id: "zip",
      symbol: "zip",
      name: "Zipper Network",
    },
    {
      id: "zipmex-token",
      symbol: "zmt",
      name: "Zipmex",
    },
    {
      id: "zipswap",
      symbol: "zip",
      name: "ZipSwap",
    },
    {
      id: "zizy",
      symbol: "zizy",
      name: "Zizy",
    },
    {
      id: "zjoe",
      symbol: "zjoe",
      name: "zJOE",
    },
    {
      id: "zkapes-token",
      symbol: "zat",
      name: "zkApes Token",
    },
    {
      id: "zkcult",
      symbol: "zcult",
      name: "zkCULT",
    },
    {
      id: "zkdoge",
      symbol: "zkdoge",
      name: "zkDoge",
    },
    {
      id: "zkfloki",
      symbol: "zfloki",
      name: "zkFloki",
    },
    {
      id: "zk-inu",
      symbol: "$zkinu",
      name: "ZK inu",
    },
    {
      id: "zklaunchpad",
      symbol: "zkpad",
      name: "zkLaunchpad",
    },
    {
      id: "zklotto",
      symbol: "zklotto",
      name: "zkLotto",
    },
    {
      id: "zknftex",
      symbol: "$zkn",
      name: "zkNFTex",
    },
    {
      id: "zkpepe",
      symbol: "zkpepe",
      name: "ZKPepe",
    },
    {
      id: "zkproof",
      symbol: "zkp",
      name: "zkProof",
    },
    {
      id: "zkshib",
      symbol: "zkshib",
      name: "zkShib",
    },
    {
      id: "zkspace",
      symbol: "zks",
      name: "ZKSpace",
    },
    {
      id: "zksvm",
      symbol: "zksvm",
      name: "zkSVM",
    },
    {
      id: "zkswap-92fc4897-ea4c-4692-afc9-a9840a85b4f2",
      symbol: "zksp",
      name: "zkSwap",
    },
    {
      id: "zksync-id",
      symbol: "zkid",
      name: "zkSync id",
    },
    {
      id: "zksync-labs",
      symbol: "zklab",
      name: "zkSync Labs",
    },
    {
      id: "zktsunami",
      symbol: ":zkt:",
      name: "ZkTsunami",
    },
    {
      id: "zkvault",
      symbol: "zkvault",
      name: "zkVAULT",
    },
    {
      id: "zmine",
      symbol: "zmn",
      name: "ZMINE",
    },
    {
      id: "zodiacsv2",
      symbol: "zdcv2",
      name: "ZodiacsV2",
    },
    {
      id: "zodium",
      symbol: "zodi",
      name: "Zodium",
    },
    {
      id: "zoid-pay",
      symbol: "zpay",
      name: "ZoidPay",
    },
    {
      id: "zombie-inu",
      symbol: "zinu",
      name: "Zombie Inu (OLD)",
    },
    {
      id: "zombie-inu-2",
      symbol: "zinu",
      name: "Zombie Inu",
    },
    {
      id: "zone",
      symbol: "zone",
      name: "Zone",
    },
    {
      id: "zone-of-avoidance",
      symbol: "zoa",
      name: "Zone of Avoidance",
    },
    {
      id: "zoocoin",
      symbol: "zoo",
      name: "ZooCoin",
    },
    {
      id: "zoo-coin",
      symbol: "zoo",
      name: "ZooCoin",
    },
    {
      id: "zoo-crypto-world",
      symbol: "zoo",
      name: "ZOO Crypto World",
    },
    {
      id: "zoodao",
      symbol: "zoo",
      name: "ZooDAO",
    },
    {
      id: "zookeeper",
      symbol: "zoo",
      name: "ZooKeeper",
    },
    {
      id: "zoomer",
      symbol: "zoomer",
      name: "Zoomer",
    },
    {
      id: "zoomswap",
      symbol: "zm",
      name: "ZoomSwap",
    },
    {
      id: "zoo-token",
      symbol: "zoot",
      name: "Zoo",
    },
    {
      id: "zoracles",
      symbol: "zora",
      name: "Zoracles",
    },
    {
      id: "zro",
      symbol: "zro",
      name: "Carb0n.fi",
    },
    {
      id: "zsol",
      symbol: "zsol",
      name: "zSOL",
    },
    {
      id: "zuki-moba",
      symbol: "zuki",
      name: "Zuki Moba",
    },
    {
      id: "zum-token",
      symbol: "zum",
      name: "ZUM",
    },
    {
      id: "zuna",
      symbol: "zuna",
      name: "Zuna",
    },
    {
      id: "zunami-eth",
      symbol: "zeth",
      name: "Zunami Ether",
    },
    {
      id: "zunami-protocol",
      symbol: "uzd",
      name: "Zunami USD",
    },
    {
      id: "zurrency",
      symbol: "zurr",
      name: "ZURRENCY",
    },
    {
      id: "zusd",
      symbol: "zusd",
      name: "ZUSD",
    },
    {
      id: "zyberswap",
      symbol: "zyb",
      name: "Zyberswap",
    },
    {
      id: "zynecoin",
      symbol: "zyn",
      name: "Zynecoin",
    },
    {
      id: "zynergy",
      symbol: "zyn",
      name: "Zynergy",
    },
    {
      id: "zyrri",
      symbol: "zyr",
      name: "Zyrri",
    },
    {
      id: "zyx",
      symbol: "zyx",
      name: "ZYX",
    },
    {
      id: "zzz",
      symbol: "zzz",
      name: "GoSleep ZZZ",
    },
  ]); // TODO: initilise to empty array and make the api call in useEffect
  const [filteredTokenList, setFilteredTokenList] =
    useState<CoinGeckoToken[]>(tokenList);
  useEffect(() => {
    // fetchTokenList(); // TODO reactivate the api call -- might need a proxy server
  }, []);

  const fetchTokenList = async () => {
    try {
      const res = await axios.get(
        "https://api.coingecko.com/api/v3/coins/list",
      );
      console.log("res", res);
      setTokenList(res.data);
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
        key={`coinGecko-token-selector-${token.id}`}
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
        <button className="flex gap-5 justify-between border border-white-200 p-4 bg-buttonBg rounded-lg">
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

export default CoinGeckoTokenSelector;
