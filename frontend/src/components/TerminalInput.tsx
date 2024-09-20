import { useState, lazy, Suspense } from "react";
import { useSetRecoilState, useRecoilValue } from "recoil";
import { inputState } from "../store/atom/atom";

const TerminalOutput = lazy(() => import("./TerminalOutput"));

export default function TerminalInput() {
  const inputData = useRecoilValue(inputState);

  return (
    <>
      <div className="w-full flex flex-col items-start justify-center text-lg">
        <span className="text-textColor font-bold mb-1">ssk@archBTW ~</span>
        <div className="flex mb-6">
          <span className="text-primary text-lg font-bold mr-2">&gt;</span>
          <TerminalInputBox />
        </div>
      </div>
      <Suspense fallback={<div>Loading...</div>}>
        {inputData && <TerminalOutput />}
      </Suspense>
    </>
  );
}

function TerminalInputBox() {
  const setInputData = useSetRecoilState(inputState);
  const [inputValue, setInputValue] = useState("whoami");

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && inputValue.trim() !== "") {
      setInputData(inputValue.trim());
      setInputValue("");
    }
  };

  return (
    <input
      type="text"
      value={inputValue}
      onChange={handleInputChange}
      onKeyDown={handleKeyDown}
      autoFocus
      className="border-none w-full focus:outline-none text-gray-400 flex-1 bg-transparent caret-transparent"
    />
  );
}
