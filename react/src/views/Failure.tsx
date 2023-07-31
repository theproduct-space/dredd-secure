import FailureSection, { FailureSectionProps } from "~sections/FailureSection";

const Failure = (props: FailureSectionProps) => {
    return (
        <div
            className="h-screen w-100 bg-black bg-no-repeat bg-center bg-cover"
        >
            <FailureSection {...props} />
        </div>
    );
};

export default Failure;
