import LogoutButton from "./LogoutBtn";

interface UserDropdownProps {
  user: { id: number; email: string };
}

export default function UserDropdown({ user }: UserDropdownProps) {
  return (
    <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 text-gray-700 z-10">
      <div className="px-4 py-2 text-sm border-b">
        <span className="font-semibold text-base">Email: </span>
        {user.email}
      </div>
      <div className="w-full p-1">
        <LogoutButton />
      </div>
    </div>
  );
}
