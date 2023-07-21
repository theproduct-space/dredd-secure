import { useEffect, useState } from "react";
import { useAddressContext } from "~def-hooks/addressContext";
import useKeplr from "~def-hooks/useKeplr";
import { env } from "../../env";
import { useClient } from "~hooks/useClient";

export default function () {
  const keplr = useKeplr();
  const chainId = env.chainId;

  const { address } = useAddressContext();
  const [offlineSigner, setOfflineSigner] = useState(
    keplr.getOfflineSigner(chainId),
  );

  useEffect(() => {
    setOfflineSigner(keplr.getOfflineSigner(chainId));
  }, [address]);

  useEffect(() => {
    console.log("::::");
    const isConnected = sessionStorage.getItem("isConnected");

    if (isConnected) {
      const obj = JSON.parse(isConnected);

      if (obj["expiration"] >= Date.now()) {
        keplr.connectToKeplr(
          () => {
            sessionStorage.setItem(
              "isConnected",
              JSON.stringify({ value: true, expiration: new Date().getTime() }),
            );
          },
          () => {
            sessionStorage.removeItem("isConnected");
          },
        );
      } else {
        sessionStorage.removeItem("isConnected");
      }
    }
  }, []);

  return {
    offlineSigner,
    address,
  };
}
