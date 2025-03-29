import { BioResponse } from "@/actions/bio_action";
import Image from "next/image";

type BioBoxProps = {
  bio: BioResponse | null;
};

export function BioBox({ bio }: BioBoxProps) {
  return (
    <div className="flex-1 rounded-lg bg-white p-6 shadow-lg">
      {bio ? (
        <div className="flex flex-col gap-6">
          <div className="flex items-center gap-4">
            {bio.image &&
            typeof bio.image === "string" &&
            bio.image.startsWith("http") ? (
              <Image
                src={bio.image}
                alt={`${bio.name || "User"}'s profile`}
                width={100}
                height={100}
                className="rounded-md shadow-md object-cover"
                priority={false}
              />
            ) : (
              <div className="flex h-[100px] w-[100px] items-center justify-center rounded-full bg-gray-200">
                <span className="text-gray-500">No Image</span>
              </div>
            )}
            <div>
              <p className="text-xl font-semibold text-gray-800">
                {bio.name ?? "No name provided"}
              </p>
              <h2 className="text-md font-medium text-gray-600">{bio.title}</h2>
            </div>
          </div>
          <p className="max-w-sm text-gray-500">
            {bio.description ?? "No description provided"}
          </p>
        </div>
      ) : (
        <p className="text-center text-gray-500">
          No bio found. Add one below!
        </p>
      )}
    </div>
  );
}
