
/* eslint-disable react/no-unescaped-entities */
export default function Home() {
  return (
    <div className="w-full min-h-screen">
      <div className="container mx-auto">
        <div className="">
          <div>
            <div>
              <div className="text-4xl font-bold">Dredd-Secure</div>
              <blockquote>
                "Discover our innovative escrow service designed for the Cosmos Hub ecosystem. Using the powerful Cosmos
                SDK, we bring enhanced security and efficiency to transactions within the ecosystem."
              </blockquote>
            </div>
            <div>
              <div className="text-3xl font-bold">What is Dredd-Secure?</div>
              <blockquote>
                "Experience secure, automated contracts allowing parties to transact confidently. By holding assets
                until specific conditions are met, we eliminate the risk of fraud and disputes for smoother, safer
                transactions."
              </blockquote>
            </div>
            <div>
              <div className="text-3xl font-bold">Why use our escrow service?</div>
              <ul>
                <li>
                  <b>Trust</b>: Assets are only released when agreed conditions are met, fostering trust between
                  parties.
                </li>
                <li>
                  <b>Security</b>: Built on the Cosmos SDK, our service follows the latest security best practices,
                  keeping your assets safe.
                </li>
                <li>
                  <b>Efficiency</b>: Our escrow service eliminates the middleman, making transactions faster and
                  cheaper.
                </li>
              </ul>
            </div>
            <div>
              <div className="text-3xl font-bold">How does it works?</div>
              <div>
                <div className="text-2xl font-bold">Step 1: Create the Escrow Contract</div>
                <div>
                  Begin a secure transaction by creating an escrow contract. Input necessary parameters and transfer
                  assets into the contract. Initiations are possible as both buyer or seller.
                </div>
              </div>
              <div>
                <div className="text-2xl font-bold">Step 2: Check Conditions</div>
                <div>
                  Our smart contract evaluates predefined conditions to ensure agreed-upon terms are met before asset
                  transfer.
                </div>
              </div>
              <div>
                <div className="text-2xl font-bold">Step 3: Release Assets</div>
                <div>
                  Upon satisfying all conditions, assets are automatically released to the respective parties, ensuring
                  a seamless, secure process.
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
