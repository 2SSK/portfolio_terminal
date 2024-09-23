import { lazy, useEffect } from "react";

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
  useEffect(() => {
    const img = new Image();
    img.src = wallpaper;
  }, []);

  return (
    <div
      className={`w-full h-screen bg-cover bg-center overflow-hidden bg-gray-900`}
      style={{ backgroundImage: `url(${wallpaper})` }}
    >
      <Navbar />
      <div className="w-full h-[100%] mt-2  flex justify-center items-center p-[20px]">
        <TerminalBox />
      </div>
    </div>
  );
}
