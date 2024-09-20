import { useState, useEffect } from "react";
import { format } from "date-fns";

import { useSetRecoilState, useRecoilValue } from "recoil";
import { inputState } from "../store/atom/atom";

export default function Navbar() {
  const navText = useRecoilValue(inputState);

  return (
    <div className="w-full h-8 bg-transparent bg-opacity-100 backdrop-blur-sm shadow-md shadow-disabled flex items-center justify-between pr-10 text-foreground text-lg font-bold">
      <div className="flex items-center gap-2">
        <p className="bg-background-alt py-1 px-3">1</p>
        <p className="text-primary">
          <span className="text-textColor">{navText}</span>@portfolio...
        </p>
      </div>

      <DateTime />

      <div className="">
        <ul className="flex items-center gap-8 text-primary cursor-pointer">
          <ListLink text="about" />
          <ListLink text="project" />
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
    setInputState(e.currentTarget.innerText);
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
