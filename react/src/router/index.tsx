import { createBrowserRouter, Outlet } from "react-router-dom";
import Home from "../views/Home";
import Header from "../components/layouts/Header";
import FAQ from "../views/FAQ";
import Dashboard from "../views/Dashboard";

const items = [
  {
    label: "Home",
    to: "/",
  },
  {
    label: "DashBoard",
    to: "/app",
  },
  {
    label: "FAQ",
    to: "/faq",
  },
];

const Layout = () => {
  return (
    <>
      <Header navItems={items} />
      <Outlet />
    </>
  );
};
const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { path: "/", element: <Home /> },
      { path: "/app", element: <Dashboard /> },
      { path: "/faq", element: <FAQ /> },
    ],
  },
]);

export default router;
