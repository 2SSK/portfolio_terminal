import { useEffect } from "react";
import PFP from "/images/anime-pfp.jpg";

import Neofetch from "./Neofetch/Neofetch";
import SkillSection from "./SkillSection/SkillSection";

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

export default About;
