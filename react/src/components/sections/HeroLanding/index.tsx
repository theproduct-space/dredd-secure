/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Link } from "react-router-dom";
import bgImage from "~assets/bg3.png";
import Button from "~baseComponents/Button";
export default function HeroLanding() {
  return (
    <div className="relative h-[1024px]">
      <img src={bgImage} alt="Dredd-Secure" className="w-full absolute bottom-0 left-96" />
      {/* <div className="absolute w-[]" /> */}
      <div className="absolute top-0 left-0 w-full h-full flex flex-col justify-center">
        <div className="container mx-auto">
          <div className="max-w-2xl py-6">
            <h1 className="text-h1 text-white-1000 font-revalia py-6">Dredd-Secure</h1>
            <blockquote className="text-white-1000 text-p1">
              Discover our innovative escrow service designed for the Cosmos Hub ecosystem. Using the powerful Cosmos
              SDK, we bring enhanced security and efficiency to transactions within the ecosystem.
            </blockquote>
          </div>
          <Link to="/app">
            <Button text="Lauch App" className="font-revalia rounded-full border-solid border-2" />
          </Link>
        </div>
      </div>
    </div>
  );
}
