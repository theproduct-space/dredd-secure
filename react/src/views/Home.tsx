/* eslint-disable import/no-unresolved */
import HeroLanding from "~sections/HeroLanding";
import FAQ from "./FAQ";
import AboutUs from "~sections/AboutUs";
import AboutDredd from "~sections/AboutDredd";
import WhyCosmos from "~sections/WhyCosmos";
import Features from "~sections/Features";

import Challenges from "~sections/Challenges";
import UseCases from "~sections/UseCases";
import CTASection from "~sections/CTASection";

export default function Home() {
  return (
    <div className="w-full min-h-screen bg-black">
      <HeroLanding />
      <AboutDredd />
      <WhyCosmos />
      <Features />
      <Challenges />
      <UseCases />
      <FAQ />
      <CTASection />
    </div>
  );
}
