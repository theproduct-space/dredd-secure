/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

import CubesImage from "~assets/random-cubes.png";

export default function WhyCosmos() {
  return (
    <ContentContainer className="mt-20">
      <div className="relative min-h-screen w-full overflow-y-clip flex justify-start items-center pb-[20rem] sm:pb-[22rem] md:pb-[32rem] lg:pt-[4rem] lg:pb-[8rem]">
        <div className="flex flex-col lg:max-w-[50%] z-10">
          <Typography variant="h2" className="font-revalia py-6">
            Why Cosmos?
          </Typography>
          <Typography variant="h6" as={"blockquote"}>
            Cosmos offers unique advantages for DreddSecure. Its
            Inter-Blockchain Communication protocol allows seamless asset and
            data transfers across different chains, providing broad usability.
            Cosmos' Hub design ensures efficient handling of numerous
            transactions, facilitating scalability. The modular Cosmos SDK
            simplifies blockchain and application development, allowing for
            complex smart contract creation. All these attributes make Cosmos an
            optimal choice for a secure, scalable, and interoperable escrow app.
          </Typography>
        </div>
        <img
          src={CubesImage}
          alt="Dredd-Secure"
          className="object-cover absolute right-0 z-0 drop-shadow-orange opacity-40 mt-[400px] lg:mt-0 lg:opacity-100 lg:w-[42%] lg:mt-0"
        />
      </div>
    </ContentContainer>
  );
}
