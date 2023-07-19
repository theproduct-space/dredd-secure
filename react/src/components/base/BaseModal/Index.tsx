import { Modal, Box, Typography, Button } from "@mui/material";
import React from "react";

interface BaseModalProps {
  open: boolean;
  handleClose: () => void;
  title: string;
  children?: React.ReactNode;
}

const BaseModal: React.FC<BaseModalProps> = ({
  open,
  handleClose,
  title,
  children,
}) => {
  const modalBody = (
    <Box
      sx={{
        position: "absolute",
        top: "50%",
        left: "50%",
        transform: "translate(-50%, -50%)",
        width: 400,
        bgcolor: "#241F16",
        border: "1px solid #000",
        borderRadius: "16px",
        boxShadow: 24,
        p: 4,
      }}
    >
      <Typography id="modal-title" variant="h6" className="text-white-1000 ">
        {title}
      </Typography>
      {children} {/* Render children here */}
      <Button onClick={handleClose}>Close</Button>
    </Box>
  );

  return (
    <Modal
      open={open}
      onClose={handleClose}
      aria-labelledby="modal-title"
      aria-describedby="modal-description"
    >
      {modalBody}
    </Modal>
  );
};

export default BaseModal;
