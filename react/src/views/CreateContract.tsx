import { useLocation } from "react-router-dom";
import CreateContract, { IContract } from "~sections/CreateContract";

const CreateContractPage = () => {
    const contract = useLocation().state as IContract | undefined;
    return <CreateContract contract={contract} />;
};

export default CreateContractPage;
