import React from "react";
import TextField from "@mui/material/TextField";
import { createTheme, ThemeProvider } from "@mui/material/styles";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#FF8A00",
    },
  },
});

interface SearchBarProps {
  placeholder: string;
  onChange: (query: string) => void;
  children?: React.ReactNode;
}

const SearchBar: React.FC<SearchBarProps> = ({
  placeholder,
  onChange,
  children,
}) => {
  return (
    <ThemeProvider theme={darkTheme}>
      <div className="relative">
        <TextField
          type="search"
          placeholder={placeholder}
          onChange={(e) => onChange(e.target.value)}
          variant="outlined"
          size="small"
          fullWidth
          className="border-white-200"
        />
        {children}
      </div>
    </ThemeProvider>
  );
};

export default SearchBar;
