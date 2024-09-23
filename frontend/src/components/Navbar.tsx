import { useState, useEffect } from "react";
import { format } from "date-fns";
import { useSetRecoilState } from "recoil";
import { inputState } from "../store/atom/atom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars } from "@fortawesome/free-solid-svg-icons/faBars";

export default function Navbar() {
  const [showSubNavbar, setShowSubNavbar] = useState(false);

  const toggleSubNavbar = () => setShowSubNavbar((prev) => !prev);

  return (
    <div className="flex flex-col relative">
      <div className="w-full h-12 bg-transparent bg-opacity-100 backdrop-blur-sm shadow-md shadow-secondary flex items-center justify-between pr-10 text-foreground text-sm md:text-lg font-bold">
        <div className="flex items-center gap-2">
          <p className="bg-background-alt py-2 px-3">1</p>
          <p className="text-primary w-[200px]">
            <span className="text-textColor">ssk</span>
            <span className="text-foreground">@</span>portfolio...
          </p>
        </div>

        <DateTime />

        <ul className="hidden lg:flex gap-8 text-primary cursor-pointer">
          <ListLink text="about" />
          <ListLink text="project" />
          <ListLink text="social" />
          <ListLink text="resume" />
          <ListLink text="help" />
        </ul>

        <FontAwesomeIcon
          icon={faBars}
          onClick={toggleSubNavbar}
          className="block lg:hidden text-lg font-extrabold text-primary cursor-pointer"
        />
      </div>

      <SubNavbar show={showSubNavbar} />
    </div>
  );
}

// Sub Navbar Component (For smaller screens)
function SubNavbar({ show }: { show: boolean }) {
  return (
    <div
      className={`absolute top-12 left-0 w-full bg-background-alt bg-opacity-60 p-4 shadow-md lg:hidden transition-all duration-300 ease-in-out ${
        show ? "max-h-40 opacity-100" : "max-h-0 opacity-0"
      } overflow-hidden`}
      style={{ zIndex: 10 }} // Ensure it overlays other content
    >
      <ul className="flex justify-around text-primary cursor-pointer">
        <ListLink text="about" />
        <ListLink text="project" />
        <ListLink text="social" />
        <ListLink text="resume" />
        <ListLink text="help" />
      </ul>
    </div>
  );
}

// Remaining components (ListLink, DateTime) remain unchanged

// Component to display list items
type ListLinkProps = { text: string };

function ListLink({ text }: ListLinkProps) {
  const setInputState = useSetRecoilState(inputState);

  const handleOnClick = (e: React.MouseEvent<HTMLLIElement>) => {
    const uniqueId = Date.now();
    setInputState({ command: e.currentTarget.innerText, id: uniqueId });
  };

  return (
    <li
      onClick={handleOnClick}
      className="outline-none hover:text-highlight transition"
    >
      {text}
    </li>
  );
}

// Component to display current date and time
function DateTime() {
  const [now, setNow] = useState(new Date());

  useEffect(() => {
    const interval = setInterval(() => {
      setNow(new Date());
    }, 1000 * 60);

    return () => clearInterval(interval);
  }, []);

  const formattedDate: string = format(now, "EEE dd-MMM yyyy hh:mm a");

  return <p className="text-textColor">{formattedDate}</p>;
}
