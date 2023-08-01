import React, { useState } from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import Typography from "~baseComponents/Typography";
import { TextField } from "@mui/material";
import Button from "~baseComponents/Button";
import { ICondition } from "~sections/CreateContract/AddConditions";
import CoinGeckoTokenInfo from "./CoinGeckoTokenInfo";

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

  const renderConditionSelector = () => {
    switch (true) {
      // Dates
      case condition.condition.prop === "startDate" ||
        condition.condition.prop === "endDate":
        return (
          <>
            <Typography variant="body-small" className="text-white-500 p-1">
              {condition.condition.type}
            </Typography>
            <CustomDatePicker value={selectedDate} onChange={setSelectedDate} />
          </>
        );
      // API
      case condition.condition.prop === "coingecko-token-info":
        return <CoinGeckoTokenInfo />;
    }
  };

  return (
    <div className={`${className}`}>
      <Typography variant="body-small" className="text-white-500 p-1">
        {children}
      </Typography>
      <div className="flex flex-col gap-2 w-full md:flex-row ">
        <div className="flex gap-5 w-full md:gap-10">
          <select
            value={condition.condition.type}
            onChange={(e) => handleChangeCondition(e, index)}
            className="w-full bg-buttonBg text-white-1000 p-4 border border-white-200 rounded focus:outline-none focus:border-orange"
          >
            <option value="Select Condition Type" disabled selected>
              Select Condition Type
            </option>
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
          <button
            onClick={() => handleRemoveCondition(index)}
            className="text-orange text-6xl"
          >
            -
          </button>
        </div>
      </div>
      <div className="pl-10 py-6">{renderConditionSelector()}</div>
    </div>
  );
};

export default Condition;
