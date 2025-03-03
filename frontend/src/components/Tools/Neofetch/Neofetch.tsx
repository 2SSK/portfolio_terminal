import InfoRow from "./InfoRow";
import UptimeComponent from "./UptimeComponent";

interface NeofetchProps {
  PFP: string;
}

const Neofetch: React.FC<NeofetchProps> = ({ PFP }) => {
  return (
    <div className="w-full max-w-2xl  px-4 py-4  bg-opacity-60 rounded-lg flex flex-col md:flex-row justify-center md:items-start gap-10 md:gap-[80px] bg-[#1a1b26]">
      <div>
        <img
          src={PFP}
          alt="ASCII Profile"
          className="w-full md:w-[80%] md:h-full rounded-md"
        />
      </div>
      <div className="w-full whitespace-pre leading-relaxed flex flex-col gap-1 md:gap-1">
        <InfoRow label="User" value="SSK" />
        <InfoRow label="Host" value="ArchBTW" />
        <InfoRow label="Uptime" value={<UptimeComponent />} />
        <InfoRow label="Shell" value="zsh" />
        <InfoRow label="Editor" value="Neovim" />
        <InfoRow label="OS" value="Arch Linux" />
        <InfoRow label="Hobby" value="Tinkering with Linux" />
        <InfoRow label="Quote" value='"I use arch BTW 🐧"' />
      </div>
    </div>
  );
};

export default Neofetch;
