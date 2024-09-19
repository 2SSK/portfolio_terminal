import { useState, useEffect } from "react";
import { format } from "date-fns";

export default function Navbar() {
  const [now, setNow] = useState(new Date());

  useEffect(() => {
    const interval = setInterval(() => {
      setNow(new Date());
    }, 1000);

    return () => clearInterval(interval);
  });

  const formatedDate: string = format(now, "EEE dd-MMM yyy hh:mm a");

  return (
    <div className="w-full h-8 bg-background bg-opacity-100 backdrop-blur-sm shadow-md shadow-disabled flex items-center justify-between pr-10 text-foreground text-lg font-bold">
      <div className="flex items-center gap-2">
        <p className="bg-background-alt py-1 px-3">1</p>
        <p className="text-primary">
          <span className="text-textColor">welcome</span>@portfolio...
        </p>
      </div>

      <p className="text-textColor">{formatedDate}</p>

      <div className="">
        <ul className="flex items-center gap-8 text-primary cursor-pointer">
          <li>about</li>
          <li>project</li>
          <li>social</li>
          <li>resume</li>
        </ul>
      </div>
    </div>
  );
}
