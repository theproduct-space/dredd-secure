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
import "./index.css";
import "@ignt/react-library/dist/style.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AddressProvider>
        <WalletProvider>
          <DenomProvider>
            <RouterProvider router={router} />
            <Analytics />
          </DenomProvider>
        </WalletProvider>
      </AddressProvider>
    </QueryClientProvider>
  </React.StrictMode>,
);
