const helpObj = {
  commands: [
    ["about", "About me."],
    ["projects", "View my projects."],
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
            <div key={index} className="flex">
              <span className="w-28 text-primary font-bold mb-[6px]">
                <span className="p-1 rounded-md bg-customBlue bg-opacity-10">
                  {command}
                </span>
              </span>
              <span className="">{description}</span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Help;
