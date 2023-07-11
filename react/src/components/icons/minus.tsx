import React from "react";

const Minus: React.FC<any> = (props) => {
  return (
    <svg
      width="20"
      height="3"
      viewBox="0 0 20 3"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className="text-white-1000 hover:text-orange"
      {...props}
    >
      <path
        d="M1.5 1.5L18 1.5"
        stroke="#FAFAFA"
        strokeWidth="3"
        strokeLinecap="round"
      />
    </svg>
  );
};

export default Minus;
