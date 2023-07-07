/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable import/no-unresolved */
import FAQItem from "~baseComponents/FAQItem/FAQItem";
import ContentContainer from "~layouts/ContentContainer";

function FAQ() {
  return (
    <section className="bg-gray">
      <ContentContainer>
        <div className="py-20 flex flex-col md:flex-row md:justify-between gap-12">
          <h2 className="text-4xl font-bold text-white-1000 font-revalia">
            Frequently asked questions
          </h2>
          <div className="flex flex-col gap-4 w-full ">
            <FAQItem
              question="What is DreddSecure?"
              answer="Experience secure, automated contracts allowing parties to transact confidently. By holding assets until specific conditions are met, we eliminate the risk of fraud and disputes for smoother, safer transactions."
            />
            <FAQItem
              question="Why use our escrow service?"
              answer={
                <ul className="flex flex-col gap-2">
                  <li>
                    <b>Trust</b>: Assets are only released when agreed
                    conditions are met, fostering trust between parties.
                  </li>
                  <li>
                    <b>Security</b>: Built on the Cosmos SDK, our service
                    follows the latest security best practices, keeping your
                    assets safe.
                  </li>
                  <li>
                    <b>Efficiency</b>: Our escrow service eliminates the
                    middleman, making transactions faster and cheaper.
                  </li>
                </ul>
              }
            />
            <FAQItem
              question="How does it work?"
              answer={
                <div className="flex flex-col gap-2">
                  <div>
                    <div className="font-bold">
                      Step 1: Create the Escrow Contract
                    </div>
                    <div>
                      Begin a secure transaction by creating an escrow contract.
                      Input necessary parameters and transfer assets into the
                      contract. Initiations are possible as both buyer or
                      seller.
                    </div>
                  </div>
                  <div>
                    <div className="font-bold">Step 2: Check Conditions</div>
                    <div>
                      Our smart contract evaluates predefined conditions to
                      ensure agreed-upon terms are met before asset transfer.
                    </div>
                  </div>
                  <div>
                    <div className="font-bold">Step 3: Release Assets</div>
                    <div>
                      Upon satisfying all conditions, assets are automatically
                      released to the respective parties, ensuring a seamless,
                      secure process.
                    </div>
                  </div>
                </div>
              }
            />
            <FAQItem
              question="How is the escrow service secure?"
              answer="The escrow service, developed with the Cosmos SDK, is secure as it uses a two-party Escrow Contract to hold assets until agreed conditions are met. It adheres to the latest security practices and undergoes third-party audits to ensure asset security."
            />
            <FAQItem
              question="What types of assets can I put in an escrow?"
              answer="Any digital asset accessible through the Cosmos Hub, including native tokens & wrapped tokens can be put in escrow."
            />
            <FAQItem
              question="How is Cosmos SDK beneficial for the escrow service?"
              answer="The Cosmos SDK enhances the escrow service by providing scalability, seamless integration with the Cosmos ecosystem, robust security mechanisms, support for cross-chain transactions, and developer-friendly tools for easy development and deployment."
            />
          </div>
        </div>
      </ContentContainer>
    </section>
  );
}

export default FAQ;
