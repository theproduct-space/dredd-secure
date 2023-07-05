import React from 'react'

export interface FilterDropDownChoice {
    label: string;
    filterValue: string;
}

export interface FilterDropDownProps {
    choices: FilterDropDownChoice[];
    filterFunction: (choice: string) => void;
}

export function FilterDropDown(props: FilterDropDownProps) {
    const { choices, filterFunction } = props;

    const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const { value } = event.target;
        filterFunction(value);
    }
    return (
        <select onChange={handleSelectChange}>
            <option>All</option>
            {
                choices.map((choice) => {
                    return (
                        <option value={choice.filterValue}>{choice.label}</option>
                    )
                })
            }
        </select>
    )
}
