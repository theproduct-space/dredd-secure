import Typography from "~baseComponents/Typography";

const NotFound = () => {
  return (
    <div className="h-screen w-100 bg-black bg-no-repeat bg-center bg-cover flex items-center justify-center">
      <Typography
        variant="h1"
        className="text-white-1000 font-revalia text-orange"
      >
        404 Not Found
      </Typography>
    </div>
  );
};

export default NotFound;
