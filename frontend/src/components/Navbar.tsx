import { format } from "date-fns";

import logo from "/logo2.png";

export default function Navbar() {
  const now = new Date();
  const formatedDate: string = format(now, "EEE dd-MM yyy hh:mm a");

  return (
    <div className="w-full h-5 flex">
      <div>
        <img src={logo} alt="" className="w-8" />
        <p>welcome@portfolio...</p>
      </div>

      <p>{formatedDate}</p>

      <div>
        <ul>
          <li>help</li>
          <li>about</li>
          <li>social</li>
          <li>resume</li>
        </ul>
      </div>
    </div>
  );
}
