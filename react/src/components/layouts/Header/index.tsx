import { IgntLink, IgntLogo } from "@ignt/react-library";
import { Link } from "react-router-dom";
import Account from "~sections/Account";

type MenuItem = {
    label: string;
    to?: string;
    href?: string;
};
interface HeaderProps {
    navHome: MenuItem;
    navPages: MenuItem[]
}
export default function Header(props: HeaderProps) {
    const { navHome, navPages } = props;

    return (
        <header className="flex p-5">
            <Link to={navHome.to || "/"} className="logo-link">
                <div className="logo"></div>
                <IgntLogo className="mx-2.5"></IgntLogo> {/* TODO (Design): Remove that line when you will have the logo. */}
            </Link>
            <nav className="flex flex-1 justify-between">
                <ul>
                    {
                        navPages.map((page) => {
                            return (
                                <Link to={page.to || "/"}>
                                    <div>{page.label}</div>
                                </Link>
                            )
                        })
                    }
                </ul>
                <div>
                    <Account />
                </div>
            </nav>
        </header>
    );
}
