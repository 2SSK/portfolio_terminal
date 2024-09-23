import { useState, useEffect } from "react";
import ASCII_PFP from "/images/ascii_pfp.jpg";

const About = () => {
  useEffect(() => {
    const img = new Image();
    img.src = ASCII_PFP;
  }, []);

  return (
    <div>
      <div className="w-[850px] px-4 py-4 md:px-8 bg-opacity-60 rounded-lg flex items-center gap-[50px] md:gap-[100px] bg-[#1a1b26]">
        <div>
          <img
            src={ASCII_PFP}
            alt="ASCII Profile"
            className="w-[200px] md:w-[300px] brightness-200"
          />
        </div>
        <div className="whitespace-pre leading-relaxed flex flex-col gap-1 md:gap-2">
          <InfoRow label="User" value="ssk" />
          <InfoRow label="Host" value="archBTW" />
          <InfoRow label="Uptime" value={<UptimeComponent />} />{" "}
          <InfoRow label="Shell" value="zsh" />
          <InfoRow label="Editor" value="Neovim" />
          <InfoRow label="OS" value="Arch Linux" />
          <InfoRow label="Hobby" value="Tinkering with Linux & Custom setups" />
          <InfoRow label="Quote" value='"I use arch BTW 🐧"' />
        </div>
      </div>

      <div className="my-8">
        <span className="font-bold text-[#bb9af7] mb-4">LANGUAGES:</span>
        <div className="flex gap-2 flex-wrap">
          <Badge text="C" />
          <Badge text="C++" />
          <Badge text="Bash" />
          <Badge text="Python" />
          <Badge text="TypeScript" />
          <Badge text="JavaScript" />
          <Badge text="Lua" />
          <Badge text="Markdown" />
        </div>
      </div>
      <div>
        <span className="font-bold text-[#bb9af7] mb-4">TOOLS:</span>
        <div className="flex gap-2 flex-wrap">
          <Badge text="Linux" />
          <Badge text="Neovim" />
          <Badge text="Vercel" />
          <Badge text="Render" />
          <Badge text="Netlify" />
          <Badge text="Postman" />
          <Badge text="Compass" />
          <Badge text="GIMP" />
          <Badge text="Canva" />
          <Badge text="Trello" />
          <Badge text="Notion" />
          <Badge text="Obsidian" />
        </div>
      </div>
      <div className="mt-8">
        <span className="font-bold text-[#bb9af7] mb-4">FRAMEWORKS:</span>
        <div className="flex gap-2 flex-wrap">
          <Badge text="React" />
          <Badge text="Next" />
          <Badge text="Node" />
          <Badge text="Express" />
          <Badge text="JWT" />
          <Badge text="ReactNative" />
          <Badge text="Vite" />
        </div>
      </div>
      <div className="mt-8">
        <span className="font-bold text-[#bb9af7] mb-4">DATABASE:</span>
        <div className="flex gap-2">
          <Badge text="MariaDB" />
          <Badge text="MySQL" />
          <Badge text="mongoDB" />
          <Badge text="POSTGRES" />
        </div>
      </div>
    </div>
  );
};

// Calculating age in years and days
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

// Component to display label and value
const InfoRow = ({
  label,
  value,
}: {
  label: string;
  value: React.ReactNode;
}) => {
  return (
    <div className="flex items-center flex-wrap gap-[2px] md:gap-[4px] text-sm md:text-md">
      <span className="w-[55px] md:w-[100px] font-extrabold text-[#7aa2f7]">
        {label}:
      </span>
      <span className="text-[#c0caf5]">{value}</span>
    </div>
  );
};

// Component to display badge with random color
const Badge = ({ text }: { text: string }) => {
  return (
    <span
      className={`px-2 py-1 text-s rounded-md bg-primary bg-opacity-10 text-customBlue`}
    >
      {text}
    </span>
  );
};

export default About;
