import React from 'react'

export interface FilterDropDownChoice {
    label: string;
    filterValue: string | undefined;
}

export interface FilterDropDownProps {
    choices: FilterDropDownChoice[];
    filterFunction: (choice: string) => void;
}

export function FilterDropDown(props: FilterDropDownProps) {
    const { choices, filterFunction } = props;

    const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        filterFunction(event.target.value);
    }
    return (
        <select onChange={handleSelectChange}>
            {
                choices.map((choice) => {
                    return (
                        <option value={choice.filterValue} key={choice.label}>{choice.label}</option>
                    )
                })
            }
        </select>
    )
}
