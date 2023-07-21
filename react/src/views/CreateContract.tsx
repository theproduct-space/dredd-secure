import { useLocation } from "react-router-dom";
import CreateContract, { IContract } from "~sections/CreateContract";

const CreateContractPage = () => {
  const contract = useLocation().state as IContract | undefined;
  return (
    <div className="bg-black">
      <CreateContract contract={contract} />;
    </div>
  );
};

export default CreateContractPage;
