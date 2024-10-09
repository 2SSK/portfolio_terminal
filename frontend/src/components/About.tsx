import { useState, useEffect } from "react";
import PFP from "/images/ascii_pfp.jpg";

const About = () => {
  useEffect(() => {
    const img = new Image();
    img.src = PFP;
  }, []);

  const languages = [
    "C",
    "C++",
    "Bash",
    "Python",
    "TypeScript",
    "JavaScript",
    "Lua",
    "Markdown",
  ];
  const tools = [
    "Linux",
    "Neovim",
    "Emacs",
    "Vercel",
    "Render",
    "Netlify",
    "Postman",
    "Compass",
    "GIMP",
    "Canva",
    "Trello",
    "Notion",
    "Obsidian",
  ];
  const frameworks = [
    "React",
    "Next",
    "Node",
    "Express",
    "JWT",
    "ReactNative",
    "Vite",
  ];
  const databases = ["MariaDB", "MySQL", "mongoDB", "POSTGRES"];

  return (
    <div>
      <Neofetch />

      {/* Languages Section */}
      <SkillsSection title="LANGUAGES" items={languages} />

      {/* Tools Section */}
      <SkillsSection title="TOOLS" items={tools} />

      {/* Frameworks Section */}
      <SkillsSection title="FRAMEWORKS" items={frameworks} />

      {/* Databases Section */}
      <SkillsSection title="DATABASE" items={databases} />
    </div>
  );
};

// SkillsSection component to avoid repetitive divs
const SkillsSection = ({
  title,
  items,
}: {
  title: string;
  items: string[];
}) => {
  return (
    <div className="my-8">
      <span className="font-bold text-[#bb9af7] mb-4">{title}:</span>
      <div className="flex gap-2 flex-wrap">
        {items.map((item, index) => (
          <Badge key={index} text={item} />
        ))}
      </div>
    </div>
  );
};

// Neofetch component
function Neofetch() {
  return (
    <div className="w-[98%] sm:w-[850px] px-4 py-4 md:px-8 bg-opacity-60 rounded-lg flex flex-col sm:flex-row items-start sm:items-center gap-[50px] md:gap-[100px] bg-[#1a1b26]">
      <div>
        <img src={PFP} alt="ASCII Profile" className="w-[300px] rounded-md" />
      </div>
      <div className="whitespace-pre leading-relaxed flex flex-col gap-1 md:gap-2">
        <InfoRow label="User" value="ssk" />
        <InfoRow label="Host" value="archBTW" />
        <InfoRow label="Uptime" value={<UptimeComponent />} />{" "}
        <InfoRow label="Shell" value="zsh" />
        <InfoRow label="Editor" value="Neovim" />
        <InfoRow label="OS" value="Arch Linux" />
        <InfoRow label="Hobby" value="Tinkering with Linux" />
        <InfoRow label="Quote" value='"I use arch BTW 🐧"' />
      </div>
    </div>
  );
}

// UptimeComponent for calculating years and days
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

// InfoRow component to display label and value
const InfoRow = ({
  label,
  value,
}: {
  label: string;
  value: React.ReactNode;
}) => {
  return (
    <div className="flex items-center">
      <span className="w-16 sm:w-20 font-bold text-[#7aa2f7]">{label}:</span>
      <span className="text-[#c0caf5] text-wrap">{value}</span>
    </div>
  );
};

// Badge component to display skills
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
