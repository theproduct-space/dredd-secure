// React Imports
import { useLocation } from "react-router-dom";

// Custom Imports
import { IContract } from "~sections/CreateContract";
import PaymentSection from "~sections/PaymentSection";

const PaymentView = () => {
  const contract = useLocation().state as IContract;
  return <PaymentSection contract={contract} />;
};

export default PaymentView;
