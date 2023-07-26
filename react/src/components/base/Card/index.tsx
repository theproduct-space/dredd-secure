import React from "react";

interface ButtonProps {
  children: React.ReactNode;
  progress?: string;
  className?: string;
}

const Card = ({ children, progress, className }: ButtonProps) => {
  return (
    <div
      className={`relative bg-gray mx-auto rounded-3xl border border-white-200 overflow-hidden ${className}`}
    >
      {progress && (
        <div className={`bg-white-500 w-full h-[24px]`}>
          <div
            className={`bg-orange h-[24px]`}
            style={{ width: `${progress}%` }}
          />
        </div>
      )}

      <div
        className={`relative bg-gray w-full rounded-3xl 
        ${progress && "-top-5"} 
        ${progress && "-mb-5"}
        `}
      >
        {children}
      </div>
    </div>
  );
};

export default Card;
