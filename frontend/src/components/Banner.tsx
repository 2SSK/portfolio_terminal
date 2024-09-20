import React, { useEffect, useState } from "react";

const Banner: React.FC = () => {
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
    <div>
      <Welcome asciiArt={asciiArt} />
    </div>
  );
};

interface WelcomeProps {
  asciiArt: string;
}

const Welcome: React.FC<WelcomeProps> = ({ asciiArt }) => {
  return (
    <div className="text-textColor text-sm leading-none">
      <pre style={{ fontFamily: "monospace", whiteSpace: "pre-wrap" }}>
        {asciiArt}
      </pre>
    </div>
  );
};

export default Banner;
