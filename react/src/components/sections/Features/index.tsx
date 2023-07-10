/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import CosmosSDK from "~icons/CosmosSDK";
import Integration from "~icons/Integration";
import Reliable from "~icons/Reliable";
import Transaction from "~icons/Transaction";
import Decentralized from "~icons/Decentralized";
import Security from "~icons/Security";
import FeatureCard from "~baseComponents/FeatureCard/Index";

interface Feature {
  icon: JSX.Element;
  title: string;
  description: string;
}

const featuresData: Feature[] = [
  {
    icon: <CosmosSDK />,
    title: "Cosmos SDK Module",
    description: "Seamless integration with the Cosmos Hub and other chains in the Cosmos ecosystem.",
  },
  {
    icon: <Transaction />,
    title: "Escrow Contract",
    description:
      "A configurable set of rules that holds the assets in escrow and releases them only when the specified conditions are met.",
  },
  {
    icon: <Integration />,
    title: "API & UI",
    description:
      "A user-friendly web application and a set of APIs for easy creation, management, and querying of escrow transactions.",
  },
  {
    icon: <Decentralized />,
    title: "Oracle Integration",
    description:
      "Integration with oracles to allow the use of external data sources for verifying transaction conditions.",
  },
  {
    icon: <Security />,
    title: "Security",
    description:
      "Built on the Cosmos SDK, our service follows the latest security best practices, keeping your assets safe.",
  },
  {
    icon: <Reliable />,
    title: "Scalability",
    description:
      "Designed to handle a large number of concurrent transactions, allowing it to scale horizontally without performance degradation.",
  },
];

export default function Features() {
  return (
    <section className="bg-black">
      <ContentContainer>
        <div className="relative min-h-screen w-full overflow-y-clip flex justify-center items-center">
          <div className="flex flex-col">
            <Typography variant="h2" className="py-16 text-white-1000 font-revalia md:py-32">
              Key Features of DreddSecure
            </Typography>
            <div className="grid gap-4 grid-cols-1 md:grid-cols-2 grid-rows-3 lg:grid-cols-3 grid-rows-2">
              {featuresData.map((feature, index) => (
                <FeatureCard key={index} {...feature} />
              ))}
            </div>
          </div>
        </div>
      </ContentContainer>
    </section>
  );
}
