import Welcome from "./Welcome";
import BannerCommand from "./BannerCommand";

const Banner: React.FC = () => {
  return (
    <div className="text-sm md:text-md mb-2">
      <Welcome />
      <div className="flex flex-col gap-2">
        <div>
          <BannerCommand command="help" /> for a list of all available commands
        </div>
        <div className="text-wrap">
          <BannerCommand command="repo" /> to visit repository{" "}
        </div>
      </div>
    </div>
  );
};

export default Banner;
