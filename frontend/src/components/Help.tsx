const helpObj = {
  commands: [
    ["whoami", "A brief introduction."],
    ["tools", "tools and technologies I use."],
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
    <div className="p-4 text-gray-300">
      <div className="mb-4">
        {helpObj.commands.map((ele, index) => {
          const command = ele[0];
          const description = ele[1];
          return (
            <div key={index} className="flex gap-4">
              <span className="border-none w-24 text-center text-primary font-sans font-bold mb-[6px] p-1 rounded-md bg-customBlue bg-opacity-10">
                {command}
              </span>
              <span className="text-lg">{description}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Help;
