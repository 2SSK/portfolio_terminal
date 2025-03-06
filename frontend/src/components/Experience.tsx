const workExperience = [
  {
    company: "DEMO",
    companyLink: "https://ssk-portfolio.vercel.app",
    position: "Software Developer",
    startDate: "Mar 2025",
    endDate: "Ongoing",
    description:
      "Responsible for writing clean, efficient code, designing software architecture, debugging, and troubleshooting. Collaborates with teams to develop secure and scalable applications, ensuring best practices in testing, documentation, and deployment. Experienced in CI/CD, secure coding, and data protection.",
  },
  //{
  //  company: "EOXS",
  //  companyLink: "https://www.eoxs.com",
  //  position: "Software Developer Intern",
  //  startDate: "03/03/2025",
  //  endDate: "Ongoing",
  //  description:
  //    "Responsible for writing clean, efficient code, designing software architecture, debugging, and troubleshooting. Collaborates with teams to develop secure and scalable applications, ensuring best practices in testing, documentation, and deployment. Experienced in CI/CD, secure coding, and data protection.",
  //},
];

const Experience = () => {
  return (
    <div className="mt-4 md:ml-4">
      {workExperience.map((job, index) => (
        <div
          key={index}
          className="max-w-lg border border-customBlue rounded-lg p-4 mb-4"
        >
          <div className="text-base md:text-xl font-semibold">
            <span className="text-primary">{job.position}</span> @{" "}
            <a href={job.companyLink} target="_blank">
              <span className="text-customBlue font-bold underline">
                {job.company}
              </span>
            </a>
          </div>
          <div className="text-textColor text-sm mt-1">
            <span>{job.startDate}</span> - <span>{job.endDate}</span>
          </div>
          <p className="mt-2 text-sm md:text-base text-gray-300 ">
            {job.description}
          </p>
        </div>
      ))}
    </div>
  );
};

export default Experience;
