import Typography from "~baseComponents/Typography";

interface FeatureProps {
  icon: JSX.Element;
  title: string;
  description: string;
}

const FeatureCard = ({ icon, title, description }: FeatureProps) => (
  <div className="flex flex-col justify-center items-center">
    <div className="drop-shadow-orangeCenter">{icon}</div>
    <Typography variant="h5" className="pt-10 text-center">
      {title}
    </Typography>
    <Typography
      variant="body-small"
      as={"blockquote"}
      className="px-10 py-6 text-center lg:p-6"
    >
      {description}
    </Typography>
  </div>
);

export default FeatureCard;
