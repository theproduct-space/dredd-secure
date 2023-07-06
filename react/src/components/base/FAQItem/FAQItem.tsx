/* eslint-disable import/no-unresolved */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
import React from "react";
import minus from "../../../assets/minusIcon.svg";
import plus from "../../../assets/plusIcon.svg";

interface FAQItemProps {
  question: string;
  answer: string | JSX.Element;
  isOpen: boolean;
  onToggle: () => void;
}

const FAQItem: React.FC<FAQItemProps> = ({ question, answer, isOpen, onToggle }) => {
  return (
    <div className="bg-white-200 rounded-xl p-4">
      <div className="flex justify-between items-center cursor-pointer" onClick={onToggle}>
        <p className="text-p1 text-white-1000">{question}</p>
        <img src={isOpen ? minus : plus} alt="toggle icon" />
      </div>
      {isOpen && <p className="text-white-1000 pt-4">{answer}</p>}
    </div>
  );
};

export default FAQItem;
