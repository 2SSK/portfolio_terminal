import { useState, useEffect } from "react";
import { format } from "date-fns";

import { useSetRecoilState } from "recoil";
import { inputState } from "../store/atom/atom";

export default function Navbar() {
  return (
    <div className="w-full h-8 bg-transparent bg-opacity-100 backdrop-blur-sm shadow-md shadow-secondary flex items-center justify-between pr-10 text-foreground text-lg font-bold">
      <div className="flex items-center gap-2">
        <p className="bg-background-alt py-1 px-3">1</p>
        <p className="text-primary w-[200px]">
          <span className="text-textColor">ssk</span>
          <span className="text-foreground">@</span>portfolio...
        </p>
      </div>

      <DateTime />

      <div className="">
        <ul className="flex items-center gap-8 text-primary cursor-pointer">
          <ListLink text="about" />
          <ListLink text="project" />
          {/* <ListLink text="experience" /> */}
          <ListLink text="social" />
          <ListLink text="resume" />
          <ListLink text="help" />
        </ul>
      </div>
    </div>
  );
}

type ListLinkProps = { text: string };

function ListLink({ text }: ListLinkProps) {
  const setInputState = useSetRecoilState(inputState);

  const handleOnClick = (e: React.MouseEvent<HTMLLIElement>) => {
    const uniqueId = Date.now();
    setInputState({ command: e.currentTarget.innerText, id: uniqueId });
  };

  return (
    <li onClick={handleOnClick} className="outline-none">
      {text}
    </li>
  );
}

function DateTime() {
  const [now, setNow] = useState(new Date());

  useEffect(() => {
    const interval = setInterval(() => {
      setNow(new Date());
    }, 1000 * 60);

    return () => clearInterval(interval);
  });

  const formatedDate: string = format(now, "EEE dd-MMM yyy hh:mm a");

  return <p className="text-textColor">{formatedDate}</p>;
}
