import SuccessSection, { SuccessSectionProps } from "~sections/SuccessSection";
import LinkedBalls from "~assets/linked-balls.png";

const Success = (props: SuccessSectionProps) => {
  return (
    <div
      className="h-screen w-100 bg-no-repeat bg-center bg-cover"
      style={{ backgroundImage: `url(${LinkedBalls})` }}
    >
      <SuccessSection {...props} />
    </div>
  );
};

export default Success;
