// Custom Imports
import Typography from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

// Assets Imports
import BlockPyramidImage from "~assets/block-pyramid.webp";

const UseCases = () => {
  return (
    <ContentContainer className="py-20 relative min-h-[750px] h-screen flex items-start md:py-[20rem]">
      <div className="relative w-full flex justify-end items-center">
        <img
          src={BlockPyramidImage}
          alt="Dredd-Secure"
          className="object-cover absolute left-0 z-0 top-44 scale-125 max-w-[400px] opacity-50 drop-shadow-lightOrange
            lg:w-[60%] lg:max-w-[700px] lg:scale-100 lg:-left-20 lg:-top-20 lg:opacity-100
            xl:-top-44 xl:left-0"
          loading="lazy"
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
