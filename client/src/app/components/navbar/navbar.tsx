import Link from "next/link";
import { Cookie } from "@/icons";
import React from "react";
import { NavBarLinks } from "./navbar-links";
import { NavBarRight } from "./navbar-right";
import { cookies } from "next/headers";
import { User } from "@/api/gen/user/v1/user_pb";
import { userByToken } from "@/api/client";

export async function Navbar() {
  let user: User | null = null;
  const cookieStore = await cookies();
  const sessionToken = cookieStore.get("x-session-token");
  if (sessionToken?.value) {
    const response = await userByToken(sessionToken.value);
    if (response.user) {
      user = response.user ?? null;
    }
  }

  return (
    <nav className="bg-white shadow-sm">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 justify-between">
          <div className="flex">
            <div className="flex shrink-0 items-center">
              <Link href="/">
                <Cookie height={36} width={36} />
              </Link>
            </div>
            <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
              <NavBarLinks />
            </div>
          </div>
          <div className="hidden sm:ml-6 sm:flex sm:items-center">
            <NavBarRight user={user} />
          </div>
        </div>
      </div>
    </nav>
  );
}
