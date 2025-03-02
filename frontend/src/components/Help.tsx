const helpObj = {
  commands: [
    ["about", "About me."],
    ["project", "View my projects."],
    ["whoami", "A brief introduction."],
    ["repo", "View the GitHub repository."],
    ["resume", "Download my resume."],
    ["banner", "Display the banner."],
    ["social", "View my social profiles."],
    ["clear", "Clear the terminal."],
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
              <span className=" w-20 text-center text-primary font-sans font-bold mb-[6px] p-1 rounded-md bg-customBlue bg-opacity-10">
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
