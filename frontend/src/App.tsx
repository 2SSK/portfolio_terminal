import Terminal from "./components/Terminal/Terminal";
import wallpaper from "/images/wallpapers/wallpaper.jpg";

import { WhoAmI } from "./components/WhoAmI";
import Tools from "./components/Tools/Tools";
import Project from "./components/Project/Project";
import Social from "./components/Social";

const App = () => {
  return (
    <div
      className="w-full h-screen bg-cover bg-center flex items-start justify-center overflow-auto md:p-4"
      style={{ backgroundImage: `url(${wallpaper})` }}
    >
      <Terminal className="hidden md:block" />
      <div className="block md:hidden">
        <div className="bg-[#011423] bg-opacity-50 backdrop-blur-xl flex flex-col gap-4 p-4">
          <WhoAmI />
          <Tools />
          <Project />
          <Social />
        </div>
      </div>
    </div>
  );
};

export default App;
