import { createContext, ReactNode, useContext } from "react";
import useWallet from "~src/components/utils/useWallet";

type NormalizedDenom = {
  normalized: string;
  path?: string;
  pathExtracted?: string[];
  isIBC: boolean;
};
type Denoms = Record<string, NormalizedDenom>;
interface Props {
  children?: ReactNode;
}
const denoms: Denoms = {};
const setDenom = (denom: string, meta: NormalizedDenom) => {
  if (!denoms[denom]) {
    denoms[denom] = meta;
  }
};
const context = { denoms, setDenom };
const DenomContext = createContext(context);
export const useDenomContext = () => useContext(DenomContext);

export default function DenomProvider({ children }: Props) {
  const { address } = useWallet();
  return (
    <DenomContext.Provider value={context}>{children}</DenomContext.Provider>
  );
}
