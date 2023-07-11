// React Imports
import { useState, useRef, useEffect } from "react";

// Custom Imports
import Typography from "~baseComponents/Typography";

// Icons Imports
import Minus from "~icons/Minus";
import Plus from "~icons/Plus";

interface FAQItemProps {
  question: string;
  answer: string | JSX.Element;
}

const FAQItem = ({ question, answer }: FAQItemProps) => {
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
      <button
        className="flex justify-between w-full items-center text-left"
        onClick={() => {
          setIsSelected((prev) => !prev);
        }}
      >
        <Typography variant="h6">{question}</Typography>
        {isSelected ? <Minus /> : <Plus />}
      </button>
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
