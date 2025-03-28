"use client";

import { useState } from "react";
import UserDropdown from "./UserDropdown";

interface UserDropdownTriggerProps {
  user: { id: number; email: string };
  initials: string;
}

export default function UserDropdownTrigger({
  user,
  initials,
}: UserDropdownTriggerProps) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="w-8 aspect-square border border-white rounded-full flex items-center justify-center text-white hover:bg-gray-700 focus:outline-none"
      >
        {initials}
      </button>
      {isOpen && <UserDropdown user={user} />}
    </>
  );
}
