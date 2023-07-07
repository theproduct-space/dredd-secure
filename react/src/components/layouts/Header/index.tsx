/* eslint-disable import/no-unresolved */
import React from "react";
import { Link, useLocation } from "react-router-dom";
import Account from "~sections/Account";
import DreddLogo from "../../../assets/Dredd-logo.png";
import Button from "~baseComponents/Button";
import TwitterIcon from "~icons/TwitterIcon";
import LockIcon from "~icons/LockIcon";
import { URL_TWITTER } from "~utils/urls";

export default function Header() {
  const location = useLocation();

  return (
    <header className="w-full bg-black z-50 fixed left-0 right-0">
      <div className="h-20 flex justify-between gap-3 items-center max-w-app-max px-3 lg:px-6 mx-auto">
        <Link to="/" className="logo-link">
          <img src={DreddLogo} alt="Dredd Logo" className="w-44" />
        </Link>
        {location.pathname !== "/app" ? (
          <div className="flex items-center gap-3">
            <a target="_blank" href={URL_TWITTER}>
              <TwitterIcon
                width="32"
                height="26"
                className="hover:fill-orange"
              />
            </a>
            {/* <Link to="/app"> */}
            <Button
              text="Lauch App"
              icon={<LockIcon />}
              disabled
              className="font-revalia rounded-full border-solid border-2"
            />
            {/* </Link> */}
          </div>
        ) : (
          <div>
            <Account />
          </div>
        )}
      </div>
    </header>
  );
}
