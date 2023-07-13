// React Imports
import { useState } from "react";
import { useLocation } from "react-router-dom";
import useKeplr from "~def-hooks/useKeplr";

// Custom Imports
import { IContract } from "~sections/CreateContract";
import PaymentSection from "~sections/PaymentSection";

const PaymentView = () => {
    const contract = useLocation().state as IContract;
    const keplr = useKeplr();
    const chainId = "dreddsecure";
    const [offlineSigner, setOfflineSigner] = useState(
        keplr.getOfflineSigner(chainId),
    );
    const [address, setAddress] = useState("");

    keplr.listenToAccChange(async () => {
        setOfflineSigner(keplr.getOfflineSigner(chainId));
        const { address } = (await offlineSigner.getAccounts())[0];
        setAddress(address);
    });

    return <PaymentSection contract={contract} wallet={{
        address: address,
        offlineSigner: offlineSigner
    }} />;
};

export default PaymentView;
