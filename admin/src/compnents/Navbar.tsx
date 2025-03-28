import { checkAuth } from "@/actions/auth";
import UserDropdownTrigger from "./UserDropdownTrigger";
import NavLinks from "./NavLinks";

export default async function Navbar() {
  const user = await checkAuth();
  const initials = user.email.charAt(0).toUpperCase();

  return (
    <div className="w-full">
      <div className="flex items-center justify-between py-1 px-8 bg-gray-800 text-white">
        {/* Logo */}
        <div className="font-bold text-xl">Admin</div>
        {/* Links */}
        <NavLinks />
        {/* User Btn */}
        <div className="relative">
          <UserDropdownTrigger user={user} initials={initials} />
        </div>
      </div>
    </div>
  );
}
