import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { IconProp } from "@fortawesome/fontawesome-svg-core";
import {
  faGithub,
  faTwitter,
  faLinkedinIn,
} from "@fortawesome/free-brands-svg-icons";
import { faEnvelope } from "@fortawesome/free-solid-svg-icons";

const emailId = "sauravchp2@gmail.com";
const github = "https://github.com/2SSK";
const linkedin = "https://linkedin.com/in/2ssk";
const twitter = "https://x.com/_2ssk";

const Social = () => {
  return (
    <div className="mt-4 md:ml-8 flex flex-col ">
      <span className="md:hidden font-bold text-xl text-[#bb9af7] mt-4">
        Contact:
      </span>
      <div>
        <span className="text-base md:text-lg text-customBlue font-semibold">
          Reach out to me.
        </span>
        <p className="text-sm md:text-base text-gray-400">
          Feel free to reach out to me via email for any queries,
          <br />
          collaboration opportunities, or further details.
        </p>
      </div>
      <div className="mt-4 flex flex-wrap gap-2">
        <SocialBox link={twitter} icon={faTwitter} title="Twitter" />
        <SocialBox link={`mailto:${emailId}`} icon={faEnvelope} title="Email" />
        <SocialBox link={linkedin} icon={faLinkedinIn} title="Linkedin" />
        <SocialBox link={github} icon={faGithub} title="Github" />
      </div>
    </div>
  );
};

const SocialBox = ({
  link,
  icon,
  title,
}: {
  link: string;
  title: string;
  icon: IconProp;
}) => {
  return (
    <div className="border border-customBlue rounded-md px-2 py-1 text-primary hover:bg-customBlue/10">
      <a target="_blank" href={link} rel="noopener noreferrer">
        <span>
          <FontAwesomeIcon icon={icon} className="mr-3" />
          {title}
        </span>
      </a>
    </div>
  );
};

export default Social;
