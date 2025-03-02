import Welcome from "./Welcome";
import BannerCommand from "./BannerCommand";

const repoLink = "https://github.com/2SSK/portfolio_terminal";

const Banner: React.FC = () => {
  return (
    <div className="text-sm sm:text-md">
      <Welcome />
      <div className="flex flex-col gap-2">
        <div>
          <BannerCommand command="help" /> for a list of all available commands
        </div>
        <div>
          <BannerCommand command="repo" /> to view the GitHub repository or
          click{" "}
          <a
            href={repoLink}
            target="_blank"
            className="underline text-customBlue"
          >
            here
          </a>
        </div>
      </div>
    </div>
  );
};

export default Banner;
