import { lazy, Suspense } from "react";

// import { useRecoilState } from "recoil";
// import { inputState } from "../store/atom/atom";

const TerminalInput = lazy(() => import("./TerminalInput"));
// const TerminalOutput = lazy(() => import("./TerminalOutput"));

// import { WhoAmI } from "./WhoAmI";
// import Help from "./Help";

export default function TerminalBox() {
  // const [history, setHistory] = useState<
  //   { input: string; output: JSX.Element }[]
  // >([]);
  // const [currentInput, setCurrentInput] = useRecoilState(inputState);

  // const handleCommandExecution = (command: string) => {
  //   const lowerCommand = command.toLowerCase();
  //
  //   switch (lowerCommand) {
  //     case "clear":
  //       setHistory([]); // Clear the terminal
  //       break;
  //     case "whoami":
  //       setHistory((prev) => [...prev, { input: command, output: <WhoAmI /> }]);
  //       break;
  //     case "help":
  //       setHistory((prev) => [
  //         ...prev,
  //         { input: command, output: <span>Help content here...</span> },
  //       ]);
  //       break;
  //     default:
  //       setHistory((prev) => [
  //         ...prev,
  //         { input: command, output: <span>{command}: command not found</span> },
  //       ]);
  //   }
  //   setCurrentInput("");
  // };

  return (
    <div className="w-[80%] h-[90%] p-7 border-4 border-customBlue rounded-lg bg-transparent bg-opacity-30 backdrop-blur-sm shadow-lg shadow-disabled font-jetbrains overflow-y-auto">
      {/* {history.map((entry: any, index: number) => ( */}
      {/*   <div key={index}> */}
      {/*     <div className="text-textColor font-bold mb-1">ssk@archBTW ~</div> */}
      {/*     <div className="flex"> */}
      {/*       <span className="text-primary font-bold mr-2">&gt;</span> */}
      {/*       <span>{entry.input}</span> */}
      {/*     </div> */}
      {/*     <div className="ml-4 mb-6">{entry.output}</div> */}
      {/*   </div> */}
      {/* ))} */}

      {/* Render new input box */}
      <Suspense fallback={<div>Loading...</div>}>
        <TerminalInput />
      </Suspense>
    </div>
  );
}
