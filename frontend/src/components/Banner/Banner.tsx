import Welcome from "./Welcome";
import BannerCommand from "./BannerCommand";

const repoLink = "https://github.com/2SSK/portfolio_terminal";

const Banner: React.FC = () => {
  return (
    <div className="text-sm sm:text-md">
      <Welcome />
      <BannerCommand command="help" /> for a list of all available commands
      <br />
      <BannerCommand command="repo" /> to view the GitHub repository or click{" "}
      <a href={repoLink} target="_blank" className="underline text-customBlue">
        here
      </a>
    </div>
  );
};

export default Banner;
