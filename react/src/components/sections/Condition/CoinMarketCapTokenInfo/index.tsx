import Typography from "~baseComponents/Typography";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";
import Dropdown, { DropdownChoice } from "~baseComponents/Dropdown";
import {
  ICondition,
  ISubCondition,
} from "~sections/CreateContract/AddConditions";
import SubConditionDetail from "../SubConditionDetail";
import CoinMarketCapTokenSelector, {
  CoinMarketCapTokenI,
} from "../CoinMarketCapTokenSelector";

interface CoinMarketCapTokenInfoProps {
  index: number;
  className?: string;
  condition: ICondition;
  setConditions: React.Dispatch<React.SetStateAction<ICondition[]>>;
  handleSetSubConditions: (subConditions: Array<ISubCondition>) => void;
  handleRemoveSubConditions: (subConditionIndex: number) => void;
}

const relevantFields =
  configuredAPIEndpoints.data["coinmarketcap-token-info"]["relevant_fields"];

const CoinMarketCapTokenInfo = ({
  index,
  className = "",
  condition,
  setConditions,
  handleSetSubConditions,
  handleRemoveSubConditions,
}: CoinMarketCapTokenInfoProps) => {
  const handleSetSelectedOption = (
    option: DropdownChoice,
    subConditionIndex: number,
  ) => {
    const pastSubConditions = condition?.subConditions?.slice() || [];

    if (pastSubConditions && pastSubConditions[subConditionIndex]) {
      pastSubConditions[subConditionIndex].label = option.label;
      pastSubConditions[subConditionIndex].name = option.value;
      pastSubConditions[subConditionIndex].dataType = option.type || "";

      // update the path to access the relevant field in the CoinMarketCap API response
      pastSubConditions[subConditionIndex].path = [
        "data",
        String(condition?.tokenOfInterest?.id),
        "quote",
        "USD",
        pastSubConditions[subConditionIndex].name,
      ];
    }

    handleSetSubConditions(pastSubConditions);
  };

  const handleSetSubConditionValue = (
    value: ISubCondition["value"],
    index: number,
  ) => {
    const pastSubConditions = condition?.subConditions?.slice() || [];

    if (pastSubConditions && pastSubConditions[index]) {
      pastSubConditions[index].value = value;
    }

    handleSetSubConditions(pastSubConditions);
  };

  const handleSetSubConditionType = (conditionType: string, index: number) => {
    const pastSubConditions = condition?.subConditions?.slice() || [];

    if (pastSubConditions && pastSubConditions[index]) {
      pastSubConditions[index].conditionType = conditionType;
    }

    handleSetSubConditions(pastSubConditions);
  };

  const handleSetTokenOfInterest = (tokenOfInterest: CoinMarketCapTokenI) => {
    setConditions((prev: ICondition[]) => {
      const temp = [...prev];
      temp[index].tokenOfInterest = { ...tokenOfInterest };

      // update every subCondition path field
      temp[index].subConditions = temp[index].subConditions?.map(
        (subCondition: ISubCondition) => {
          const newSubCondition = { ...subCondition };
          newSubCondition.path[1] = String(tokenOfInterest?.id);

          return newSubCondition;
        },
      );
      return temp;
    });
  };

  return (
    <div className={`w-full flex flex-col gap-5 ${className}`}>
      <div>
        <CoinMarketCapTokenSelector
          selectedToken={condition.tokenOfInterest}
          setSelectedToken={handleSetTokenOfInterest}
        />
      </div>
      {condition?.subConditions?.map((subCondition, index) => (
        <div key={`${condition.type}-subcondition-${index}`}>
          <Typography variant="body-small" className="text-white-500 p-1">
            Subcondition #{index + 1}
          </Typography>
          <div className="flex items-start w-full justify-between gap-5 md:gap-10 mb-1">
            <Dropdown
              choices={relevantFields.map((relevantField) => ({
                label: relevantField.label,
                value: relevantField.name,
                type: relevantField.type,
              }))}
              selectedOption={{
                label: subCondition.label,
                value: subCondition.value,
                type: subCondition.dataType,
              }}
              setSelectedOption={(choice) =>
                handleSetSelectedOption(choice, index)
              }
              className="w-full"
            />
            <button
              onClick={() => handleRemoveSubConditions(index)}
              className="relative text-orange text-5xl"
            >
              -
            </button>
          </div>
          <div className="pl-3">
            <SubConditionDetail
              dataType={subCondition.dataType}
              conditionType={subCondition.conditionType}
              value={subCondition.value}
              handleSetValue={(value) =>
                handleSetSubConditionValue(value, index)
              }
              handleSetConditionType={(conditionType) =>
                handleSetSubConditionType(conditionType, index)
              }
            />
          </div>
        </div>
      ))}
    </div>
  );
};

export default CoinMarketCapTokenInfo;
