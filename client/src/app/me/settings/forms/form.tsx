"use client";

import { ChangeEventHandler, useRef, useState } from "react";
import { Image as ImageType } from "@/api/gen/media/v1/image_pb";
import { env } from "@/env";
import { updateUserAction } from "./action";
import { updateUserSchema } from "./schema";
import { User } from "@/api/gen/user/v1/user_pb";
import { useHookFormAction } from "@next-safe-action/adapter-react-hook-form/hooks";
import { zodResolver } from "@hookform/resolvers/zod";
import clsx from "clsx";

export function ProfileInfoForm({ user, avatar }: { user: User; avatar: ImageType | undefined }) {
  const [showAvatar, setShowAvatar] = useState(!!avatar);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const imgRef = useRef<HTMLImageElement>(null);

  const handlePhotoChange: ChangeEventHandler<HTMLInputElement> = () => {
    if (fileInputRef.current && imgRef.current) {
      const file = fileInputRef.current.files?.item(0);
      if (file) {
        imgRef.current.src = URL.createObjectURL(file);
        form.setValue("imageFile", file);
        setShowAvatar(true);
      }
    }
  };

  const { form, handleSubmitWithAction, resetFormAndAction } = useHookFormAction(
    updateUserAction,
    zodResolver(updateUserSchema),
    {
      actionProps: {
        onSuccess: () => {
          window.alert("Logged in successfully!");
          resetFormAndAction();
        },
      },
      formProps: {
        defaultValues: {
          userId: user.id,
          givenName: user.givenName,
          familyName: user.familyName,
          email: user.email,
          handle: user.handle,
          imageFile: undefined,
        },
      },
      errorMapProps: {},
    },
  );

  const errors = form.formState.errors;

  return (
    <>
      <form onSubmit={handleSubmitWithAction} noValidate className="md:col-span-2">
        <div className="grid grid-cols-1 gap-x-6 gap-y-8 sm:max-w-xl sm:grid-cols-6">
          <div className="col-span-full flex items-center gap-x-8">
            <img
              ref={imgRef}
              src={env.NEXT_PUBLIC_SERVER_URL + avatar?.url}
              alt="profile photo"
              className={clsx(
                "size-24 flex-none rounded-lg bg-gray-800 object-cover",
                showAvatar ? "flex-none" : "hidden",
              )}
            />
            <div className={clsx("size-24 rounded-lg bg-gray-800", showAvatar ? "hidden" : "flex-none")} />
            <div>
              <input
                type="file"
                className="hidden"
                {...form.register("imageFile")}
                onChange={handlePhotoChange}
                multiple={false}
                ref={fileInputRef}
                accept="image/*"
              />
              <label htmlFor="imageFile">
                <button
                  onClick={() => fileInputRef.current?.click()}
                  type="button"
                  className="rounded-md bg-white/10 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-white/20"
                >
                  Change avatar
                </button>
              </label>
              <p className="mt-2 text-xs/5 text-gray-400">JPG, GIF or PNG. 1MB max.</p>
            </div>
          </div>

          <input {...form.register("userId")} type="hidden" />
          <div className="sm:col-span-3">
            <label htmlFor="first-name" className="block text-sm/6 font-medium text-white">
              First name
            </label>
            <div className="mt-2">
              <input
                {...form.register("givenName")}
                type="text"
                autoComplete="given-name"
                className="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
              />
            </div>
            {errors.givenName?.message ? (
              <p id="givenName-error" className="mt-2 text-sm text-red-600">
                {errors.givenName?.message}
              </p>
            ) : null}
          </div>

          <div className="sm:col-span-3">
            <label htmlFor="last-name" className="block text-sm/6 font-medium text-white">
              Last name
            </label>
            <div className="mt-2">
              <input
                {...form.register("familyName")}
                type="text"
                autoComplete="family-name"
                className="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
              />
            </div>
            {errors.familyName?.message ? (
              <p id="email-error" className="mt-2 text-sm text-red-600">
                {errors.email?.message}
              </p>
            ) : null}
          </div>

          <div className="col-span-full">
            <label htmlFor="email" className="block text-sm/6 font-medium text-white">
              Email address
            </label>
            <div className="mt-2">
              <input
                {...form.register("email")}
                type="email"
                autoComplete="email"
                className="block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
              />
            </div>
            {errors.email?.message ? (
              <p id="email-error" className="mt-2 text-sm text-red-600">
                {errors.email?.message}
              </p>
            ) : null}
          </div>

          <div className="col-span-full">
            <label htmlFor="handle" className="block text-sm/6 font-medium text-white">
              Handle
            </label>
            <div className="mt-2">
              <div className="flex items-center rounded-md bg-white/5 pl-3 outline-1 -outline-offset-1 outline-white/10 focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-500">
                <div className="shrink-0 text-base text-gray-500 select-none sm:text-sm/6">example.com/user/</div>
                <input
                  {...form.register("handle")}
                  type="text"
                  className="block min-w-0 grow bg-transparent py-1.5 pr-3 pl-1 text-base text-white placeholder:text-gray-500 focus:outline-none sm:text-sm/6"
                />
              </div>
              {errors.handle?.message ? (
                <p id="email-error" className="mt-2 text-sm text-red-600">
                  {errors.handle?.message}
                </p>
              ) : null}
            </div>
          </div>
        </div>

        <div className="mt-8 flex">
          <button
            type="submit"
            className="rounded-md bg-indigo-500 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
          >
            Save
          </button>
        </div>
      </form>
    </>
  );
}
