import { userByToken } from "@/api/client";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import Image from "next/image";
import { env } from "@/env";

export default async function CurrentUserPage() {
  const cookieStore = await cookies();
  const sessionToken = cookieStore.get("x-session-token");
  if (!sessionToken) {
    return redirect("/");
  }

  const response = await userByToken(sessionToken.value);
  const user = response.user;

  if (!user) {
    return redirect("/");
  }
  const avatarImage = user.avatar;

  return (
    <>
      <div>
        <div>Current User</div>
        <div>First Name</div>
        <div>{user.givenName}</div>
        <div>Last Name</div>
        <div>{user.familyName}</div>
        <div>Email</div>
        <div>{user.email}</div>
        <div>{avatarImage?.url}</div>
        {avatarImage ? <Image src={env.SERVER_URL + avatarImage.url} width={avatarImage.width} height={avatarImage.height} alt="avatar" /> : null}
      </div>
    </>
  );
}
