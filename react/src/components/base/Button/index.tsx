import Typography from "~baseComponents/Typography";

interface ButtonProps {
  onClick?: () => void;
  text: string;
  secondary?: boolean;
  disabled?: boolean;
  className?: string;
  icon?: JSX.Element;
  orangeText?: boolean;
}

const Button = ({
  onClick,
  text,
  disabled,
  icon,
  secondary,
  orangeText,
  className,
}: ButtonProps) => {
  return (
    <button
      className={`${
        secondary
          ? "bg-transparent hover:bg-orange "
          : "bg-orange hover:bg-transparent "
      } 
        ${disabled && "opacity-50 pointer-events-none"} 
        transition duration-300 ease-in-out hover:-translate-y-0.5 active:-translate-y-0.5
        text-white-1000 py-3 px-5 rounded uppercase flex gap-3 items-center justify-center sm:py-3 sm:px-6 ${className}`}
      onClick={onClick}
    >
      {icon}
      <Typography
        variant="body-small"
        as="div"
        className={`${orangeText && "text-orange"}`}
      >
        {text}
      </Typography>
    </button>
  );
};

export default Button;
