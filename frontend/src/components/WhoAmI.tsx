const whoamiObj = {
  message: [
    [
      "In the kaleidoscope of existence,",
      "I am but a reflection questioning the enigma - ",
    ],
    [
      "Amidst cosmic whispers,",
      "I navigate the maze of self-discovery,",
      "echoing the eternal refrain - ",
    ],
    [
      "In the symphony of life,",
      "I am a note inquiring its own melody,",
      "harmonizing with the universal query - ",
    ],
    [
      "As stardust contemplating its journey,",
      "I ponder the cosmic query,",
      "silently asking - ",
    ],
    [
      "In the tapestry of reality,",
      "I am the thread of self-inquiry,",
      "weaving through the eternal question - ",
    ],
  ],
};

const WhoAmI = () => {
  const getRandomQuote = () => {
    const randomIndex = Math.floor(Math.random() * whoamiObj.message.length);
    return whoamiObj.message[randomIndex];
  };

  const selectedQuote = getRandomQuote();

  return (
    <div>
      <div className="p-4">
        <div className="mb-4">
          {selectedQuote.map((line, index) => {
            if (index === selectedQuote.length - 1) {
              return (
                <p key={index}>
                  {line}
                  <span className="text-primary"> Who Am I?</span>
                </p>
              );
            }
            return <p key={index}>{line}</p>;
          })}
        </div>
      </div>
    </div>
  );
};

export default WhoAmI;
