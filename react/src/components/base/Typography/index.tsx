import React, { ElementType } from "react";

type Variant =
  | "h1"
  | "h2"
  | "h3"
  | "h4"
  | "h5"
  | "h6"
  | "body"
  | "body-med"
  | "body-small"
  | "small";

interface Props {
  variant: Variant;
  children: React.ReactNode;
  className?: string;
  as?: ElementType;
}

const tags: Record<Variant, ElementType> = {
  h1: "h1",
  h2: "h2",
  h3: "h3",
  h4: "h4",
  h5: "h5",
  h6: "h6",
  body: "p",
  "body-med": "p",
  "body-small": "p",
  small: "span",
};

const sizes: Record<Variant, string> = {
  h1: "text-4xl sm:text-6xl md:text-7xl",
  h2: "text-3xl sm:text-5xl md:text-6xl",
  h3: "text-2xl sm:text-5xl md:text-5xl",
  h4: "text-xl sm:text-3xl md:text-4xl",
  h5: "text-lg sm:text-2xl md:text-3xl",
  h6: "text-base sm:text-xl md:text-2xl",
  body: "text-lg sm:text-2xl",
  "body-med": "text-base sm:text-lg",
  "body-small": "text-sm sm:text-base",
  small: "text-xxs sm:text-xs",
};

const Typography = ({ variant, children, className, as }: Props) => {
  const sizeClasses = sizes[variant];
  const Tag = as || tags[variant];

  return (
    <Tag className={`${sizeClasses} text-white-1000 ${className}`}>
      {children}
    </Tag>
  );
};

export default Typography;
