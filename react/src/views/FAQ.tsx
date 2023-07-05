/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable import/no-unresolved */
import React, { useState } from "react";
import minus from "../assets/minusIcon.svg";
import plus from "../assets/plusIcon.svg";

function FAQ() {
  const [toggles, setToggles] = useState([false, false, false]);

  const handleToggle = (index: number) => {
    setToggles((prevState) => {
      const updatedToggles = [...prevState];
      updatedToggles[index] = !updatedToggles[index];
      return updatedToggles;
    });
  };

  return (
    <section className="bg-gray">
      <div className="container mx-auto py-20 flex gap-12">
        <h2 className="text-4xl font-bold text-white-1000 font-revalia">Frequently asked questions</h2>
        <div className="flex flex-col gap-4 w-1/2">
          <div className="bg-white-200 rounded-xl p-4">
            <div className="flex justify-between items-center cursor-pointer" onClick={() => handleToggle(0)}>
              <p className="text-p1 text-white-1000">How is the escrow service secure ?</p>
              <img src={toggles[0] ? minus : plus} alt="toggle icon" />
            </div>
            {toggles[0] && (
              <p className="text-white-1000">
                The escrow service, developed with the Cosmos SDK, is secure as it uses a two-party Escrow Contract to
                hold assets until agreed conditions are met. It adheres to latest security practices and undergoes
                third-party audits to ensure asset security.
              </p>
            )}
          </div>
          <div className="bg-white-200 rounded-xl p-4">
            <div className="flex justify-between items-center cursor-pointer" onClick={() => handleToggle(1)}>
              <p className="text-p1 text-white-1000">What types of assets can I put in an escrow ?</p>
              <img src={toggles[1] ? minus : plus} alt="toggle icon" />
            </div>
            {toggles[1] && (
              <p className="text-white-1000">
                Any digital asset accessible through the Cosmos Hub, including native tokens, wrapped tokens, and
                Non-Fungible Tokens (NFTs), can be put in escrow.
              </p>
            )}
          </div>
          <div className="bg-white-200 rounded-xl p-4">
            <div className="flex justify-between items-center cursor-pointer" onClick={() => handleToggle(2)}>
              <p className="text-p1 text-white-1000">How is Cosmos SDK beneficial for the escrow service ?</p>
              <img src={toggles[2] ? minus : plus} alt="toggle icon" />
            </div>
            {toggles[2] && (
              <p className="text-white-1000">
                The Cosmos SDK enhances the escrow service by providing scalability, seamless integration with the
                Cosmos ecosystem, robust security mechanisms, support for cross-chain transactions, and
                developer-friendly tools for easy development and deployment.
              </p>
            )}
          </div>
        </div>
      </div>
    </section>
  );
}

export default FAQ;
