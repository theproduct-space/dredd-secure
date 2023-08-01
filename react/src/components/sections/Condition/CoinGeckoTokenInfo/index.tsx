import React, { useState } from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import Typography from "~baseComponents/Typography";
import { TextField } from "@mui/material";
import Button from "~baseComponents/Button";
import { toast } from "react-toastify";
import axios from "axios";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";
import Dropdown from "~baseComponents/Dropdown";

interface ICondition {
  type: string;
  prop: string;
}

interface ConditionProps {
  className?: string;
  condition: { condition: ICondition; value: string };
}

const isDevelopment = process.env.NODE_ENV === "development";
const API_BASE_URL = isDevelopment ? "http://localhost:3000" : "";
const relevantFields =
  configuredAPIEndpoints.data["coingecko-token-info"]["relevant_fields"];
console.log("relevantFields", relevantFields);

// const ApiCondition = ({ className, condition }: ConditionProps) => {
const CoinGeckoTokenInfo = () => {
  return (
    <>
      <Typography variant="body-small" className="text-white-500 p-1">
        Subcondition #1
      </Typography>

      {/* <TextField
        fullWidth
        variant="outlined"
        placeholder="Enter the API endpoint URL"
        type="text"
        value={apiUrl}
        onChange={(e) => setApiUrl(e.target.value)}
      />
      <Button text="Fetch" className="capitalize" onClick={handleApiFetch} /> */}
    </>
  );
};

export default CoinGeckoTokenInfo;
