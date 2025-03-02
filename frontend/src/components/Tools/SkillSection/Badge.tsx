const Badge = ({ text }: { text: string }) => {
  return (
    <span
      className={`px-2 py-1 text-s rounded-md bg-primary bg-opacity-10 font-bold text-customBlue`}
    >
      {text}
    </span>
  );
};

export default Badge;
