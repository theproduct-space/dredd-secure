import React from "react";
import Typography from "~baseComponents/Typography";

interface ButtonProps {
  status: string;
  className?: string;
}

const statusStyles = {
  open: {
    fill: "#FFB648",
    border: "#AA7627",
  },
  pending: {
    fill: "#FFEB38",
    border: "#B4A526",
  },
  closed: {
    fill: "#FFF5E7",
    border: "#D0C9BF",
  },
  cancelled: {
    fill: "#FFF5E7",
    border: "#D0C9BF",
  },
};

const StatusPill = ({ status, className }: ButtonProps) => {
  return (
    <div
      className={`inline-flex items-center px-2 border-2 rounded min-h-[32px] ${className}`}
      style={{
        backgroundColor: statusStyles[status].fill,
        borderColor: statusStyles[status].border,
      }}
    >
      <Typography variant="body-small" className="font-semibold text-grayText">
        {status.charAt(0).toUpperCase() + status.slice(1)}
      </Typography>
    </div>
  );
};

export default StatusPill;
