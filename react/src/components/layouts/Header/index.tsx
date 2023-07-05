/* eslint-disable import/no-unresolved */
import React from "react";
import { Link, useLocation } from "react-router-dom";
import Account from "~sections/Account";
import DreddLogo from "../../../assets/Dredd-logo.png";
import Button from "~baseComponents/Button";

export default function Header() {
  const location = useLocation();

  return (
    <header className="bg-black h-20 flex justify-between items-center px-4 sticky top-0 z-50">
      <Link to="/" className="logo-link">
        <img src={DreddLogo} alt="Dredd Logo" className="mx-2.5 w-44" />
      </Link>
      {location.pathname !== "/app" ? (
        <div>
          <Link to="/app">
            <Button text="Lauch App" className="font-revalia rounded-full border-solid border-2" />
          </Link>
        </div>
      ) : (
        <div>
          <Account />
        </div>
      )}
    </header>
  );
}
