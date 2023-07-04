import React from "react";

function FAQ() {
  return (
    <div>
      <div className="text-4xl font-bold">FAQ</div>
      <div>
        <div className="text-2xl font-bold">How is the escrow service secure?</div>
        <div>
          The escrow service, developed with the Cosmos SDK, is secure as it uses a two-party Escrow Contract to hold
          assets until agreed conditions are met. It adheres to latest security practices and undergoes third-party
          audits to ensure asset security.
        </div>
      </div>
      <div>
        <div className="text-2xl font-bold">What types of assets can I put in an escrow?</div>
        <div>
          Any digital asset accessible through the Cosmos Hub, including native tokens, wrapped tokens, and Non-Fungible
          Tokens (NFTs), can be put in escrow.
        </div>
      </div>
      <div>
        <div className="text-2xl font-bold">How is Cosmos SDK beneficial for the escrow service?</div>
        <div>
          The Cosmos SDK enhances the escrow service by providing scalability, seamless integration with the Cosmos
          ecosystem, robust security mechanisms, support for cross-chain transactions, and developer-friendly tools for
          easy development and deployment.
        </div>
      </div>
    </div>
  );
}

export default FAQ;
