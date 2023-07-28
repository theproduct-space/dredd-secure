import React, { useState } from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import Typography from "~baseComponents/Typography";
import { TextField } from "@mui/material";
import Button from "~baseComponents/Button";
import { toast } from "react-toastify";
import axios from "axios";

interface ICondition {
  type: string;
  prop: string;
}

interface ConditionProps {
  className?: string;
  condition: { condition: ICondition; value: string };
}

// const ApiCondition = ({ className, condition }: ConditionProps) => {
const ApiCondition = () => {
  const [apiUrl, setApiUrl] = useState<string>( // TODO initialise to ""
    // "https://get.geojs.io/v1/ip/country/206.172.129.109",
    "http://api.weatherapi.com/v1/current.json?key=0a1e46cc0a524ccfbc4162404232707&q=London&aqi=no",
  );

  const handleApiFetch = async () => {
    if (apiUrl == "") {
      toast.warning("Please provide a url");
      return;
    }

    try {
      const apiRes = await axios.get(apiUrl);
      console.log("apiRes", apiRes);
    } catch (error: any) {
      toast.error(error?.message);
    }
  };

  return (
    <>
      <TextField
        fullWidth
        variant="outlined"
        placeholder="Enter the API endpoint URL"
        type="text"
        value={apiUrl}
        onChange={(e) => setApiUrl(e.target.value)}
      />
      <Button text="Fetch" className="capitalize" onClick={handleApiFetch} />
    </>
  );
};

export default ApiCondition;
