interface LinkProps {
  href: string;
  title: string;
  description: string;
}

const Link = ({ href, title, description }: LinkProps) => {
  return (
    <div className="flex flex-col md:flex-row items-start">
      <a
        href={href}
        target="_blank"
        className="underline text-[#bb9af7] w-[150px] text-md"
      >
        {title}
      </a>
      <p className="text-[#c0caf5] text-sm">{description}</p>
    </div>
  );
};

export default Link;
