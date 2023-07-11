/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import logoImage from "~assets/3d-logoNoBg.webp";

export default function AboutDredd() {
  return (
    <ContentContainer className="-mt-20 relative min-h-[1024px] h-screen flex items-start lg:-mt-8 lg:items-center">
      <div className="relative w-full flex justify-end items-center">
        <img
          src={logoImage}
          alt="Dredd-Secure"
          className="object-cover absolute left-4 z-0 top-64 scale-110 w-full max-w-[400px] opacity-50 drop-shadow-lightOrange 
            sm:left-16
            md:left-1/2 md:-translate-x-1/4
            lg:w-[60%] lg:max-w-[500px] lg:scale-100 lg:left-4 lg:top-4 lg:opacity-100 lg:translate-x-0
            xl:-top-20 xl:left-0 xl:max-w-[600px]"
          loading="lazy"
        />
        <div className="flex flex-col lg:max-w-[50%] z-10">
          <Typography variant="h2" className="font-revalia pt-6 pb-12">
            About DreddSecure
          </Typography>
          <Typography variant="h6" as={"blockquote"}>
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
