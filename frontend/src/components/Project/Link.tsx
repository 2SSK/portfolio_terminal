import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
import { faExternalLinkAlt } from "@fortawesome/free-solid-svg-icons";

interface LinkProps {
  href: string;
  github: string;
  title: string;
  description: string;
  image?: string;
}

const Link = ({ href, github, title, description, image }: LinkProps) => {
  return (
    <div className="max-w-[700px] border border-blue-300 rounded-lg p-4 flex flex-col justify-start items-start md:flex-row gap-4 hover:bg-blue-800/15 transition duration-150">
      {image && (
        <img
          src={image}
          alt={title}
          className="rounded-md w-full h-[180px] md:max-w-[200px] md:min-w-[200px] md:h-[130px] md:hover:scale-150 md:duration-300 md:transition-all md:ease-in-out object-cover"
        />
      )}
      <div className="flex flex-col justify-start">
        <div className="flex justify-between items-center md:text-lg font-semibold text-gray-100">
          <p className="text-primary">{title}</p>
          <div className="flex gap-3 text-customBlue">
            <a href={href} target="_blank" rel="noopener noreferrer">
              <FontAwesomeIcon icon={faExternalLinkAlt} />
            </a>
            <a href={github} target="_blank" rel="noopener noreferrer">
              <FontAwesomeIcon icon={faGithub} />
            </a>
          </div>
        </div>
        <p className="text-gray-100 text-sm mt-2">{description}</p>
      </div>
    </div>
  );
};

export default Link;
