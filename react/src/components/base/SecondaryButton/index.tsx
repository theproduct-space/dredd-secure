import Button from "~baseComponents/Button";

interface SecondaryButtonProps {
  onClick?: () => void;
  text: string;
  secondary?: boolean;
  className?: string;
  icon?: JSX.Element;
  orangeText?: boolean;
}

const SecondaryButton = ({
  onClick,
  text,
  icon,
  secondary,
  orangeText,
  className,
}: SecondaryButtonProps) => {
  return (
    <Button
      text={text}
      className={`${`p-2 border rounded-3xl border-white-500 bg-buttonBg capitalize w-max ${className}`}`}
      orangeText={orangeText}
      icon={icon}
      secondary={secondary}
      onClick={onClick}
    />
  );
};

export default SecondaryButton;
