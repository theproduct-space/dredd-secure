// Custom Imports
import Typography from "~baseComponents/Typography";
import Condition from "~sections/Condition";
// import ICondition from "../../CreateContract";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";
import { CoinMarketCapTokenI } from "~sections/Condition/CoinMarketCapTokenSelector";

export interface ISubConditions {
  conditionType: string; // gt, lt, equal
  dataType: string; // number, string
  name: string; // relevant_fields.name
  path: Array<string>; // path to data, from API call
  label: string; // relevant_fields.label
  value: string | number | undefined; // input from user
}

export interface ICondition {
  label: string;
  name: string;
  value?: string | number;
  tokenOfInterest?: CoinMarketCapTokenI;
  subConditions?: Array<ISubConditions>;
}

export const ConditionTypes: ICondition[] = [
  {
    label: "Starting Date",
    name: "startDate",
  },
  {
    label: "Deadline",
    name: "endDate",
  },
  ...configuredAPIEndpoints.list.map((endpoint) => ({
    label: configuredAPIEndpoints.data[endpoint].label,
    name: endpoint,
    subConditions: [
      {
        conditionType: "eq",
        dataType: configuredAPIEndpoints.data[endpoint].relevant_fields[0].type,
        name: configuredAPIEndpoints.data[endpoint].relevant_fields[0].name,
        path: [""],
        label: configuredAPIEndpoints.data[endpoint].relevant_fields[0].label,
        value: undefined,
      },
    ],
  })),
];

interface AddConditionsProps {
  conditions: ICondition[];
  // setConditions: (newConditions: ICondition[]) => void;
  setConditions: React.Dispatch<React.SetStateAction<ICondition[]>>;
}

const AddConditions = ({ conditions, setConditions }: AddConditionsProps) => {
  const displayConditionTypes = () => {
    return ConditionTypes.map((condition) => {
      return (
        <option key={condition.label} value={condition.label}>
          {condition.label}
        </option>
      );
    });
  };

  const handleAddNewEmptyCondition = () => {
    const array = [...conditions].concat({
      label: "Select Condition Type",
      name: "select",
      value: "",
    });
    setConditions(array);
  };

  const handleRemoveCondition = (index: number) => {
    const array = conditions
      .slice(0, index)
      .concat(conditions.slice(index + 1));
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

  const handleSelectCondition = (
    e: React.ChangeEvent<HTMLSelectElement>,
    index: number,
  ) => {
    const array = [...conditions];
    array[index] =
      ConditionTypes.find((element) => element.label === e.target.value) ??
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
            setConditions={setConditions}
            index={index}
            handleSelectCondition={handleSelectCondition}
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
