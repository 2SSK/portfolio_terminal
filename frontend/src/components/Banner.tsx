import React, { useEffect, useState } from "react";

const repoLink = "https://github.com/2SSK/portfolio_terminal";

const Banner: React.FC = () => {
  return (
    <div className="text-sm sm:text-md">
      <Welcome />
      <BannerCommand command="help" /> for a list of all available commands
      <br />
      <BannerCommand command="repo" /> to view the GitHub repository or click{" "}
      <a href={repoLink} target="_blank" className="underline text-customBlue">
        here
      </a>
    </div>
  );
};

interface BannerCommandProps {
  command: string;
}

const BannerCommand = ({ command }: BannerCommandProps) => {
  return (
    <span className="text-primary rounded-md bg-customBlue bg-opacity-15">
      '{command}'
    </span>
  );
};

const Welcome: React.FC = () => {
  const [asciiArt, setAsciiArt] = useState<string>("");

  useEffect(() => {
    const fetchAsciiArt = async () => {
      try {
        const response = await fetch("/images/ascii-text-art.txt");
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const text = await response.text();
        setAsciiArt(text);
      } catch (error) {
        console.error("Failed to fetch ASCII art:", error);
      }
    };

    fetchAsciiArt();
  }, []);

  return (
    <div className="text-textColor leading-none w-fit mb-4 hidden sm:block">
      <pre style={{ fontFamily: "monospace", whiteSpace: "pre-wrap" }}>
        {asciiArt}
      </pre>
    </div>
  );
};

export default Banner;
