import { getCurrentUser } from "@/api/client";
import { LayoutNavigation } from "./components/layout-nav";
import { redirect } from "next/navigation";

interface Props {
  children: React.ReactNode;
}
export default async function Layout({ children }: Props) {
  const response = await getCurrentUser();
  const user = response?.user;
  if (!user) {
    return redirect("/");
  }
  const avatar = user.avatar;

  return (
    <>
      <div>
        <LayoutNavigation user={user} avatar={avatar} />
        <main className="py-10 lg:pl-72">
          <div className="px-4 sm:px-6 lg:px-8">{children}</div>
        </main>
      </div>
    </>
  );
}
