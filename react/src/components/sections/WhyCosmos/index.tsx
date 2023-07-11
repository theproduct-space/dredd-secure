/* eslint-disable react/no-unescaped-entities */
/* eslint-disable import/no-unresolved */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

import CubesImage from "~assets/random-cubes.webp";

export default function WhyCosmos() {
  return (
    <ContentContainer className="-mt-16 md:mt-0">
      <div className="relative min-h-screen w-full overflow-y-clip flex justify-start items-center pb-[20rem] sm:pb-[35rem] md:pb-[24rem] lg:pb-[4rem]">
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
          className="object-cover absolute right-0 z-0 drop-shadow-orange opacity-40 mt-[400px] md:max-w-[75%] lg:mt-0 lg:opacity-100 lg:w-[42%] lg:mt-0"
          loading="lazy"
        />
      </div>
    </ContentContainer>
  );
}
