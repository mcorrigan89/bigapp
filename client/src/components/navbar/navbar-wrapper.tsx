import { getCurrentUser } from "@/api/client";
import { NavBarContent } from "./navbar-content";

const links = [
  { href: "/", label: "Home" },
  { href: "/about", label: "About" },
  { href: "/services", label: "Services" },
  { href: "/contact", label: "Contact" },
];

export async function NavBar() {
  const currentUserResponse = await getCurrentUser();
  const currentUser = currentUserResponse?.user ?? null;
  return <NavBarContent links={links} currentUser={currentUser} />;
}
