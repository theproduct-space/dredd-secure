// dredd-secure-client-ts Imports
import { Coin } from "dredd-secure-client-ts/cosmos.bank.v1beta1/types/cosmos/base/v1beta1/coin";

// Custom Imports
import TokenElement from "~baseComponents/TokenElement";
import Typography from "~baseComponents/Typography";

//Images
import charityImage from "~assets/charity.png";

interface TipsProps {
  selectedToken: Coin | undefined;
  onClick: () => void;
}

function Tips(props: TipsProps) {
  const { selectedToken, onClick } = props;

  return (
    <div className="border-t-[1px] border-white-200">
      <div className="p-8 flex justify-between items-center">
        <div>
          <Typography variant="body-small" className="flex gap-2">
            <img src={charityImage} alt="charity-heart" className="w-6 h-6" />
            Tips and donations go a long way.
          </Typography>
          <Typography variant="body-small">We are a free service.</Typography>
        </div>
        <div>
          {/* Will take as a prop another component for the base display. Here, it will be a "Add Tip" link or button */}
          <TokenElement
            selectedToken={selectedToken}
            onClick={onClick}
            baseButton={
              <Typography variant="body-small" className="text-orange">
                Add Tip
              </Typography>
            }
          />
        </div>
      </div>
    </div>
  );
}

export default Tips;
