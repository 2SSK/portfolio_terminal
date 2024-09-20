import { useState, useEffect, memo } from "react";

import { useRecoilValue } from "recoil";
import { inputState } from "../store/atom/atom";

import { WhoAmI } from "./WhoAmI";
import Help from "./Help";
import About from "./About";
import Project from "./Project";
import Social from "./Social";

const TerminalOutput = memo(() => {
  const inputData = useRecoilValue(inputState);

  const [output, setOutput] = useState<string>("");

  useEffect(() => {
    setOutput(inputData);
  }, [inputData]);

  type Command = "whoami" | "clear" | "about" | "projects" | "social" | "help";

  const CommandMap: Record<Command, JSX.Element> = {
    whoami: <WhoAmI />,
    clear: <span>clear</span>,
    about: <About />,
    projects: <Project />,
    social: <Social />,
    help: <Help />,
  };

  const renderOutput = () => {
    const commandKey = output.toLowerCase() as keyof typeof CommandMap;

    if (commandKey in CommandMap) {
      return CommandMap[commandKey];
    } else {
      return <span>{output}: command not found</span>;
    }
  };

  return <div className="text-foreground ml-4 mb-6">{renderOutput()}</div>;
});

export default TerminalOutput;
