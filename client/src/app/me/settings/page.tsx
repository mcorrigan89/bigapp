import { getCurrentUser } from "@/api/client";
import { redirect } from "next/navigation";
import { ProfileInfoForm } from "./forms/form";

const secondaryNavigation = [
  { name: "Account", href: "#", current: true },
  { name: "Notifications", href: "#", current: false },
  { name: "Billing", href: "#", current: false },
  { name: "Teams", href: "#", current: false },
  { name: "Integrations", href: "#", current: false },
];

export default async function SettingsPage() {
  const response = await getCurrentUser();
  const user = response?.user;
  if (!user) {
    return redirect("/");
  }
  const avatar = user.avatar;

  return (
    <main>
      <h1 className="sr-only">Account Settings</h1>

      <header className="border-b border-white/5">
        {/* Secondary navigation */}
        <nav className="flex overflow-x-auto py-4">
          <ul role="list" className="flex min-w-full flex-none gap-x-6 px-4 text-sm/6 font-semibold text-gray-400 sm:px-6 lg:px-8">
            {secondaryNavigation.map((item) => (
              <li key={item.name}>
                <a href={item.href} className={item.current ? "text-indigo-400" : ""}>
                  {item.name}
                </a>
              </li>
            ))}
          </ul>
        </nav>
      </header>

      {/* Settings forms */}
      <div className="divide-y divide-white/5">
        <div className="grid max-w-7xl grid-cols-1 gap-x-8 gap-y-10 px-4 py-16 sm:px-6 md:grid-cols-3 lg:px-8">
          <div>
            <h2 className="text-base/7 font-semibold text-white">Personal Information</h2>
            <p className="mt-1 text-sm/6 text-gray-400">Use a permanent address where you can receive mail.</p>
          </div>
          <ProfileInfoForm user={user} avatar={avatar} />
        </div>

        <div className="grid max-w-7xl grid-cols-1 gap-x-8 gap-y-10 px-4 py-16 sm:px-6 md:grid-cols-3 lg:px-8">
          <div>
            <h2 className="text-base/7 font-semibold text-white">Delete account</h2>
            <p className="mt-1 text-sm/6 text-gray-400">
              No longer want to use our service? You can delete your account here. This action is not reversible. All information related to this account will
              be deleted permanently.
            </p>
          </div>

          <form className="flex items-start md:col-span-2">
            <button type="submit" className="rounded-md bg-red-500 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-red-400">
              Yes, delete my account
            </button>
          </form>
        </div>
      </div>
    </main>
  );
}
