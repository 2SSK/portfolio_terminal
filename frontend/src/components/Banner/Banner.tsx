import Welcome from "./Welcome";
import BannerCommand from "./BannerCommand";

const Banner: React.FC = () => {
  return (
    <div className="mb-2">
      <Welcome />
      <div className="flex flex-col gap-2 text-sm md:text-base">
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
