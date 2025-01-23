"use client";

import Link from "next/link";
import { cn } from "@/lib/utils";
import React from "react";
import { usePathname } from "next/navigation";

function NavBarLink({
  href,
  currentPath,
  children,
}: {
  href: string;
  currentPath: string;
  children: React.ReactNode;
}) {
  return (
    <Link
      href={href}
      className={cn(
        "font-medium, inline-flex items-center px-1 pt-1 text-sm",
        currentPath === href
          ? "text-gray-900 border-b-2 border-primary-600"
          : "text-mono-1200 hover:border-mono-1000 hover:text-mono-1400",
      )}
    >
      {children}
    </Link>
  );
}
export function NavBarLinks() {
  const pathname = usePathname();
  return (
    <>
      <NavBarLink href="/me" currentPath={pathname}>
        Profile
      </NavBarLink>
      <NavBarLink href="/login" currentPath={pathname}>
        Login
      </NavBarLink>
      <NavBarLink href="/avatar-upload" currentPath={pathname}>
        Avatar Upload
      </NavBarLink>
    </>
  );
}
