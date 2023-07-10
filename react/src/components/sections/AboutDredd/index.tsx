/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

import logoImage from "~assets/3d-logoNoBg.png";

export default function AboutDredd() {
  return (
    <ContentContainer>
      <div className="relative min-h-screen w-full overflow-y-clip flex flex-col justify-center items-center gap-44 lg:flex-row">
        <img
          src={logoImage}
          alt="Dredd-Secure"
          className="object-cover z-0 w-[35%] drop-shadow-yellow"
        />
        <div className="flex flex-col z-10 lg:max-w-[50%]">
          <Typography
            variant="h3"
            className="text-white-1000 font-revalia py-6"
          >
            About DreddSecure
          </Typography>
          <Typography
            variant="h6"
            as={"blockquote"}
            className="text-white-1000"
          >
            DreddSecure is an advanced blockchain-based escrow app for the
            Cosmos Hub ecosystem, designed by ProductShop, a web development
            studio based in Montréal, Canada. We're committed to enhancing
            secure, reliable, and scalable blockchain transactions. With
            DreddSecure, we're not just adapting to the digital revolution,
            we're actively shaping it, creating a safer digital transaction
            landscape.
          </Typography>
        </div>
      </div>
    </ContentContainer>
  );
}
