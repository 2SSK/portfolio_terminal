import { useState, useEffect } from "react";
import ASCII_PFP from "/images/ascii_pfp.jpg";

const About = () => {
  return (
    <div className="w-full flex items-center gap-[100px]">
      <div>
        <img
          src={ASCII_PFP}
          alt="ASCII Profile"
          className="w-[300px] brightness-200"
        />
      </div>
      <div className="whitespace-pre leading-relaxed flex flex-col gap-2">
        <InfoRow label="User" value="ssk" />
        <InfoRow label="Host" value="archBTW" />
        <InfoRow label="Uptime" value={<UptimeComponent />} />{" "}
        <InfoRow label="Shell" value="zsh" />
        <InfoRow label="Editor" value="Neovim" />
        <InfoRow label="OS" value="Arch Linux" />
        <InfoRow label="Language" value="TypeScript, JavaScript" />
        <InfoRow label="Tools" value="React, Recoil, Express, Node.js" />
        <InfoRow label="Hobby" value="Tinkering with Linux & Custom setups" />
        <InfoRow label="Quote" value='"I use arch BTW"' />
      </div>
    </div>
  );
};

const InfoRow = ({
  label,
  value,
}: {
  label: string;
  value: React.ReactNode;
}) => {
  return (
    <div className="flex items-center gap-[4px] text-md">
      <span className="w-[100px] font-extrabold text-textColor">{label}:</span>
      <span>{value}</span>
    </div>
  );
};

function UptimeComponent() {
  const [uptime, setUptime] = useState("");

  useEffect(() => {
    const startDate = new Date("2003-06-15");
    const updateUptime = () => {
      const now = new Date();
      const differenceInTime = now.getTime() - startDate.getTime();
      const differenceInDays = differenceInTime / (1000 * 3600 * 24);
      const years = Math.floor(differenceInDays / 365);
      const remainingDays = Math.floor(differenceInDays % 365);

      setUptime(`${years} years, ${remainingDays} days`);
    };

    updateUptime();
    const intervalId = setInterval(updateUptime, 1000 * 60 * 60 * 24);

    return () => clearInterval(intervalId);
  }, []);

  return <>{uptime}</>;
}

export default About;
