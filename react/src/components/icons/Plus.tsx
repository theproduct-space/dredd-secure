/* eslint-disable @typescript-eslint/no-explicit-any */
import React from "react";

const Plus: React.FC<any> = (props) => {
  return (
    <svg
      width="20"
      height="21"
      viewBox="0 0 20 21"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className="text-white-1000 hover:text-orange"
      {...props}
    >
      <path
        d="M9.75 2.25V18.75"
        stroke="#FAFAFA"
        strokeWidth="3"
        strokeLinecap="round"
      />
      <path
        d="M1.5 10.5L18 10.5"
        stroke="#FAFAFA"
        strokeWidth="3"
        strokeLinecap="round"
      />
    </svg>
  );
};

export default Plus;
