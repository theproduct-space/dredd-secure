/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import HeroLanding from "~sections/HeroLanding";
import FAQ from "./FAQ";
import AboutUs from "~sections/AboutUs";
import AboutDredd from "~sections/AboutDredd";
import WhyCosmos from "~sections/WhyCosmos";
import Features from "~sections/Features";

export default function Home() {
  return (
    <div className="w-full min-h-screen bg-black">
      <HeroLanding />
      <AboutDredd />
      <WhyCosmos />
      <Features />
      {/* <AboutUs /> */}
      <FAQ />
    </div>
  );
}
