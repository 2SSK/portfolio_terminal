import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faGithub,
  faLinkedin,
  faTwitter,
} from "@fortawesome/free-brands-svg-icons";

import { IconProp } from "@fortawesome/fontawesome-svg-core"; // Import the IconProp type

import profile from "/images/pfp.jpeg";

const WhoAmI = () => {
  return (
    <div className="w-full h-full flex flex-col items-center justify-center p-5">
      <div className="w-[200px] h-[200px] rounded-full shadow-lg shadow-textColor overflow-hidden">
        <img
          src={profile}
          alt="Profile"
          className="w-full h-full object-cover"
        />
      </div>
      <div className="w-full px-4 mt-8 text-center">
        <span className="text-textColor text-2xl font-extrabold">
          WHO AM I?
        </span>
        <p className="text-lg mt-4 bg-customBlue bg-opacity-40 p-5 rounded-lg max-w-3xl w-full mx-auto break-words">
          Hey, I’m Saurav, a{" "}
          <b className="text-textColor">developer in making</b> with a passion
          for web development and all things Linux.
          <b className="text-textColor"> I use Arch Linux BTW</b> (with i3 and
          Neovim as my daily tools) and love customizing my system to make it
          uniquely mine. I’m working towards becoming a full-stack web
          developer, focusing on building functional, user-friendly web
          applications. Alongside that, I’m diving into DevOps, learning how to
          automate workflows and streamline deployments with tools like Docker
          and CI/CD. In the future, I want to combine my web development skills
          with DevOps practices to create scalable, efficient apps. When I’m not
          coding or tweaking my Linux setup, I’m usually exploring new tech,
          contributing to open source, or gaming.
        </p>

        <div className="mt-6 flex space-x-4 items-center justify-center gap-5">
          <SocialLink
            icon={faGithub}
            link="https://github.com/2SSK"
            text="2SSK"
          />
          <SocialLink
            icon={faLinkedin}
            link="https://linkedin.com/in/2ssk"
            text="2ssk"
          />
          <SocialLink
            icon={faTwitter}
            link="https://twitter.com/_2ssk"
            text="_2ssk"
          />
        </div>
      </div>
    </div>
  );
};

function SocialLink({
  icon,
  link,
  text,
}: {
  icon: IconProp; // Use IconProp for FontAwesome icons
  link: string;
  text: string;
}) {
  return (
    <a
      href={link}
      target="_blank"
      rel="noopener noreferrer"
      className="flex items-center space-x-2 text-customBlue hover:underline border-2 rounded-2xl px-3 py-1 border-customBlue hover:scale-110 transition-transform"
    >
      <FontAwesomeIcon icon={icon} size="lg" />
      <span>{text}</span>
    </a>
  );
}

export default WhoAmI;
