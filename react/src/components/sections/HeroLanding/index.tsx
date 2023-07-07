/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Link } from "react-router-dom";
import bgImage from "~assets/bg3.png";
import Button from "~baseComponents/Button";
import LockIcon from "~icons/LockIcon";
import ContentContainer from "~layouts/ContentContainer";
import { Typography } from "~baseComponents/Typography";

export default function HeroLanding() {
  return (
    <ContentContainer>
      <div className="relative h-screen w-full overflow-y-clip">
        <img
          src={bgImage}
          alt="Dredd-Secure"
          className="object-contain absolute opacity-70 
          bottom-0 scale-[3]
          sm:scale-[2]
          md:-right-[200px] md:scale-150 md:top-unset
          xl:opacity-100"
        />
        <div className="absolute top-0 left-0 w-full h-full flex flex-col py-40 lg:justify-center items-start">
          <div className="">
            <div className="max-w-2xl py-6">
              <Typography
                variant="h1"
                className="text-white-1000 font-revalia py-6"
              >
                DreddSecure
              </Typography>
              <Typography
                variant="body"
                className="text-white-1000 text-p1"
                as={"blockquote"}
              >
                Discover our innovative escrow service designed for the Cosmos
                Hub ecosystem. Using the powerful Cosmos SDK, we bring enhanced
                security and efficiency to transactions within the ecosystem.
              </Typography>
            </div>
            <div className="flex gap-3">
              {/* <Link to="/app"> */}
              <Button
                text="lauch app"
                icon={<LockIcon />}
                disabled
                className="font-revalia rounded-full border-solid border"
              />
              {/* </Link> */}
              {/* <Link to="/app"> */}
              <Button
                text="join waitlist"
                icon={<LockIcon />}
                secondary
                disabled
                className="font-revalia rounded-full border-solid border"
              />
              {/* </Link> */}
            </div>
          </div>
        </div>
      </div>
    </ContentContainer>
  );
}
