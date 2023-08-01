import * as React from "react";
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { createTheme, ThemeProvider } from "@mui/material/styles";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#FF8A00",
    },
  },
});

interface DatePickerProps {
  value: Date | null;
  onChange: (newValue: Date | null) => void;
}

export default function CustomDatePicker({ value, onChange }: DatePickerProps) {
  return (
    <ThemeProvider theme={darkTheme}>
      <LocalizationProvider dateAdapter={AdapterDayjs}>
        <DatePicker
          value={value}
          onChange={onChange}
          sx={{ borderRadius: "8px", width: "100%" }}
        />
      </LocalizationProvider>
    </ThemeProvider>
  );
}
