// React Imports
import { Link } from "react-router-dom";
import Button from "~baseComponents/Button";
import Card from "~baseComponents/Card";
import Typography from "~baseComponents/Typography";
import DoubleCheck from "~icons/DoubleCheck";
import ContentContainer from "~layouts/ContentContainer";

export interface SuccessSectionProps {
  successTitle: string;
  successBody?: string;
  continueButton?: {
    buttonText: string;
    buttonLink: string;
  };
}

const SuccessSection = (props: SuccessSectionProps) => {
  const { successTitle, successBody, continueButton } = props;

  return (
    <ContentContainer className="pt-64 w-7/12">
      <Card>
        <div className="flex flex-col items-center text-center pt-10 pb-10">
          <div className="card-header">
            <Typography variant="h5" className="font-revalia text-orange">
              Complete
            </Typography>
          </div>
          <DoubleCheck />
          <Typography variant="h5" className="pt-4">
            {successTitle}
          </Typography>
          <Typography variant="small" className="pt-4 opacity-50">
            {successBody}
          </Typography>
          <div className="flex justify-center gap-10 pt-4">
            <Link to="/dashboard">
              <button
                className={
                  "transition hover:-translate-y-0.5 active:-translate-y-0.5 \
                                rounded uppercase flex gap-3 items-center border border-orange py-3 px-3 sm:py-3 sm:px-4"
                }
              >
                <Typography variant="body-small" className="text-orange">
                  Return Home
                </Typography>
              </button>
            </Link>
            {continueButton && (
              <Link to={continueButton.buttonLink}>
                <Button text={continueButton.buttonText} />
              </Link>
            )}
          </div>
        </div>
      </Card>
    </ContentContainer>
  );
};

export default SuccessSection;
