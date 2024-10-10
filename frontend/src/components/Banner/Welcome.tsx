import React, { useState, useEffect } from "react";

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

export default Welcome;
