import { lazy, Suspense } from "react";

const TerminalInput = lazy(() => import("./TerminalInput"));
const TerminalOutput = lazy(() => import("./TerminalOutput"));

export default function TerminalBox() {
  return (
    <div className="w-[80%] h-[90%] p-7 border-4 border-customBlue rounded-lg bg-transparent bg-opacity-30 backdrop-blur-sm shadow-lg shadow-disabled font-jetbrains overflow-y-auto">
      <Suspense fallback={<div>Loading...</div>}>
        <TerminalOutput />
        <TerminalInput />
      </Suspense>
    </div>
  );
}
