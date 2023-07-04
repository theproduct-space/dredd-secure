import { createBrowserRouter, Outlet } from "react-router-dom";
import Home from "../components/layouts/Home";
import DataView from "../views/DataView";
import Header from "../components/layouts/Header";

const items = [
  {
    label: "Home",
    to: "/",
  },
  {
    label: "App",
    to: "/app",
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
      { path: "/app", element: <DataView /> },
    ],
  },
]);

export default router;
