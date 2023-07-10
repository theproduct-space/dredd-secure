/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Link } from "react-router-dom";
import LayersImage from "~assets/layers-and-logo.png";
import Button from "~baseComponents/Button";
import LockIcon from "~icons/LockIcon";
import ContentContainer from "~layouts/ContentContainer";
import { Typography } from "~baseComponents/Typography";
import { URL_MEDIUM, URL_WAITLISTFORM } from "~utils/urls";

export default function HeroLanding() {
  return (
    <ContentContainer>
      <div className="relative min-h-screen w-full flex flex-col overflow-y-clip pt-20 lg:flex-row-reverse">
        <div className="flex flex-1 items-center justify-center">
          <img
            src={LayersImage}
            alt="Dredd-Secure"
            className="object-cover w-full max-w-lg lg:max-w-full relative drop-shadow-[0_35px_35px_rgba(255,199,90,0.25)]
          "
          />
        </div>
        <div className="flex-1 relative -top-12 left-0 w-full flex flex-col lg:justify-center items-start">
          <div className="max-w-2xl py-6">
            <Typography variant="h1" className="font-revalia pt-6 md:pb-3">
              DreddSecure
            </Typography>
            <Typography variant="h4" className="text-orange py-6 font-semibold">
              Secure Escrow Services for the Cosmos Ecosystem
            </Typography>
            <Typography variant="body" className="pb-3" as={"blockquote"}>
              Mitigating fraud and disputes by securely holding assets until
              pre-determined conditions are met. Facilitating smooth, safe, and
              confident transactions within the Cosmos Hub ecosystem.
            </Typography>
          </div>
          <div className="flex gap-3">
            <a href={URL_WAITLISTFORM} target="_blank">
              <Button
                text="get started"
                className="font-revalia rounded-full border-solid border"
              />
            </a>
            <a href={URL_MEDIUM} target="_blank">
              <Button
                text="learn more"
                secondary
                className="font-revalia rounded-full border-solid border"
              />
            </a>
          </div>
        </div>
      </div>
    </ContentContainer>
  );
}
