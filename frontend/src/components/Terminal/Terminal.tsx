import { lazy, Suspense, useRef, useEffect } from "react";
import { useSetRecoilState } from "recoil";
import { focusInputState } from "../../store/atom/atom";

const TerminalInput = lazy(() => import("./TerminalInput"));
const TerminalOutput = lazy(() => import("./TerminalOutput"));

export default function TerminalBox({ className }: { className?: string }) {
  const inputRef = useRef<HTMLInputElement>(null);
  const setFocusInput = useSetRecoilState(focusInputState);
  const outputRef = useRef<HTMLDivElement>(null); // Ref for auto-scrolling

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
      className={`w-full h-[85%] sm:h-full mb-5 p-4 border-4 border-customBlue rounded-lg bg-[#011423] bg-opacity-50 backdrop-blur-xl shadow-lg shadow-secondary overflow-y-auto overflow-x-hidden ${className}`}
      ref={outputRef}
    >
      <style>{`
        body {
          overflow: hidden; /* Disable scrolling on the webpage */
        }
        .terminal-box::-webkit-scrollbar {
          display: none; /* Hide scrollbars inside the terminal */
        }
      `}</style>
      <Suspense fallback={<div>Loading...</div>}>
        <TerminalOutput outputRef={outputRef} />
        <TerminalInput inputRef={inputRef} />
      </Suspense>
    </div>
  );
}
