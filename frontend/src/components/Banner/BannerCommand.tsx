const BannerCommand = ({ command }: { command: string }) => {
  return (
    <span className="w-20 text-center text-primary rounded-md bg-customBlue bg-opacity-15 py-0.5 px-2 mr-2">
      {command}
    </span>
  );
};

export default BannerCommand;
