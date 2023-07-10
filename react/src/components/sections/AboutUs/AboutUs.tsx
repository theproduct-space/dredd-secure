/* eslint-disable import/no-unresolved */
/* eslint-disable react/no-unescaped-entities */
import { Typography } from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";

export default function AboutUs() {
  return (
    <section className="bg-gray">
      <ContentContainer>
        <Typography variant="h2" className="text-white-1000 font-revalia py-6">
          About Us
        </Typography>
        <Typography variant="h6" as={"blockquote"} className="text-white-1000 text-p1">
          ProductShop is a web development studio based in Montréal, Canada. We’re a team of skilled engineers and
          professionals with expertise in payment solutions, blockchain development, and smart contract creation. We are
          committed to delivering high-quality and customized solutions for our diverse clientele across various
          industries. Our team consistently innovates, plans, and strategizes to create cutting-edge solutions.
        </Typography>
        <Typography variant="h6" as={"blockquote"} className="text-white-1000 text-p1">
          Our recent venture, DreddSecure, is a testament to our commitment. An advanced blockchain-based escrow app for
          the Cosmos Hub ecosystem, designed to enhance secure, reliable, and scalable blockchain transactions. At
          ProductShop, we're not just adapting to the digital revolution, we're actively shaping it, and creating a
          safer digital transaction landscape.
        </Typography>
      </ContentContainer>
    </section>
  );
}
