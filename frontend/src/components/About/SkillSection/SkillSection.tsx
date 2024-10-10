import Badge from "./Badge";

interface SkillSectionProps {
  title: string;
  items: string[];
}

const SkillSection = ({ title, items }: SkillSectionProps) => {
  return (
    <div className="my-8">
      <span className="font-bold text-[#bb9af7] mb-4">{title}:</span>
      <div className="flex gap-2 flex-wrap">
        {items.map((item, index) => (
          <Badge key={index} text={item} />
        ))}
      </div>
    </div>
  );
};

export default SkillSection;
