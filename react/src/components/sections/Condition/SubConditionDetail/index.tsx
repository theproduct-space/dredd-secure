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

interface SubConditionDetailProps {
  className?: string;
  subConditionType: string;
  // handleSetSubConditions: (subConditions: Array<ISubConditions>) => void;
}

// const relevantFields =
//   configuredAPIEndpoints.data["coingecko-token-info"]["relevant_fields"];
// console.log({ configuredAPIEndpoints });
// console.log("relevantFields", relevantFields);

const getDropdownChoices = (subConditionType: string) => {
  switch (subConditionType) {
    case "text":
      return [
        {
          label: "Equal",
          value: "eq",
        },
      ];
    case "number":
      return [
        {
          label: "Equal",
          value: "eq",
        },
        {
          label: "Greater Than",
          value: "gt",
        },
        {
          label: "Less Than",
          value: "lt",
        },
      ];
    case "default":
      return [
        {
          label: "Equal",
          value: "eq",
        },
      ];
  }
};

const SubConditionDetail = ({
  className,
  subConditionType,
}: // handleSetSubConditions,
SubConditionDetailProps) => {
  console.log("condition", subConditionType);

  const [value, setValue] = useState("");
  const dropdownChoices = getDropdownChoices(subConditionType);
  const [selectedOption, setSelectedOption] = useState(
    dropdownChoices && dropdownChoices[0],
  );

  // const handleSetSelectedOption = (option) => {
  //   console.log("option", option);

  //   // handleSetSubConditions(subConditions: Array<ISubConditions>)
  // };

  return (
    <div className={className}>
      <div>
        <Typography variant="body-small" className="text-white-500 p-1">
          Details
        </Typography>
        <div className="flex gap-3 items-center">
          <Dropdown
            choices={dropdownChoices}
            selectedOption={selectedOption}
            setSelectedOption={setSelectedOption}
          />
          <TextField
            type={subConditionType}
            variant="outlined"
            placeholder="Value"
            value={value}
            onChange={(e) => setValue(e.target.value)}
          />
        </div>
      </div>
    </div>
  );
};

export default SubConditionDetail;
