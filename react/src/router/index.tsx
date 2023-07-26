// React Imports
import { createBrowserRouter, Outlet } from "react-router-dom";
import { ToastContainer } from "react-toastify";

import 'react-toastify/dist/ReactToastify.css';

// Custom Imports
import Header from "~layouts/Header";
import CreateContractPage from "~views/CreateContract";
import Dashboard from "~sections/Dashboard";
import Landing from "~views/Home";
import PaymentView from "~views/Payment";
import ReviewContract from "~views/ReviewContract";
import Footer from "~layouts/Footer";

const home = {
    label: "Landing",
    to: "/",
    view: <Landing />,
};
const otherPages = [
    {
        label: "Dashboard",
        to: "/dashboard",
        view: <Dashboard />,
    },
    {
        label: "CreateContract",
        to: "/escrow/create",
        view: <CreateContractPage />,
    },
    {
        label: "PayEscrow",
        to: "/escrow/pay",
        view: <PaymentView />,
    },
    {
        label: "ViewEscrow",
        to: "/escrow/:id",
        view: <ReviewContract />,
    },
];
const allPages = otherPages.concat(home);

const Layout = () => {
    return (
        <div className="relative overflow-hidden bg-black">
        <Header />
            <ToastContainer
                position="top-right"
                autoClose={2000}
                hideProgressBar={false}
                newestOnTop={false}
                rtl={false}
                pauseOnFocusLoss={false}
                draggable={false}
                theme="dark"
            />
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
