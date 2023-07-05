import { createBrowserRouter, Outlet } from "react-router-dom";
import Home from "../views/Home";
import Header from "../components/layouts/Header";
import Dashboard from "../views/Dashboard";

const Layout = () => {
  return (
    <div className="relative">
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
      { path: "/app", element: <Dashboard /> },
    ],
  },
]);

export default router;
