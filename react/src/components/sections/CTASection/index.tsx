// Custom Imports
import Typography from "~baseComponents/Typography";
import Button from "~baseComponents/Button";
import ContentContainer from "~layouts/ContentContainer";
import { URL_WAITLISTFORM } from "~utils/urls";

const CTASection = () => {
  return (
    <ContentContainer className="pt-20 pb-40 md:pt-16 md:pb-52 relative">
      <div className="relative w-full flex justify-center ">
        <div className="flex flex-col z-10 items-center text-center">
          <Typography variant="h2" className="font-revalia pt-6 pb-6 md:pb-12">
            Ready to Secure Your Transactions?
          </Typography>
          <Typography variant="h6" as={"blockquote"} className="max-w-6xl mb-8">
            With DreddSecure, you can enhance trust, security, and protection
            for users in various use cases. Get started with DreddSecure today
            and experience the difference.
          </Typography>
          <a href={URL_WAITLISTFORM} target="_blank" rel="noreferrer">
            <Button
              text="Get Started"
              className="font-revalia rounded-full border-solid border"
            />
          </a>
        </div>
      </div>
    </ContentContainer>
  );
};

export default CTASection;
