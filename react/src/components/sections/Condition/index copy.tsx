import React, { useState } from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import Typography from "~baseComponents/Typography";
import { TextField } from "@mui/material";
import Button from "~baseComponents/Button";

interface ICondition {
  type: string;
  prop: string;
}

interface ConditionProps {
  children?: React.ReactNode;
  className?: string;
  condition: { condition: ICondition; value: string };
  index: number;
  handleChangeCondition: (
    e: React.ChangeEvent<HTMLSelectElement>,
    id: number,
  ) => void;
  handleRemoveCondition: (id: number) => void;
  displayConditionTypes: () => JSX.Element[];
}

const Condition = ({
  children,
  className,
  condition,
  index,
  handleChangeCondition,
  handleRemoveCondition,
  displayConditionTypes,
}: ConditionProps) => {
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [apiUrl, setApiUrl] = useState<string | null>(null);

  const renderCondition = () => {
    switch (true) {
      case condition.condition.prop === "startDate" ||
        condition.condition.prop === "endDate":
        return (
          <div className="flex gap-2">
            <select
              value={condition.condition.type}
              onChange={(e) => handleChangeCondition(e, index)}
              className="w-60 bg-buttonBg text-white-1000 p-4 border border-white-200 rounded focus:outline-none focus:border-orange"
            >
              {displayConditionTypes().map((option, optionIndex) => (
                <option
                  key={optionIndex}
                  value={option.props.value}
                  className="bg-buttonBg text-white-1000 p-4 border border-white-500"
                >
                  {option.props.children}
                </option>
              ))}
            </select>
            <CustomDatePicker value={selectedDate} onChange={setSelectedDate} />
            <button
              onClick={() => handleRemoveCondition(index)}
              className="text-orange text-4xl"
            >
              -
            </button>
          </div>
        );
      case condition.condition.prop === "api":
        return (
          <>
            <TextField
              fullWidth
              variant="outlined"
              placeholder="Enter the API endpoint URL"
              type="text"
              value={apiUrl}
              onChange={(e) => setApiUrl(e.target.value)}
              sx={{ mb: "12px" }}
            />
            <Button text="Fetch API" className="capitalize" />
          </>
        );
    }
  };

  return (
    <div className={`${className}`}>
      <Typography variant="body-small" className="text-white-500 p-1">
        {children}
      </Typography>
      {renderCondition()}
    </div>
  );
};

export default Condition;
