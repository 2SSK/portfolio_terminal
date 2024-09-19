import Navbar from "./components/Navbar";
import TerminalBox from "./components/TerminalBox";

export default function App() {
  return (
    <>
      <MyApp />
    </>
  );
}

function MyApp() {
  return (
    <div className="w-full h-screen bg-background-light overflow-auto">
      <Navbar />
      <div className="w-full h-[95%] mt-[20px] flex justify-center items-center p-[20px]">
        <TerminalBox />;
      </div>
    </div>
  );
}
