import { userByToken } from "@/api/client";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

export default async function CurrentUserPage() {
  const cookieStore = await cookies();
  const sessionToken = cookieStore.get("x-session-token");
  if (!sessionToken) {
    return redirect("/");
  }

  const response = await userByToken(sessionToken.value);
  return (
    <div>
      <div>Hello world</div>
      <div>{JSON.stringify(response.user)}</div>
    </div>
  );
}
