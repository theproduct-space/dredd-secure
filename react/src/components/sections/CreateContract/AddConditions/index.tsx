// Custom Imports
import Typography from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import Condition from "~sections/Condition";
// import ICondition from "../../CreateContract";

export interface ICondition {
  type: string;
  prop: string;
  input?: string;
}

export const ConditionTypes: ICondition[] = [
  {
    type: "Starting Date",
    prop: "startDate",
  },
  {
    type: "Deadline",
    prop: "endDate",
  },
  {
    type: "API request",
    prop: "api",
  },
];

interface AddConditionsProps {
  conditions: { condition: ICondition; value: string }[];
  setConditions: (
    newConditions: { condition: ICondition; value: string }[],
  ) => void;
}

const AddConditions = ({ conditions, setConditions }: AddConditionsProps) => {
  const displayConditionTypes = () => {
    return ConditionTypes.map((condition) => {
      return (
        <option key={condition.type} value={condition.type}>
          {condition.type}
        </option>
      );
    });
  };

  const handleAddNewEmptyCondition = () => {
    const array = [...conditions].concat({
      condition: ConditionTypes[0],
      value: "",
    });
    setConditions(array);
  };

  const handleRemoveCondition = (id: number) => {
    const array = conditions.slice(0, id).concat(conditions.slice(id + 1));
    setConditions(array);
  };

  // const handleChangeConditionValue = (
  //   e: React.ChangeEvent<HTMLInputElement>,
  //   id: number,
  // ) => {
  //   const array = [...conditions];
  //   array[id].value = e.target.value;
  //   setConditions(array);
  // };

  const handleChangeCondition = (
    e: React.ChangeEvent<HTMLSelectElement>,
    id: number,
  ) => {
    const array = [...conditions];
    array[id].condition =
      ConditionTypes.find((element) => element.type === e.target.value) ??
      ConditionTypes[0];
    setConditions(array);
  };

  return (
    <>
      <div className="subtitle">
        <Typography variant="h6" className="font-revalia">
          Add Conditions<span className="text-orange">*</span>
        </Typography>
      </div>
      {conditions.map((condition, index) => (
        <div className="condition" key={`add-condition-${index}`}>
          <Condition
            condition={condition}
            index={index}
            handleChangeCondition={handleChangeCondition}
            handleRemoveCondition={handleRemoveCondition}
            displayConditionTypes={displayConditionTypes}
            className="pb-2"
          >
            Condition #{index + 1}
          </Condition>
        </div>
      ))}
      <button className="add-condition" onClick={handleAddNewEmptyCondition}>
        <Typography
          variant="body-small"
          className="font-revalia text-orange py-4"
        >
          + Add Condition
        </Typography>
      </button>
    </>
  );
};

export default AddConditions;
