"use client";

import { useRouter } from "next/navigation";
import { logout } from "@/actions/auth";

export default function LogoutButton() {
  const router = useRouter();

  const handleLogout = async () => {
    await logout();
    router.refresh();
  };

  return (
    <button
      onClick={handleLogout}
      className="w-full text-left px-4 py-2 rounded-sm text-sm font-semibold bg-blue-500 text-white hover:bg-blue-600"
    >
      Logout
    </button>
  );
}
