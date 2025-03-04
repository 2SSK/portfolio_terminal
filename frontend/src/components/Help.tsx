const helpObj = {
  commands: [
    ["whoami", "A brief introduction."],
    ["tools", "tools/tech I use."],
    ["project", "View my projects."],
    ["repo", "View the GitHub repository."],
    ["resume", "Download my resume."],
    ["banner", "Display the banner."],
    ["social", "View my social profiles."],
    ["clear", "Clear the terminal."],
    ["experience", "my work experiences"],
  ],
};

const Help = () => {
  return (
    <div className="p-1 md:p-4 text-gray-300">
      <div className="mb-1 md:mb-4">
        {helpObj.commands.map((ele, index) => {
          const command = ele[0];
          const description = ele[1];
          return (
            <div key={index} className="flex gap-2 md:gap-4 items-center">
              <span className="border-none min-w-[80px] px-2 text-sm md:text-base text-center text-primary font-sans font-bold mb-[6px] p-0.5 md:p-1 rounded-md bg-customBlue bg-opacity-10">
                {command}
              </span>
              <span className="text-sm md:text-lg text-center">
                {description}
              </span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Help;
