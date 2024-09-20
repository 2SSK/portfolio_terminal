import { lazy } from "react";

const Navbar = lazy(() => import("./components/Navbar"));
const TerminalBox = lazy(() => import("./components/TerminalBox"));

import wallpaper from "/images/wallpapers/wallpaper.jpg";

export default function App() {
  return (
    <>
      <MyApp />
    </>
  );
}

function MyApp() {
  return (
    <div
      className={`w-full h-screen bg-cover bg-center overflow-auto `}
      style={{ backgroundImage: `url(${wallpaper})` }}
    >
      <Navbar />
      <div className="w-[full] h-[95%] mt-[20px] flex justify-center items-center p-[20px]">
        <TerminalBox />;
      </div>
    </div>
  );
}
