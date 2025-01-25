import { userByHandle } from "@/api/client";
import { env } from "@/env";
import { EnvelopeIcon, PhoneIcon } from "@heroicons/react/20/solid";
import Image from "next/image";
import { redirect } from "next/navigation";

type Params = Promise<{ handle: string }>;
interface Props {
  params: Params;
}

export default async function UserProfile({ params }: Props) {
  const { handle } = await params;
  const response = await userByHandle(handle);
  const user = response.user;
  if (!user) {
    return redirect("/");
  }
  const avatar = user.avatar;

  return (
    <div>
      <div className="flex justify-center">
        {avatar ? (
          <Image
            alt=""
            src={env.NEXT_PUBLIC_SERVER_URL + avatar.url}
            width={avatar.width}
            height={avatar.height}
            className="m-2 h-48 w-full rounded-lg object-cover lg:mx-8 lg:h-80 xl:mx-36 xl:h-96 xl:rounded-xl"
          />
        ) : null}
      </div>
      <div className="mx-auto max-w-5xl px-4 sm:px-6 lg:px-8">
        <div className="-mt-12 sm:-mt-16 sm:flex sm:items-end sm:space-x-5">
          <div className="flex">
            {avatar ? (
              <Image
                alt=""
                src={env.NEXT_PUBLIC_SERVER_URL + avatar.url}
                width={avatar.width}
                height={avatar.height}
                className="size-24 rounded-full ring-4 ring-white sm:size-32"
              />
            ) : null}
          </div>
          <div className="mt-6 sm:flex sm:min-w-0 sm:flex-1 sm:items-center sm:justify-end sm:space-x-6 sm:pb-1">
            <div className="mt-6 min-w-0 flex-1 sm:hidden md:block">
              <h1 className="truncate text-2xl font-bold font-medium text-gray-900">{user.fullName}</h1>
            </div>
            <div className="mt-6 flex flex-col justify-stretch space-y-3 sm:flex-row sm:space-y-0 sm:space-x-4">
              <button
                type="button"
                className="inline-flex justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 ring-1 shadow-xs ring-gray-300 ring-inset hover:bg-gray-50"
              >
                <EnvelopeIcon aria-hidden="true" className="mr-1.5 -ml-0.5 size-5 text-gray-400" />
                <span>Message</span>
              </button>
              <button
                type="button"
                className="inline-flex justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 ring-1 shadow-xs ring-gray-300 ring-inset hover:bg-gray-50"
              >
                <PhoneIcon aria-hidden="true" className="mr-1.5 -ml-0.5 size-5 text-gray-400" />
                <span>Call</span>
              </button>
            </div>
          </div>
        </div>
        <div className="mt-6 hidden min-w-0 flex-1 sm:block md:hidden">
          <h1 className="truncate text-2xl font-bold font-medium text-gray-900">{user.fullName}</h1>
        </div>
      </div>
    </div>
  );
}
