/* eslint-disable import/no-unresolved */
/* eslint-disable jsx-a11y/click-events-have-key-events */
/* eslint-disable jsx-a11y/no-static-element-interactions */
import React, { useState, useRef, useEffect } from "react";
import { Typography } from "~baseComponents/Typography";
import Plus from "~icons/Plus";
import Minus from "~icons/minus";

interface FAQItemProps {
  question: string;
  answer: string | JSX.Element;
}

const FAQItem: React.FC<FAQItemProps> = ({ question, answer }) => {
  const [isSelected, setIsSelected] = useState(false);
  const [maxHeight, setMaxHeight] = useState("");
  const heightRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (heightRef.current) {
      setMaxHeight(
        isSelected ? `${heightRef?.current?.scrollHeight}px` : "0px",
      );
    }
  }, [isSelected]);

  return (
    <div className="bg-white-200 rounded-xl p-4">
      <div
        className="flex justify-between items-center cursor-pointer"
        onClick={() => {
          setIsSelected((prev) => !prev);
        }}
      >
        <Typography variant="h6">{question}</Typography>
        {isSelected ? <Minus /> : <Plus />}
      </div>
      <div
        ref={heightRef}
        style={{ maxHeight }}
        className={` transition-max-height overflow-hidden duration-100 ease-in`}
      >
        <Typography variant="body-small" as="p" className="pt-4">
          {answer}
        </Typography>
      </div>
    </div>
  );
};

export default FAQItem;
