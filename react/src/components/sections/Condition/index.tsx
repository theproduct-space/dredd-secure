import React from "react";
import CustomDatePicker from "~baseComponents/DatePicker";
import FilterDropDown from "~baseComponents/FilterDropDown";
import Typography from "~baseComponents/Typography";

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
  const [selectedDate, setSelectedDate] = React.useState<Date | null>(null);
  return (
    <div className={`${className}`}>
      <Typography variant="body-small" className="text-white-500 p-1">
        {children}
      </Typography>
      <div className="flex gap-2">
        <select
          value={condition.condition.type}
          onChange={(e) => handleChangeCondition(e, index)}
          className="bg-buttonBg text-white-1000 p-4 border-[1px] border-white-200 rounded focus:outline-none focus:border-orange"
        >
          {displayConditionTypes().map((option, optionIndex) => (
            <option
              key={optionIndex}
              value={option.props.value}
              className="bg-buttonBg text-white-1000 p-4 border-[1px] border-white-500"
            >
              {option.props.children}
            </option>
          ))}
        </select>
        {/* <input
          value={condition.value}
          className="bg-buttonBg text-white-1000 p-4 border-[1px] border-white-200 rounded-lg focus:outline-none focus:border-orange"
        ></input> */}
        <CustomDatePicker value={selectedDate} onChange={setSelectedDate} />
        <button
          onClick={() => handleRemoveCondition(index)}
          className="text-orange text-4xl"
        >
          -
        </button>
      </div>
    </div>
  );
};

export default Condition;
