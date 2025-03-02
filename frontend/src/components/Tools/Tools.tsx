import { useEffect } from "react";
import PFP from "/images/anime-pfp.jpg";

import Neofetch from "./Neofetch/Neofetch";
import SkillSection from "./SkillSection/SkillSection";

const Tools = () => {
  useEffect(() => {
    const img = new Image();
    img.src = PFP;
  }, []);

  const languages = [
    "C",
    "C++",
    "Golang",
    "Python",
    "Bash Script",
    "TypeScript",
    "JavaScript",
    "Markdown",
  ];
  const tools = [
    "Linux",
    "Neovim",
    "Tmux",
    "Vercel",
    "Render",
    "Netlify",
    "Postman",
    "Compass",
    "Figma",
    "Obsidian",
  ];
  const frameworks = ["Reactjs", "Nextjs", "Expressjs", "GoFiber"];
  const databases = ["MySQL", "mongoDB", "POSTGRESQL"];

  return (
    <div>
      <Neofetch PFP={PFP} />

      {/* Languages Section */}
      <SkillSection title="LANGUAGES" items={languages} />

      {/* Tools Section */}
      <SkillSection title="TOOLS" items={tools} />

      {/* Frameworks Section */}
      <SkillSection title="FRAMEWORKS" items={frameworks} />

      {/* Databases Section */}
      <SkillSection title="DATABASE" items={databases} />
    </div>
  );
};

export default Tools;
