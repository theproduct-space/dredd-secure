import React from "react";

interface ModalContainerProps {
  children?: React.ReactNode;
  className?: string;
}

const ModalContainer = ({ children, className }: ModalContainerProps) => {
  return (
    <div className={`max-w-4xl px-4 md:px-8 xl:px-16 mx-auto ${className}`}>
      {children}
    </div>
  );
};

export default ModalContainer;
