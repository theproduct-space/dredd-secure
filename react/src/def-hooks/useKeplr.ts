import { useClient } from "../hooks/useClient";
import { useDispatchWalletContext } from "../def-hooks/walletContext";

export default function () {
  const client = useClient();
  const walletStore = useDispatchWalletContext();

  const connectToKeplr = async (
    onSuccessCb: () => void,
    onErrorCb: () => void,
  ) => {
    try {
      walletStore.connectWithKeplr();
      onSuccessCb();
    } catch (e) {
      console.error(e);
      onErrorCb();
    }
  };

  const isKeplrAvailable = !!window.keplr;

  const getOfflineSigner = (chainId: string) =>
    window.keplr.getOfflineSigner(chainId);

  const getKeplrAccParams = async (chainId: string) => {
    if (isKeplrAvailable)
        return window.keplr.getKey(chainId);
    else
        return new Promise<{ name: string, bech32Address: string }>((resolve) => {
            resolve({ name: "", bech32Address: "" })
        });
  }

  const listenToAccChange = (cb: EventListener) => {
    client.on("signer-changed", cb);
  };

  return {
    connectToKeplr,
    isKeplrAvailable,
    getOfflineSigner,
    getKeplrAccParams,
    listenToAccChange,
  };
}
