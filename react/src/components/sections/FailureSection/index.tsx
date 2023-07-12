// React Imports
import { useNavigate } from "react-router-dom";

export interface FailureSectionProps {
  errorTitle: string;
  errorBody: string;
  continueButton: JSX.Element;
}

const FailureSection = (props: FailureSectionProps) => {
  const { errorTitle, errorBody, continueButton } = props;
  const navigate = useNavigate();

  return (
    <div className="card">
      <div className="card-title centered">{errorTitle}</div>
      <div className="card-body">{errorBody}</div>
      <div>
        <button onClick={() => navigate(-1)}>Return to Escrow</button>
        {continueButton}
      </div>
    </div>
  );
};

export default FailureSection;
