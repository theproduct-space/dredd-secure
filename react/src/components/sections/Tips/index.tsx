// Custom Imports
import TokenElement from "~baseComponents/TokenElement";
import Typography from "~baseComponents/Typography";

//Images
import charityImage from "~assets/charity.png";
import { IToken } from "~baseComponents/TokenSelector";

interface TipsProps {
  token: IToken | undefined;
  onClick: () => void;
  selectedAmount: number;
  setSelectedAmount: (amount: number) => void;
}

function Tips(props: TipsProps) {
  const { token, onClick, selectedAmount, setSelectedAmount } = props;
  console.log(token);
  const handleTokenClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    onClick();
  };

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
          {token ? (
            <TokenElement
              token={token}
              onClick={onClick}
              selectedAmount={selectedAmount}
              setSelectedAmount={setSelectedAmount}
            />
          ) : (
            <button onClick={handleTokenClick}>
              <Typography variant="body-small" className="text-orange">
                Add Tip
              </Typography>
            </button>
          )}
        </div>
      </div>
    </div>
  );
}

export default Tips;
