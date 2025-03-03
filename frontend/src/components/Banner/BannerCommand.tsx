const BannerCommand = ({ command }: { command: string }) => {
  return (
    <span className="w-15 md:w-20 text-center break-words  text-primary rounded-md bg-customBlue bg-opacity-15 py-0.5 px-1 md:px-2 mr-1 md:mr-2">
      {command}
    </span>
  );
};

export default BannerCommand;
