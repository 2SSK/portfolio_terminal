import { useEffect, memo, useState, useCallback } from "react";
import { useRecoilValue } from "recoil";
import { inputState } from "../store/atom/atom";

import Help from "./Help";
import About from "./About";
import Project from "./Project";
import Social from "./Social";
import Banner from "./Banner";
import Resume from "./Resume";
import WhoAmI from "./WhoAmI";

const REPO_LINK = "https://github.com/2SSK/i3_theme_portfolio";

const TerminalOutput = memo(() => {
  const inputData = useRecoilValue(inputState);
  const [output, setOutput] = useState<(string | JSX.Element)[]>([
    <Banner key="banner" />,
  ]);

  const handleCommand = useCallback((command: string) => {
    const trimmedCommand = command.toLowerCase();

    setOutput((prevOutput) => {
      const newOutput = [...prevOutput];

      newOutput.push(
        <div key={newOutput.length} className="mb-2">
          <span className="text-textColor font-bold">ssk@archBTW ~ </span>
          <span className="text-primary">&gt; {command}</span>
        </div>,
      );

      switch (trimmedCommand) {
        case "clear":
          return [];
        case "whoami":
          newOutput.push(<WhoAmI key={newOutput.length + 1} />);
          break;
        case "about":
          newOutput.push(<About key={newOutput.length + 1} />);
          break;
        case "project":
          newOutput.push(<Project key={newOutput.length + 1} />);
          break;
        case "social":
          newOutput.push(<Social key={newOutput.length + 1} />);
          break;
        case "help":
          newOutput.push(<Help key={newOutput.length + 1} />);
          break;
        case "resume":
          newOutput.push(<Resume key={newOutput.length + 1} />);
          break;
        case "banner":
          newOutput.push(<Banner key={newOutput.length + 1} />);
          break;
        case "repo":
          newOutput.push("Redirecting to github.com...");
          setTimeout(() => {
            window.open(REPO_LINK, "_blank");
          }, 500);
          break;
        default:
          newOutput.push(
            <div key={newOutput.length + 1} className="text-red-500">
              {command}: command not found
            </div>,
          );
          break;
      }
      return newOutput;
    });
  }, []);

  useEffect(() => {
    if (inputData) {
      handleCommand(inputData); // Process the command
    }
  }, [inputData, handleCommand]);

  return (
    <div className="text-foreground ml-4 mb-6">
      {output.map((item, index) => (
        <div key={index}>{item}</div>
      ))}
    </div>
  );
});

export default TerminalOutput;
