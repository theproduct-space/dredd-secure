// Custom Imports
import HeroLanding from "~sections/HeroLanding";
import FAQ from "~sections/FAQ";
import AboutDredd from "~sections/AboutDredd";
import WhyCosmos from "~sections/WhyCosmos";
import Features from "~sections/Features";
import Challenges from "~sections/Challenges";
import UseCases from "~sections/UseCases";
import CTASection from "~sections/CTASection";

const Home = () => {
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
};

export default Home;
