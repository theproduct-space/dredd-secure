import React, { useEffect, useState } from "react";
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

const getDropdownChoices = (dataType: string) => {
  switch (dataType) {
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
const getSelectedChoice = (conditionType: string) => {
  switch (conditionType) {
    case "eq":
      return {
        label: "Equal",
        value: "eq",
      };
    case "gt":
      return {
        label: "Greater Than",
        value: "gt",
      };
    case "lt":
      return {
        label: "Less Than",
        value: "lt",
      };
  }
};

interface SubConditionDetailProps {
  className?: string;
  dataType: string;
  conditionType: string;
  value: string | number | undefined;
  handleSetValue: (value: any) => void;
  handleSetConditionType: (conditionType: string) => void;
}

const SubConditionDetail = ({
  className,
  dataType,
  conditionType,
  value,
  handleSetValue,
  handleSetConditionType,
}: SubConditionDetailProps) => {
  const dropdownChoices = getDropdownChoices(dataType);
  const selectedChoice = getSelectedChoice(conditionType);

  return (
    <div className={className}>
      <div>
        <Typography variant="body-small" className="text-white-500 p-1">
          Details
        </Typography>
        <div className="flex gap-3 items-center">
          <Dropdown
            choices={dropdownChoices}
            selectedOption={selectedChoice}
            setSelectedOption={(option) => {
              handleSetConditionType(option.value);
            }}
          />
          <TextField
            type={dataType}
            variant="outlined"
            placeholder="Value"
            value={value}
            onChange={(e) => {
              if (dataType === "number") {
                handleSetValue(Number(e.target.value));
              } else {
                handleSetValue(String(e.target.value));
              }
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default SubConditionDetail;
