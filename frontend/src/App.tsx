import { useState, useEffect, lazy } from "react";
import { useRecoilValue } from "recoil";
import { bgState } from "./store/atom/atom";

const Navbar = lazy(() => import("./components/Navbar"));
const TerminalBox = lazy(() => import("./components/TerminalBox"));
const BackgroundImageRotator = lazy(() => import("./hooks/bgImageRotator"));

export default function App() {
  return (
    <>
      <BackgroundImageRotator />
      <MyApp />
    </>
  );
}

function MyApp() {
  const bgImage = useRecoilValue(bgState);
  const [fade, setFade] = useState(false);

  useEffect(() => {
    setFade(false);
    const timer = setTimeout(() => {
      setFade(true);
    }, 50);

    return () => clearTimeout(timer);
  }, [bgImage]);

  return (
    <div
      className={`w-full h-screen bg-cover bg-center overflow-auto transition-opacity duration-1000 ${fade ? "opacity-100" : "opacity-0"}`}
      style={{ backgroundImage: `url(${bgImage})` }}
    >
      <Navbar />
      <div className="w-[full] h-[95%] mt-[20px] flex justify-center items-center p-[20px]">
        <TerminalBox />;
      </div>
    </div>
  );
}
