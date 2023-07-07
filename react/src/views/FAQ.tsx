/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable import/no-unresolved */
import React, { useState } from "react";
import FAQItem from "~baseComponents/FAQItem/FAQItem";
import ContentContainer from "~layouts/ContentContainer";

function FAQ() {
  const [toggles, setToggles] = useState([
    false,
    false,
    false,
    false,
    false,
    false,
  ]);

  const handleToggle = (index: number) => {
    setToggles((prevState) => {
      const updatedToggles = [...prevState];
      updatedToggles[index] = !updatedToggles[index];
      return updatedToggles;
    });
  };

  return (
    <section className="bg-gray">
      <ContentContainer>
        <div className="py-20 flex flex-col md:flex-row gap-12">
          <h2 className="text-4xl font-bold text-white-1000 font-revalia">
            Frequently asked questions
          </h2>
          <div className="flex flex-col gap-4 w-full md:w-1/2">
            <FAQItem
              question="What is Dredd-Secure?"
              answer="Experience secure, automated contracts allowing parties to transact confidently. By holding assets until specific conditions are met, we eliminate the risk of fraud and disputes for smoother, safer transactions."
              isOpen={toggles[3]}
              onToggle={() => handleToggle(3)}
            />
            <FAQItem
              question="Why use our escrow service?"
              answer={
                <ul className="text-white-1000">
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
              isOpen={toggles[4]}
              onToggle={() => handleToggle(4)}
            />
            <FAQItem
              question="How does it work?"
              answer={
                <div>
                  <div>
                    <div className="text-2xl font-bold">
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
                    <div className="text-2xl font-bold">
                      Step 2: Check Conditions
                    </div>
                    <div>
                      Our smart contract evaluates predefined conditions to
                      ensure agreed-upon terms are met before asset transfer.
                    </div>
                  </div>
                  <div>
                    <div className="text-2xl font-bold">
                      Step 3: Release Assets
                    </div>
                    <div>
                      Upon satisfying all conditions, assets are automatically
                      released to the respective parties, ensuring a seamless,
                      secure process.
                    </div>
                  </div>
                </div>
              }
              isOpen={toggles[5]}
              onToggle={() => handleToggle(5)}
            />
            <FAQItem
              question="How is the escrow service secure?"
              answer="The escrow service, developed with the Cosmos SDK, is secure as it uses a two-party Escrow Contract to hold assets until agreed conditions are met. It adheres to the latest security practices and undergoes third-party audits to ensure asset security."
              isOpen={toggles[0]}
              onToggle={() => handleToggle(0)}
            />
            <FAQItem
              question="What types of assets can I put in an escrow?"
              answer="Any digital asset accessible through the Cosmos Hub, including native tokens, wrapped tokens, and Non-Fungible Tokens (NFTs), can be put in escrow."
              isOpen={toggles[1]}
              onToggle={() => handleToggle(1)}
            />
            <FAQItem
              question="How is Cosmos SDK beneficial for the escrow service?"
              answer="The Cosmos SDK enhances the escrow service by providing scalability, seamless integration with the Cosmos ecosystem, robust security mechanisms, support for cross-chain transactions, and developer-friendly tools for easy development and deployment."
              isOpen={toggles[2]}
              onToggle={() => handleToggle(2)}
            />
          </div>
        </div>
      </ContentContainer>
    </section>
  );
}

export default FAQ;
