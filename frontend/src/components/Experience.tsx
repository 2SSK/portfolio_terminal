const workExperience = [
  {
    company: "EOXS",
    position: "Software Developer Intern",
    startDate: "03/03/2025",
    endDate: "Ongoing",
    description:
      "Responsible for writing clean, efficient code, designing software architecture, debugging, and troubleshooting. Collaborates with teams to develop secure and scalable applications, ensuring best practices in testing, documentation, and deployment. Experienced in CI/CD, secure coding, and data protection.",
  },
];

const Experience = () => {
  return (
    <div className="mt-4 ml-4">
      {workExperience.map((job, index) => (
        <div
          key={index}
          className="max-w-lg border border-customBlue rounded-lg p-4 mb-4"
        >
          <div className="text-xl font-semibold">
            <span className="text-primary">{job.position}</span> @{" "}
            <span className="text-customBlue font-bold">{job.company}</span>
          </div>
          <div className="text-textColor text-sm mt-1">
            <span>{job.startDate}</span> - <span>{job.endDate}</span>
          </div>
          <p className="mt-2 text-base text-gray-300 ">{job.description}</p>
        </div>
      ))}
    </div>
  );
};

export default Experience;
