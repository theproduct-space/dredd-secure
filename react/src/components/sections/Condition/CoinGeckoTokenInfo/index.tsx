import Typography from "~baseComponents/Typography";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";
import Dropdown from "~baseComponents/Dropdown";
import {
  ICondition,
  ISubConditions,
} from "~sections/CreateContract/AddConditions";
import SubConditionDetail from "../SubConditionDetail";
import { DropdownChoice } from "~baseComponents/Dropdown";
import axios from "axios";
import { useState } from "react";
import CoinGeckoTokenSelector from "../CoinGeckoTokenSelector";

interface CoinGeckoTokenInfoProps {
  className?: string;
  condition: ICondition;
  handleSetSubConditions: (subConditions: Array<ISubConditions>) => void;
  handleRemoveSubConditions: (subConditionIndex: number) => void;
}

const relevantFields =
  configuredAPIEndpoints.data["coingecko-token-info"]["relevant_fields"];

const CoinGeckoTokenInfo = ({
  className = "",
  condition,
  handleSetSubConditions,
  handleRemoveSubConditions,
}: CoinGeckoTokenInfoProps) => {
  const [selectedToken, setSelectedToken] = useState(null); // TODO add this data into the "condition" state

  const handleSetSelectedOption = (option: DropdownChoice, index: number) => {
    const pastSubConditions = condition?.subConditions?.slice() || [];

    if (pastSubConditions && pastSubConditions[index]) {
      pastSubConditions[index].label = option.label;
      pastSubConditions[index].name = option.value;
      pastSubConditions[index].dataType = option.type || "";
    }

    handleSetSubConditions(pastSubConditions);
  };

  const handleSetSubConditionValue = (value: any, index: number) => {
    console.log("value", value);
    console.log("typeof value", typeof value);
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

  return (
    <div className={`w-full flex flex-col gap-5 ${className}`}>
      <div>
        <CoinGeckoTokenSelector
          selectedToken={selectedToken}
          setSelectedToken={setSelectedToken}
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

export default CoinGeckoTokenInfo;
