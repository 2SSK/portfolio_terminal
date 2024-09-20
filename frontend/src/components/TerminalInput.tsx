import { useState } from "react";
import { useSetRecoilState } from "recoil";
import { inputState } from "../store/atom/atom";

export default function TerminalInput() {
  return (
    <div className="w-full flex flex-col items-start justify-center text-lg">
      <p className="text-textColor font-bold mb-1">
        ssk<span className="text-foreground">@</span>
        archBTW ~
      </p>
      <div className="flex mb-6">
        <span className="text-primary text-lg font-bold mr-2">&gt;</span>
        <TerminalInputBox />
      </div>
    </div>
  );
}

const TerminalInputBox = () => {
  const setInputData = useSetRecoilState(inputState);
  const [inputValue, setInputValue] = useState("");

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && inputValue.trim() !== "") {
      setInputData(inputValue.trim());
      setInputValue(""); // Clear input after submission
    }
  };

  return (
    <input
      type="text"
      value={inputValue}
      onChange={handleInputChange}
      onKeyDown={handleKeyDown}
      autoFocus
      className="border-none w-[800px] focus:outline-none text-gray-400 flex-1 bg-transparent"
    />
  );
};
