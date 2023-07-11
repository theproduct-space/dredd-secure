import React from "react";

interface ContentContainerProps {
  children?: React.ReactNode;
  className?: string;
}

const ContentContainer = ({ children, className }: ContentContainerProps) => {
  return (
    <div
      className={`w-full max-w-app-max px-4 md:px-8 xl:px-16 mx-auto ${className}`}
    >
      {children}
    </div>
  );
};

export default ContentContainer;
