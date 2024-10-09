import Terminal from "./components/Terminal/Terminal";
import wallpaper from "/images/wallpapers/wallpaper.jpg";

const App = () => {
  return (
    <div
      className="w-full h-screen bg-cover bg-center flex items-start justify-center p-4"
      style={{ backgroundImage: `url(${wallpaper})` }}
    >
      <Terminal />
    </div>
  );
};

export default App;
