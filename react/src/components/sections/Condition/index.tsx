import React, { useEffect, useState } from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import Typography from "~baseComponents/Typography";
import { TextField } from "@mui/material";
import Button from "~baseComponents/Button";
import {
  ICondition,
  ISubCondition,
} from "~sections/CreateContract/AddConditions";
import CoinMarketCapTokenInfo from "./CoinMarketCapTokenInfo";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";

interface ConditionProps {
  children?: React.ReactNode;
  className?: string;
  condition: ICondition;
  index: number;
  handleSelectCondition: (
    e: React.ChangeEvent<HTMLSelectElement>,
    index: number,
  ) => void;
  handleRemoveCondition: (index: number) => void;
  displayConditionTypes: () => JSX.Element[];
  // setConditions: (newConditions: ICondition[]) => void;
  setConditions: React.Dispatch<React.SetStateAction<ICondition[]>>;
}

const Condition = ({
  children,
  className,
  condition,
  index,
  handleSelectCondition,
  handleRemoveCondition,
  displayConditionTypes,
  setConditions,
}: ConditionProps) => {
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);

  const handleSetSubConditions = (subConditions: Array<ISubCondition>) => {
    setConditions((prev: ICondition[]) => {
      const temp = [...prev];
      temp[index].subConditions = subConditions;
      return temp;
    });
  };

  const handleRemoveSubConditions = (subConditionIndex: number) => {
    setConditions((prev: ICondition[]) => {
      const temp = [...prev];
      const subConditionsArray = temp[index]?.subConditions;

      if (subConditionsArray) {
        const newSubConditions = [
          ...subConditionsArray.slice(0, subConditionIndex),
          ...subConditionsArray.slice(subConditionIndex + 1),
        ];

        temp[index].subConditions = newSubConditions;
      }

      return temp;
    });
  };

  const handleAddNewEmptySubCondition = (
    condition: ICondition,
    endpoint: string,
  ) => {
    const { subConditions } = condition;
    const tempSubConditions = subConditions?.slice();
    tempSubConditions?.push({
      conditionType: "eq",
      dataType: configuredAPIEndpoints.data[endpoint].relevant_fields[0].type,
      name: configuredAPIEndpoints.data[endpoint].relevant_fields[0].name,
      path: [
        // TODO this should depend on the condition.name (this case is for coinmarketcap-token-info)
        "data",
        String(condition.tokenOfInterest?.id ?? "1"),
        "quote",
        "USD",
        configuredAPIEndpoints.data[endpoint].relevant_fields[0].name,
      ],
      label: configuredAPIEndpoints.data[endpoint].relevant_fields[0].label,
      value: undefined,
    });

    setConditions((prev: ICondition[]) => {
      const temp = [...prev];
      temp[index].subConditions = tempSubConditions;
      return temp;
    });
  };

  const handleSetValue = (value: string | number) => {
    setConditions((prev: ICondition[]) => {
      const temp = [...prev];
      temp[index].value = value;
      return temp;
    });
  };

  const renderConditionSelector = () => {
    switch (true) {
      // Dates
      case condition.name === "startDate" || condition.name === "endDate":
        return (
          <>
            <Typography variant="body-small" className="text-white-500 p-1">
              {condition.label}
            </Typography>
            <CustomDatePicker
              value={selectedDate}
              onChange={(date: any) => {
                const epochDate = String(new Date(date.$d).getTime() / 1000);
                handleSetValue(epochDate);
                setSelectedDate(date);
              }}
            />
          </>
        );
      // API
      case condition.name === "coinmarketcap-token-info":
        return (
          <CoinMarketCapTokenInfo
            index={index}
            condition={condition}
            setConditions={setConditions}
            handleSetSubConditions={handleSetSubConditions}
            handleRemoveSubConditions={handleRemoveSubConditions}
          />
        );
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
            value={condition.label}
            onChange={(e) => handleSelectCondition(e, index)}
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
      <div className="pl-6 py-6">
        {renderConditionSelector()}
        {condition.subConditions && (
          <button
            onClick={() =>
              handleAddNewEmptySubCondition(condition, condition.name)
            }
          >
            <Typography
              variant="body-small"
              className="font-revalia text-orange py-4"
            >
              + Add Subcondition
            </Typography>
          </button>
        )}
      </div>
    </div>
  );
};

export default Condition;
