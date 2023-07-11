import FAQItem from "~baseComponents/FAQItem";
import ContentContainer from "~layouts/ContentContainer";

function FAQ() {
  return (
    <ContentContainer className="py-20 md:py-40">
      <div className="py-20 flex flex-col lg:flex-row lg:justify-between gap-12">
        <h2 className="text-4xl font-bold text-white-1000 font-revalia">
          Frequently asked questions
        </h2>
        <div className="flex flex-col gap-4 w-full ">
          <FAQItem
            question="What is an Escrow Contract?"
            answer="An Escrow Contract is a self-executing contract with the agreement terms directly written into code. It holds the assets involved in a transaction and only releases them when the specified conditions are met, eliminating the need for a trusted intermediary."
          />
          <FAQItem
            question="Why use our escrow service?"
            answer={
              <ul className="flex flex-col gap-2">
                <li>
                  <b>Trust</b>: Assets are only released when agreed conditions
                  are met, fostering trust between parties.
                </li>
                <li>
                  <b>Security</b>: Built on the Cosmos SDK, our service follows
                  the latest security best practices, keeping your assets safe.
                </li>
                <li>
                  <b>Efficiency</b>: Our escrow service eliminates the
                  middleman, making transactions faster and cheaper.
                </li>
              </ul>
            }
          />
          <FAQItem
            question="What assets are compatible with DreddSecure?"
            answer="Any assets that can be accessed through the Cosmos Hub, including native token and wrapped tokens are compatible with DreddSecure."
          />
          <FAQItem
            question="What are some use cases of DreddSecure?"
            answer="DreddSecure can be used for peer-to-peer exchanges, token offerings, cross-chain asset swaps, decentralized marketplaces, decentralized crowdfunding, peer-to-peer lending, and decentralized insurance."
          />
          <FAQItem
            question="How does it work?"
            answer="Initiate a secure transaction in DreddSecure by creating an escrow contract, input parameters and transfer assets, allow our smart contract to evaluate predefined conditions, and assets will be automatically be released to the respective parties upon satisfying all conditions, ensuring a seamless, secure process."
          />
          <FAQItem
            question="How can I start using DreddSecure?"
            answer="DreddSecure is still in the development stage. Once it is implemented, users will be able to initiate transactions through the DreddSecure web interface, integrating with popular wallet solutions for easy access to the escrow app. Stay tuned for updates."
          />
        </div>
      </div>
    </ContentContainer>
  );
}

export default FAQ;
