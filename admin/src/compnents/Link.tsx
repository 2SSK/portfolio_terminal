"use client";
import Link from "next/link";

export const LinkTxt = ({ title, route }: { title: string; route: string }) => {
  return (
    <Link href={route}>
      <span className="text-blue-500 cursor-pointer hover:underline">
        {title}
      </span>
    </Link>
  );
};
