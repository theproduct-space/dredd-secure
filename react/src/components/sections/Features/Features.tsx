/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

import bgImage from "~assets/bg2.png";

export default function Features() {
  return (
    <section className="bg-black">
      <ContentContainer>
        <div className="relative h-screen w-full overflow-y-clip flex justify-end items-center">
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
              Key Features of DreddSecure
            </Typography>
            <div>
              <div className="">
                <Typography variant="h5" as={"blockquote"} className="text-white-1000">
                  Cosmos SDK Module
                </Typography>
                <Typography variant="body-small" as={"blockquote"} className="text-white-1000">
                  Seamless integration with the Cosmos Hub and other chains in the Cosmos ecosystem.
                </Typography>
              </div>
              <div className="">
                <Typography variant="h5" as={"blockquote"} className="text-white-1000">
                  Escrow Contract
                </Typography>
                <Typography variant="body-small" as={"blockquote"} className="text-white-1000">
                  A configurable set of rules that holds the assets in escrow and releases them only when the specified
                  conditions are met.
                </Typography>
              </div>
              <div className="">
                <Typography variant="h5" as={"blockquote"} className="text-white-1000">
                  API & UI
                </Typography>
                <Typography variant="body-small" as={"blockquote"} className="text-white-1000">
                  A user-friendly web application and a set of APIs for easy creation, management, and querying of
                  escrow transactions.
                </Typography>
              </div>
              <div className="">
                <Typography variant="h5" as={"blockquote"} className="text-white-1000">
                  Oracle Integration
                </Typography>
                <Typography variant="body-small" as={"blockquote"} className="text-white-1000">
                  Built on the Cosmos SDK, our service follows the latest security best practices, keeping your assets
                  safe.
                </Typography>
              </div>
              <div className="">
                <Typography variant="h5" as={"blockquote"} className="text-white-1000">
                  Scalability
                </Typography>
                <Typography variant="body-small" as={"blockquote"} className="text-white-1000">
                  Designed to handle a large number of concurrent transactions, allowing it to scale horizontally
                  without performance degradation.
                </Typography>
              </div>
              <div className="">
                <Typography variant="h5" as={"blockquote"} className="text-white-1000">
                  Cosmos SDK Module
                </Typography>
                <Typography variant="body-small" as={"blockquote"} className="text-white-1000">
                  Seamless integration with the Cosmos Hub and other chains in the Cosmos ecosystem.
                </Typography>
              </div>
            </div>
          </div>
        </div>
      </ContentContainer>
    </section>
  );
}
