import { User } from "@/api/gen/user/v1/user_pb";
import { env } from "@/env";
import Image from "next/image";
import Link from "next/link";

export function NavBarRight({ user }: { user: User | null }) {
  if (user?.avatar) {
    return (
      <Link href="/me">
        <div className="h-12 w-12 overflow-hidden rounded-full bg-mono-600 ring-1 ring-primary-600">
          <Image
            alt={"profile picture"}
            src={env.SERVER_URL + user.avatar.url}
            width={user.avatar.width}
            height={user.avatar.height}
          />
        </div>
      </Link>
    );
  } else {
    return (
      <div className="h-12 w-12 rounded-full bg-mono-600 ring-1 ring-primary-600" />
    );
  }
}
