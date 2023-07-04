/* eslint-disable import/no-unresolved */
import { Link } from "react-router-dom";
import Account from "~sections/Account";
import DreddLogo from "../../../assets/Dredd-logo.png";

type MenuItem = {
  label: string;
  to?: string;
  href?: string;
};
interface HeaderProps {
  navItems: Array<MenuItem>;
}
export default function Header(props: HeaderProps) {
  const { navItems } = props;

  return (
    <header className="bg-gray h-20 flex justify-center items-center">
      <Link to="/" className="logo-link">
        <img src={DreddLogo} alt="Dredd Logo" className="mx-2.5 w-32" />
      </Link>
      <nav className="flex flex-1 justify-between">
        <ul className="flex">
          {navItems.map((page, index) => {
            return (
              <Link key={index} to={page.to || "/"}>
                <div>{page.label}</div>
              </Link>
            );
          })}
        </ul>
        <div>
          <Account />
        </div>
      </nav>
    </header>
  );
}
