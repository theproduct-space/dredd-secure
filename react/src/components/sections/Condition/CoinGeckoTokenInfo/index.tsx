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

  const handleSetSelectedOption = (option) => {
    console.log("option", option);

    // handleSetSubConditions(subConditions: Array<ISubConditions>)
  };

  return (
    <div className={className}>
      {condition?.subConditions?.map((subCondition, index) => (
        <div key={`${condition.type}-subcondition-${index}`}>
          <Typography variant="body-small" className="text-white-500 p-1">
            Subcondition #{index}
          </Typography>
          <Dropdown
            choices={relevantFields.map((relevantField) => ({
              label: relevantField.label,
              value: relevantField.name,
            }))}
            selectedOption={{
              label: subCondition.label,
              value: subCondition.value,
            }}
            setSelectedOption={handleSetSelectedOption}
          />
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
