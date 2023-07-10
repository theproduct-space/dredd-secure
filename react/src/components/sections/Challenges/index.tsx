import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import LayersImage from "~assets/layers.png";

const Challenges = () => {
  return (
    <ContentContainer className="py-20 md:py-32 relative">
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
      </div>
      <img
        src={LayersImage}
        alt="Dredd-Secure"
        className="object-cover absolute z-0 scale-150 -right-20 top-44 opacity-50 
          md:w-[50%] md:scale-100 md:right-0 md:top-0 "
      />
    </ContentContainer>
  );
};

export default Challenges;
