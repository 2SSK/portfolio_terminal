import { useEffect, memo, useState } from "react";
import { useRecoilValue } from "recoil";
import { inputState } from "../store/atom/atom";

import Help from "./Help";
import About from "./About";
import Project from "./Project";
import Social from "./Social";
import Banner from "./Banner";

const TerminalOutput = memo(() => {
  const inputData = useRecoilValue(inputState);
  const [output, setOutput] = useState<string | JSX.Element[]>([
    <Banner key="banner" />,
  ]);

  useEffect(() => {
    if (inputData) {
      handleCommand(inputData);
    }
  }, [inputData]);

  const handleCommand = (command: string) => {
    const trimmedCommand = command.toLowerCase();

    switch (trimmedCommand) {
      case "clear":
        setOutput(""); // Clear the output
        break;
      case "about":
        setOutput([<About />]);
        break;
      case "projects":
        setOutput([<Project />]);
        break;
      case "social":
        setOutput([<Social />]);
        break;
      case "help":
        setOutput([<Help />]);
        break;
      default:
        setOutput([<span>{command}: command not found</span>]);
        break;
    }
  };

  return <div className="text-foreground ml-4 mb-6">{output}</div>;
});

export default TerminalOutput;
