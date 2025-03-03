import { useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faGithub,
  faTwitter,
  faLinkedinIn,
} from "@fortawesome/free-brands-svg-icons";
import { SiLeetcode } from "react-icons/si";
import image from "/images/anime-pfp.jpg";

export const WhoAmI = () => {
  const [expanded, setExpanded] = useState(false);

  return (
    <div className="mt-4 ml-4 md:ml-8 flex flex-col gap-3 md:gap-4 items-start">
      <div className="flex flex-row items-start gap-3 md:gap-4">
        <img
          src={image}
          alt=""
          className="w-24 md:w-32 aspect-square rounded-md"
        />
        <div className="text-left">
          <p className="text-sm md:text-2xl font-bold text-primary">
            Saurav Singh Karmwar
          </p>
          <span className="text-sm md:text-base text-gray-300">
            Software Engineer
          </span>
          <div className="mt-3 flex gap-2 text-base md:text-lg text-customBlue">
            <a target="_blank" href="https://github.com/2SSK">
              <FontAwesomeIcon
                icon={faGithub}
                className="border border-primary/70 rounded-md p-1 md:p-2"
              />
            </a>
            <a target="_blank" href="https://x.com/_2SSK">
              <FontAwesomeIcon
                icon={faTwitter}
                className="border border-primary/70 rounded-md p-1 md:p-2"
              />
            </a>
            <a target="_blank" href="https://www.linkedin.com/in/2ssk/">
              <FontAwesomeIcon
                icon={faLinkedinIn}
                className="border border-primary/70 rounded-md p-1 md:p-2"
              />
            </a>
            <a target="_blank" href="https://leetcode.com/u/2SSK/">
              <SiLeetcode className="border border-primary/70 rounded-md p-1 md:p-2 size-7 md:size-9" />
            </a>
          </div>
        </div>
      </div>
      <div>
        <span className="text-customBlue font-semibold text-sm md:text-base">
          Who Am I.
        </span>
        <p className="text-wrap max-w-[95vw] md:max-w-[800px] text-gray-300 text-sm md:text-base">
          Hey, I’m Saurav, an aspiring full-stack web developer with a strong
          passion for DevOps and Linux. I specialize in building efficient,
          user-centric web applications and have a deep understanding of system
          administration and automation. My expertise lies in working with
          technologies like React, Next.js, and Node.js for both frontend and
          backend development, along with containerization and CI/CD pipelines
          to streamline deployments.
          <br />
          {!expanded && (
            <span
              className="text-blue-400 cursor-pointer"
              onClick={() => setExpanded(true)}
            >
              Show more...
            </span>
          )}
          {expanded && (
            <>
              <br />
              As a Linux power user, I enjoy customizing and optimizing my
              system to enhance productivity and workflow efficiency. My
              interest in DevOps has led me to explore automation tools like
              Docker, Kubernetes, and CI/CD, enabling me to create scalable,
              maintainable applications with smooth deployment processes.
              <br />
              <span
                className="text-red-400 cursor-pointer"
                onClick={() => setExpanded(false)}
              >
                Show less...
              </span>
            </>
          )}
        </p>
      </div>
    </div>
  );
};
