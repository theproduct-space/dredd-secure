import { Fragment } from "react";
import { Menu, Transition } from "@headlessui/react";
import Typography from "~baseComponents/Typography";
import ChevronDownIcon from "~icons/ChevronDownIcon";

function classNames(...classes: Array<string>) {
  return classes.filter(Boolean).join(" ");
}

export interface DropdownChoice {
  label: string;
  value: string | undefined;
}

export interface DropdownProps {
  choices: DropdownChoice[];
  selectedOption: DropdownChoice;
  setSelectedOption: (option: any) => void;
}

const Dropdown = (props: DropdownProps) => {
  const { choices, selectedOption, setSelectedOption } = props;

  return (
    <Menu as="div" className="relative inline-block text-left w-[120px]">
      <div>
        <Menu.Button className="inline-flex w-full justify-between items-center gap-x-1.5 rounded-lg bg-white px-4 py-2 border border-orange hover:bg-gray-50">
          <Typography variant={"x-small"} className="font-revalia">
            {selectedOption.label}
          </Typography>
          <ChevronDownIcon
            className="-mr-1 h-5 w-5 text-gray-400"
            aria-hidden="true"
          />
        </Menu.Button>
      </div>

      <Transition
        as={Fragment}
        enter="transition ease-out duration-100"
        enterFrom="transform opacity-0 scale-95"
        enterTo="transform opacity-100 scale-100"
        leave="transition ease-in duration-75"
        leaveFrom="transform opacity-100 scale-100"
        leaveTo="transform opacity-0 scale-95"
      >
        <Menu.Items className="absolute right-0 z-10 mt-1 w-full origin-top-right rounded-lg bg-white focus:outline-none border border-orange bg-gray overflow-hidden">
          <div className="">
            {choices.map((choice) => {
              return (
                <Menu.Item key={choice.label}>
                  {({ active }) => (
                    <button
                      className={classNames(
                        active
                          ? "bg-gray-100 text-white-1000"
                          : "text-white-500",
                        "block px-4 py-2 w-full text-sm text-left hover:bg-orange border-b last:border-b-0 border-orange",
                      )}
                      onClick={() => {
                        setSelectedOption(choice);
                      }}
                    >
                      <Typography variant={"x-small"} className="font-revalia">
                        {choice.label}
                      </Typography>
                    </button>
                  )}
                </Menu.Item>
              );
            })}
          </div>
        </Menu.Items>
      </Transition>
    </Menu>
  );
};

export default Dropdown;
