/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import HeroLanding from "~sections/HeroLanding";
import FAQ from "./FAQ";
import AboutUs from "~sections/AboutUs/AboutUs";

export default function Home() {
  return (
    <div className="w-full min-h-screen bg-black">
      <HeroLanding />
      {/* <AboutUs /> */}
      <FAQ />
    </div>
  );
}
