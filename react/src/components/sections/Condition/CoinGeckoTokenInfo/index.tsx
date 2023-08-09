import React, { useState } from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import Typography from "~baseComponents/Typography";
import { TextField } from "@mui/material";
import Button from "~baseComponents/Button";
import { toast } from "react-toastify";
import axios from "axios";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";
import Dropdown from "~baseComponents/Dropdown";
import {
  ICondition,
  ISubConditions,
} from "~sections/CreateContract/AddConditions";
import SubConditionDetail from "../SubConditionDetail";
import { DropdownChoice } from "~baseComponents/Dropdown";

interface CoinGeckoTokenInfoProps {
  className?: string;
  condition: ICondition;
  handleSetSubConditions: (subConditions: Array<ISubConditions>) => void;
}

const relevantFields =
  configuredAPIEndpoints.data["coingecko-token-info"]["relevant_fields"];
console.log({ configuredAPIEndpoints });
console.log("relevantFields", relevantFields);

const CoinGeckoTokenInfo = ({
  className,
  condition,
  handleSetSubConditions,
}: CoinGeckoTokenInfoProps) => {
  console.log("condition", condition);

  const handleSetSelectedOption = (
    option: DropdownChoice,
    index: number,
    type: string,
  ) => {
    console.log("optionnnnnn", option);
    console.log("index", index);

    const pastSubConditions = condition?.subConditions?.slice() || [];

    if (pastSubConditions && pastSubConditions[index]) {
      pastSubConditions[index].label = option.label;
      pastSubConditions[index].name = option.value;
    }

    handleSetSubConditions(pastSubConditions);
  };

  return (
    <div className={`w-full ${className}`}>
      {condition?.subConditions?.map((subCondition, index) => (
        <div key={`${condition.type}-subcondition-${index}`}>
          <Typography variant="body-small" className="text-white-500 p-1">
            Subcondition #{index}
          </Typography>
          <div className="flex items-start w-full justify-between gap-5 md:gap-10">
            <Dropdown
              choices={relevantFields.map((relevantField) => ({
                label: relevantField.label,
                value: relevantField.name,
              }))}
              selectedOption={{
                label: subCondition.label,
                value: subCondition.value,
              }}
              setSelectedOption={(choice) =>
                handleSetSelectedOption(choice, index, subCondition.dataType)
              }
              className="w-full"
            />
            <button
              onClick={() => console.log("hey")}
              className="relative text-orange text-5xl"
            >
              -
            </button>
          </div>
          <div className="pl-6">
            <SubConditionDetail subConditionType={subCondition.dataType} />
          </div>
        </div>
      ))}

      {/* <TextField
        fullWidth
        variant="outlined"
        placeholder="Value"
        type="text"
        // value={apiUrl}
        // onChange={(e) => setApiUrl(e.target.value)}
      /> */}
    </div>
  );
};

export default CoinGeckoTokenInfo;
