// React Imports
import { Link, useNavigate } from "react-router-dom";
import Card from "~baseComponents/Card";
import Typography from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

export interface FailureSectionProps {
    errorTitle: string;
    errorBody: string;
    continueButton: JSX.Element;
}

const FailureSection = (props: FailureSectionProps) => {
    const { errorTitle, errorBody, continueButton } = props;
    const navigate = useNavigate();

    return (
        <ContentContainer className="grid grid-cols-1 grid-rows-3 content-center h-full lg:w-7/12 m-0">
            <span></span>
            <Card className="h-64">
                <div className="flex flex-col items-center text-center">
                    <Typography variant="h5" className="font-revalia pt-5">
                        {errorTitle}
                    </Typography>
                    <Typography variant="small" className="pt-7">
                        {errorBody}
                    </Typography>
                    <div className="flex lg:flex-row flex-col justify-center gap-3 lg:gap-10 pt-7 lg:pt-10">
                        <Link to="/dashboard">
                            <button
                                className={
                                    "transition hover:-translate-y-0.5 active:-translate-y-0.5 \
                                rounded uppercase flex gap-3 items-center border border-orange py-3 px-3 sm:py-3 sm:px-4"
                                }
                                onClick={() => navigate(-1)}>
                                <Typography variant="body-small" className="text-orange">
                                    Return to Escrow
                                </Typography>
                            </button>
                        </Link>
                        {continueButton}
                    </div>
                </div>
            </Card>
            <span></span>
        </ContentContainer>
    );
};

export default FailureSection;
