import { createBrowserRouter, Outlet } from "react-router-dom";
import Home from "../views/Home";
import Header from "~layouts/Header";
import CreateContractPage from "~views/CreateContract";
import Dashboard from "~views/Dashboard";
import FAQ from "~views/FAQ";
import Landing from "~views/Home";
import PaymentView from "~views/Payment";
import ReviewContract from "~views/ReviewContract";
import Footer from "~layouts/Footer";

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
    },
    {
        label: "ViewEscrow",
        to: "/escrow/:id",
        view: <ReviewContract />
    },
];
const allPages = otherPages.concat(home).concat(navPages);

const Layout = () => {
    return (
        <div className="relative overflow-hidden">
            <Header />
            <Outlet />
            <Footer />
        </div>
    );
};

const router = createBrowserRouter([
    {
        path: "/",
        element: <Layout />,
        children: allPages.map((page) => {
            return { path: page.to, element: page.view };
        }),
    },
]);

export default router;
