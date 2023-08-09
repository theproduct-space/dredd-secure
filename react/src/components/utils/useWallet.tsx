import { useEffect, useState } from "react";
import { useAddressContext } from "~def-hooks/addressContext";
import useKeplr from "~def-hooks/useKeplr";
import { env } from "../../env";

export default function () {
  const keplr = useKeplr();
  const chainId = env.chainId;

  const { address } = useAddressContext();

  useEffect(() => {
    const isConnected = sessionStorage.getItem("isConnected");
    if (isConnected && keplr.isKeplrAvailable) {
      const obj = JSON.parse(isConnected);

      if (obj["expiration"] >= Date.now()) {
        keplr.connectToKeplr(
          () => {
            sessionStorage.setItem(
              "isConnected",
              JSON.stringify({
                value: true,
                expiration: new Date().getTime() + 3600000 * 24,
              }),
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
    offlineSigner: keplr.isKeplrAvailable
      ? keplr.getOfflineSigner(chainId)
      : undefined,
    address,
  };
}
