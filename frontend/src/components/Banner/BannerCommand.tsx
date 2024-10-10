const BannerCommand = ({ command }: { command: string }) => {
  return (
    <span className="text-primary rounded-md bg-customBlue bg-opacity-15">
      '{command}'
    </span>
  );
};

export default BannerCommand;
