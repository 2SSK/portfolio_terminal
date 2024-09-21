import { useState, useEffect, useRef } from "react";
import { useSetRecoilState, useRecoilValue } from "recoil";
import { inputState, focusInputState } from "../store/atom/atom";

export default function TerminalInput({
  inputRef,
}: {
  inputRef: React.RefObject<HTMLInputElement>;
}) {
  return (
    <div className="w-full flex flex-col items-start justify-center text-lg">
      {" "}
      <p className="text-textColor font-bold mb-1">
        ssk<span className="text-foreground">@</span>
        archBTW ~
      </p>
      <div className="flex mb-6">
        <span className="text-primary text-lg font-bold mr-2">&gt;</span>
        <TerminalInputBox inputRef={inputRef} />
      </div>
    </div>
  );
}

const TerminalInputBox = ({
  inputRef,
}: {
  inputRef: React.RefObject<HTMLInputElement>;
}) => {
  const setInputData = useSetRecoilState(inputState);
  const focusInput = useRecoilValue(focusInputState);
  const [inputValue, setInputValue] = useState("");
  const [commandHistory, setCommandHistory] = useState<string[]>([]);
  const [tempInput, setTempInput] = useState("");
  const historyIndex = useRef(-1);

  useEffect(() => {
    inputRef.current?.focus();
  }, [inputRef]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    switch (e.key) {
      case "Enter":
        if (inputValue.trim() !== "") {
          setCommandHistory((prev) => [...prev, inputValue.trim()]);
          setInputData(inputValue.trim());
          setInputValue(""); 
          historyIndex.current = -1;
          setTempInput("");

          if (typeof focusInput === "function") {
            focusInput();
          }
        }
        break;

      case "Tab": {
        e.preventDefault();

        const COMMANDS: string[] = [
          "help",
          "about",
          "projects",
          "whoami",
          "repo",
          "banner",
          "clear",
        ];
        const currentInput = inputValue;

        for (const ele of COMMANDS) {
          if (ele.startsWith(currentInput)) {
            setInputValue(ele);
            return;
          }
        }
        break;
      }

      case "ArrowUp":
        if (historyIndex.current === -1) {
          setTempInput(inputValue);
        }
        if (historyIndex.current < commandHistory.length - 1) {
          historyIndex.current++;
          setInputValue(
            commandHistory[commandHistory.length - 1 - historyIndex.current] ||
              "",
          );
        }
        break;

      case "ArrowDown":
        if (historyIndex.current > 0) {
          historyIndex.current--;
          setInputValue(
            commandHistory[commandHistory.length - 1 - historyIndex.current] ||
              "",
          );
        } else if (historyIndex.current === 0) {
          setInputValue(tempInput);
          historyIndex.current = -1;
        }
        break;

      default:
        break;
    }
  };

  return (
    <input
      ref={inputRef}
      type="text"
      value={inputValue}
      onChange={handleInputChange}
      onKeyDown={handleKeyDown}
      autoFocus
      className="border-none w-[800px] focus:outline-none text-gray-400 flex-1 bg-transparent"
    />
  );
};
