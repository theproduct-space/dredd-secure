import { Link } from "react-router-dom";
import DreddLogo from "../../../assets/Dredd-logo.png";
import {
  URL_MEDIUM,
  URL_PRODUCTSHOP,
  URL_TWITTER,
  URL_WHITEPAPER,
  URL_WAITLISTFORM,
} from "~utils/urls";
import { Typography } from "~baseComponents/Typography";

const Footer = () => {
  return (
    <footer className="w-full bg-black z-50 relative pb-3">
      <div className="max-w-app-max mx-auto px-3 lg:px-6">
        <div className="flex justify-between flex-col items-start border-b border-white-200 pb-3 sm:py-8 sm:flex-row sm:items-center">
          <Link
            to="/"
            className="logo-link py-7 mb-3 border-b border-white-200 w-full 
                sm:border-none sm:w-fit sm:py-0 sm:mb-0"
          >
            <img src={DreddLogo} alt="Dredd Logo" className="w-40" />
          </Link>

          <div className="flex gap-3 md:gap-8 lg:gap-12 flex-col sm:flex-row">
            <a href={URL_WAITLISTFORM} target="_blank">
              <Typography variant="body-med" className="hover:text-orange">
                Get Started
              </Typography>
            </a>
            <a href={URL_WHITEPAPER} target="_blank">
              <Typography variant="body-med" className="hover:text-orange">
                Whitepaper
              </Typography>
            </a>
            <a href={URL_MEDIUM} target="_blank">
              <Typography variant="body-med" className="hover:text-orange">
                Medium
              </Typography>
            </a>
            <a target="_blank" href={URL_TWITTER} className="flex items-center">
              <Typography variant="body-med" className="hover:text-orange">
                Twitter
              </Typography>
            </a>
          </div>
        </div>
        <div className="flex justify-between items-center pt-3 gap-1">
          <Typography variant="body-small">
            Made with 🧡 by{" "}
            <a
              href={URL_PRODUCTSHOP}
              target="_blank"
              className="hover:text-orange"
            >
              ProductShop
            </a>
          </Typography>
          <Typography variant="small">
            DreddSecure © 2023. All rights reserved.
          </Typography>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
