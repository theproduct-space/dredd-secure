/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

import bgImage from "~assets/bg2.png";

export default function WhyCosmos() {
  return (
    <section className="bg-black">
      <ContentContainer>
        <div className="relative h-screen w-full overflow-y-clip flex justify-start items-center">
          {/* <img
            src={bgImage}
            alt="Dredd-Secure"
            className="object-contain absolute opacity-70 
          bottom-0 scale-[3]
          sm:scale-[2]
          md:-left-[50%] md:scale-125 md:top-unset
          xl:opacity-100"
          /> */}
          <div className="flex flex-col lg:max-w-[50%]">
            <Typography variant="h2" className="text-white-1000 font-revalia py-6">
              Why Cosmos?
            </Typography>
            <Typography variant="h6" as={"blockquote"} className="text-white-1000">
              Cosmos offers unique advantages for DreddSecure. Its Inter-Blockchain Communication protocol allows
              seamless asset and data transfers across different chains, providing broad usability. Cosmos' Hub design
              ensures efficient handling of numerous transactions, facilitating scalability. The modular Cosmos SDK
              simplifies blockchain and application development, allowing for complex smart contract creation. All these
              attributes make Cosmos an optimal choice for a secure, scalable, and interoperable escrow app.
            </Typography>
          </div>
        </div>
      </ContentContainer>
    </section>
  );
}
