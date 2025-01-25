import { getCurrentUser } from "@/api/client";
import { LayoutNav } from "./components/layout-nav";
import { redirect } from "next/navigation";

export default async function Homepage({ children }: { children: React.ReactNode }) {
  const response = await getCurrentUser();
  const user = response?.user;
  if (!user) {
    return redirect("/");
  }
  const avatar = user.avatar;

  return (
    <>
      <div>
        <LayoutNav user={user} avatar={avatar}>
          {children}
        </LayoutNav>
      </div>
    </>
  );
}
