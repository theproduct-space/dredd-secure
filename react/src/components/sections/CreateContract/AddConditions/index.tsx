// Custom Imports
import Typography from "~baseComponents/Typography";
import Condition from "~sections/Condition";
import configuredAPIEndpoints from "~utils/configuredApiEndpoints.json";
import { CoinMarketCapTokenI } from "~sections/Condition/CoinMarketCapTokenSelector";

export interface ISubCondition {
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
  type: string;
  value?: string | number;
  tokenOfInterest?: CoinMarketCapTokenI;
  subConditions?: Array<ISubCondition>;
}

export const ConditionTypes: ICondition[] = [
  {
    label: "Starting Date",
    name: "startDate",
    type: "startDate",
  },
  {
    label: "Deadline",
    name: "endDate",
    type: "endDate",
  },
  ...configuredAPIEndpoints.list.map((endpoint) => ({
    label: configuredAPIEndpoints.data[endpoint].label,
    name: endpoint,
    type: "apiCondition",
    subConditions: [
      {
        conditionType: "eq",
        dataType: configuredAPIEndpoints.data[endpoint].relevant_fields[0].type,
        name: configuredAPIEndpoints.data[endpoint].relevant_fields[0].name,
        path: [
          "data",
          "1",
          "quote",
          "USD",
          configuredAPIEndpoints.data[endpoint].relevant_fields[0].name,
        ],
        label: configuredAPIEndpoints.data[endpoint].relevant_fields[0].label,
        value: undefined,
      },
    ],
  })),
];

interface AddConditionsProps {
  conditions: ICondition[];
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
      type: "select",
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

  const handleSelectCondition = (
    e: React.ChangeEvent<HTMLSelectElement>,
    index: number,
  ) => {
    const array = [...conditions];

    console.log("ConditionTypes", ConditionTypes);

    // Deep copy using spread operator for objects and arrays
    const selectedCondition = JSON.parse(
      JSON.stringify(
        ConditionTypes.find((element) => element.label === e.target.value) ??
          ConditionTypes[0],
      ),
    );

    array[index] = selectedCondition;

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
