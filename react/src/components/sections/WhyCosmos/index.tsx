/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

import CubesImage from "~assets/random-cubes.png";

export default function WhyCosmos() {
  return (
    <section className="bg-black">
      <ContentContainer>
        <div className="relative h-screen w-full overflow-y-clip flex justify-start items-center">
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
              complex smart contract creation. All these attributes make Cosmos
              an optimal choice for a secure, scalable, and interoperable escrow
              app.
            </Typography>
          </div>
          <img
            src={CubesImage}
            alt="Dredd-Secure"
            className="object-cover absolute right-0 z-0 w-[50%]"
          />
        </div>
      </ContentContainer>
    </section>
  );
}
