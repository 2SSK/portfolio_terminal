import { lazy, Suspense, useRef, useEffect } from "react";
import { useSetRecoilState } from "recoil";
import { focusInputState } from "../store/atom/atom";

const TerminalInput = lazy(() => import("./TerminalInput"));
const TerminalOutput = lazy(() => import("./TerminalOutput"));

export default function TerminalBox() {
  const inputRef = useRef<HTMLInputElement>(null);
  const setFocusInput = useSetRecoilState(focusInputState);

  useEffect(() => {
    setFocusInput(() => () => {
      inputRef.current?.focus();
    });
  }, [setFocusInput]);

  const handleFocusInput = () => {
    inputRef.current?.focus();
  };

  return (
    <div
      onClick={handleFocusInput}
      className="w-[95%] md:w-[80%] h-[85%] md:h-[90%] mb-5 p-7 border-4 border-customBlue rounded-lg bg-transparent bg-opacity-30 backdrop-blur-md shadow-lg shadow-secondary overflow-auto"
    >
      <style>{`
        .terminal-box::-webkit-scrollbar {
          display: none;
        }
      `}</style>
      <Suspense fallback={<div>Loading...</div>}>
        <TerminalOutput />
        <TerminalInput inputRef={inputRef} />
      </Suspense>
    </div>
  );
}
