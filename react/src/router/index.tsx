import { createBrowserRouter, Outlet } from "react-router-dom";

import Header from "~layouts/Header";
import CreateContractPage from "~views/CreateContract";
import Dashboard from "~views/Dashboard";
import FAQ from "~views/FAQ";
import Landing from "~views/Landing";

const home = {
    label: "Landing",
    to: "/",
    view: <Landing />
};
const navPages = [
    {
        label: "FAQ",
        to: "/faq",
        view: <FAQ />
    }
]
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
    }
];
const allPages = otherPages.concat(home).concat(navPages);

const Layout = () => {
    return (
        <>
            <Header navHome={home} navPages={navPages} />
            <Outlet />
        </>
    );
};
const router = createBrowserRouter([
    {
        path: "/",
        element: <Layout />,
        children: allPages.map((item) => {
            return { path: item.to, element: item.view };
        }),
    },
]);

export default router;
