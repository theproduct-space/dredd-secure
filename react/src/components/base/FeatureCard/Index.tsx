import React from "react";
import { Typography } from "~baseComponents/Typography";

interface Feature {
  icon: JSX.Element;
  title: string;
  description: string;
}

const FeatureCard: React.FC<Feature> = ({ icon, title, description }) => (
  <div className="flex flex-col justify-center items-center">
    <div className="drop-shadow-orangeCenter">{icon}</div>
    <Typography variant="h5" className="text-white-1000 pt-10 text-center">
      {title}
    </Typography>
    <Typography
      variant="body-small"
      as={"blockquote"}
      className="px-10 py-6 text-white-1000 text-center lg:p-6"
    >
      {description}
    </Typography>
  </div>
);

export default FeatureCard;
