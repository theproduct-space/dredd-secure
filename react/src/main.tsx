// React Imports
import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";

// Contexts Imports
import AddressProvider from "./def-hooks/addressContext";
import DenomProvider from "./def-hooks/denomContext";
import WalletProvider from "./def-hooks/walletContext";

// Custom Imports
import router from "./router";

// 3rd Party Imports
import { Analytics } from "@vercel/analytics/react";

// Styles Imports
import "@ignt/react-library/dist/style.css";
import "./index.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { TextField, ThemeProvider, createTheme } from "@mui/material";

const queryClient = new QueryClient();

const theme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#FF8A00",
    },
  },
});

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AddressProvider>
        <WalletProvider>
          <DenomProvider>
            <ThemeProvider theme={theme}>
              <RouterProvider router={router} />
            </ThemeProvider>
            <Analytics />
          </DenomProvider>
        </WalletProvider>
      </AddressProvider>
    </QueryClientProvider>
  </React.StrictMode>,
);
