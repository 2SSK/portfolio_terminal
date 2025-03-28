"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function NavLinks() {
  const pathname = usePathname(); // Client-side current path
  const links = [
    { href: "/bio", label: "Bio" },
    { href: "/resume", label: "Resume" },
    { href: "/socials", label: "Socials" },
    { href: "/tools", label: "Tools" },
    { href: "/projects", label: "Projects" },
    { href: "/experience", label: "Experience" },
  ];

  return (
    <ul className="flex space-x-4">
      {links.map((link) => {
        const isActive =
          pathname === link.href || (pathname === "/" && link.href === "/bio"); // /bio is landing page
        return (
          <li key={link.href}>
            <Link
              href={link.href}
              className={`${
                isActive
                  ? "bg-blue-600 text-white font-semibold"
                  : "text-white hover:text-blue-200 hover:bg-gray-700"
              } w-24 text-center px-1 py-0.5 rounded-md transition-colors block`}
            >
              {link.label}
            </Link>
          </li>
        );
      })}
    </ul>
  );
}
