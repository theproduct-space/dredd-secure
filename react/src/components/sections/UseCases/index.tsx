import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import BlockPyramidImage from "~assets/block-pyramid.png";

const UseCases = () => {
  return (
    <ContentContainer className="py-20 md:py-32 relative min-h-screen flex items-center">
      <div className="relative w-full flex justify-end items-center">
        <img
          src={BlockPyramidImage}
          alt="Dredd-Secure"
          className="object-cover absolute left-0 z-0 top-44 scale-150 opacity-50 md:w-[50%] md:scale-100 md:left-0 md:top-0 md:opacity-100"
        />
        <div className="flex flex-col lg:max-w-[50%] z-10">
          <Typography variant="h2" className="font-revalia pt-6 pb-12">
            Diverse Use Cases
          </Typography>
          <Typography variant="h6" as={"blockquote"}>
            From peer-to-peer exchanges and token offerings to cross-chain asset
            swaps, decentralized marketplaces, and peer-to-peer lending,
            DreddSecure can be used in a wide range of scenarios to ensure
            secure and transparent transactions.
          </Typography>
        </div>
      </div>
    </ContentContainer>
  );
};

export default UseCases;
