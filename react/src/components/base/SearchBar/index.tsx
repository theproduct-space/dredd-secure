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
}

const SearchBar: React.FC<SearchBarProps> = ({ placeholder, onChange }) => {
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
      </div>
    </ThemeProvider>
  );
};

export default SearchBar;
