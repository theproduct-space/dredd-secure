import { createBrowserRouter, Outlet } from "react-router-dom";
import Home from "../views/Home";
import Header from "~layouts/Header";
import CreateContractPage from "~views/CreateContract";
import Dashboard from "~views/Dashboard";
import FAQ from "~views/FAQ";
import Landing from "~views/Home";
import PaymentView from "~views/Payment";

const home = {
  label: "Landing",
  to: "/",
  view: <Landing />,
};
const navPages = [
  {
    label: "FAQ",
    to: "/faq",
    view: <FAQ />,
  },
];
const otherPages = [
  {
      label: "Dashboard",
      to: "/dashboard",
      view: <Dashboard />
  },
  {
      label: "CreateContract",
      to: "/escrow/create",
      view: <CreateContractPage />
  },
  {
      label: "PayEscrow",
      to: "/escrow/pay",
      view: <PaymentView />
  }
];
const allPages = otherPages.concat(home).concat(navPages);

const Layout = () => {
  return (
    <div className="relative overflow-hidden">
      <Header />
      <Outlet />
    </div>
  );
};

const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { path: "/", element: <Home /> },
      // { path: "/app", element: <Dashboard /> },
    ],
  },
//   {
//       path: "/",
//       element: <Layout />,
//       children: allPages.map((page) => {
//           return { path: page.to, element: page.view };
//       }),
//   },
]);

export default router;
