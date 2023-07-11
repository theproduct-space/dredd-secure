/* eslint-disable import/no-unresolved */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import LayersImage from "~assets/layers.webp";

const Challenges = () => {
  return (
    <ContentContainer className="py-20 pb-32 relative min-h-screen flex items-center md:py-72">
      <div className="relative w-full flex justify-start items-center">
        <div className="flex flex-col lg:max-w-[50%] z-10">
          <Typography variant="h2" className="font-revalia pt-6 pb-12">
            Overcoming Challenges with DreddSecure
          </Typography>
          <Typography variant="h6" as={"blockquote"}>
            <strong>Problem: </strong>The absence of an escrow service can lead
            to a lack of trust, inadequate security, and insufficient protection
            for buyers and sellers.
          </Typography>
          <br />
          <Typography variant="h6" as={"blockquote"}>
            <strong>Solution:</strong> DreddSecure provides a two-party Escrow
            Contract, asset escrow, flexible conditions, a user interface, APIs
            and integration to overcome these challenges.
          </Typography>
        </div>
        <img
          src={LayersImage}
          alt="Dredd-Secure"
          className="object-cover drop-shadow-lightOrange absolute z-0 -right-10 top-[7rem] opacity-30 max-w-[500px] w-full
          lg:w-[50%] lg:right-0 lg:top-auto lg:bottom-auto lg:opacity-100 lg:max-w-full"
          loading="lazy"
        />
      </div>
    </ContentContainer>
  );
};

export default Challenges;
