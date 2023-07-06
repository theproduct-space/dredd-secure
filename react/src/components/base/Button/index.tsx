import React from "react";

interface ButtonProps {
  onClick?: () => void;
  text: string;
  className?: string;
}

const Button: React.FC<ButtonProps> = ({ onClick, text, className }) => {
  return (
    <button
      className={`bg-orange hover:bg-yellow focus:translate-y-0.5 text-white-1000 py-2 px-4 rounded ${className}`}
      onClick={onClick}
    >
      {text}
    </button>
  );
};

export default Button;
