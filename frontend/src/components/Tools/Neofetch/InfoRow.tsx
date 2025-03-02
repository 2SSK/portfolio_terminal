interface InfoRowProps {
  label: string;
  value: React.ReactNode;
}

const InfoRow = ({ label, value }: InfoRowProps) => {
  return (
    <div className="flex items-center">
      <span className="w-16 sm:w-20 font-bold text-[#7aa2f7] mr-2">
        {label}:
      </span>
      <span className="text-[#c0caf5] text-wrap">{value}</span>
    </div>
  );
};

export default InfoRow;
