import { Modal, Box, Typography, Button } from "@mui/material";
import React from "react";

interface BaseModalProps {
  open: boolean;
  handleClose: () => void;
  children?: React.ReactNode;
}

const BaseModal: React.FC<BaseModalProps> = ({
  open,
  handleClose,
  children,
}) => {
  const modalBody = (
    <Box
      sx={{
        position: "absolute",
        top: "50%",
        left: "50%",
        transform: "translate(-50%, -50%)",
        width: 500,
        height: 500,
        bgcolor: "#241F16",
        border: "1px solid #000",
        borderRadius: "16px",
        boxShadow: 24,
        overflow: "hidden",
      }}
    >
      <Box
        sx={{
          width: "100%",
          height: "100%",
          overflowY: "scroll",
          "&::-webkit-scrollbar": {
            width: "10px",
            background: "#F5F5F5",
          },
          "&::-webkit-scrollbar-thumb": {
            background: "#FF8A00",
            borderRadius: "4px",
          },
        }}
      >
        {children}
      </Box>
    </Box>
  );

  return (
    <Modal open={open} onClose={handleClose}>
      {modalBody}
    </Modal>
  );
};

export default BaseModal;
