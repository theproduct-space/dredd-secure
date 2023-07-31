// React imports
import { Link, useLocation } from "react-router-dom";

// Custom Imports
import Account from "~sections/Account";
import { URL_TWITTER } from "~utils/urls";

// Assets Imports
import DreddLogo from "~assets/Dredd-logo.png";
import TwitterIcon from "~icons/TwitterIcon";

const Header = () => {
  const location = useLocation();

  return (
    <header className="w-full bg-gray z-50 fixed left-0 right-0">
      <div className="h-20 flex justify-between gap-3 items-center max-w-app-max px-3 lg:px-6 mx-auto">
        <Link to="/" className="logo-link">
          <img src={DreddLogo} alt="Dredd Logo" className="w-44" />
        </Link>
        {location.pathname !== "/" ? (
          <div>
            <Account />
          </div>
        ) : (
          <div className="flex items-center gap-3">
            <a target="_blank" href={URL_TWITTER} rel="noreferrer">
              <TwitterIcon
                width="32"
                height="26"
                className="hover:fill-orange"
              />
            </a>
            {/* <Link to="/app"> */}
            {/* <Button
            text="Lauch App"
            icon={<LockIcon />}
            disabled
            className="font-revalia rounded-full border-solid border-2"
          /> */}
            {/* </Link> */}
          </div>
        )}
      </div>
    </header>
  );
};

export default Header;
