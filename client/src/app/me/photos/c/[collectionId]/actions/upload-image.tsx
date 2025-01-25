"use client";

import { useHookFormAction } from "@next-safe-action/adapter-react-hook-form/hooks";
import { uploadImageToCollectionAction } from "./upload-image-action";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRef } from "react";
import { uploadImageToCollectionSchema } from "./upload-image-schema";

export function UploadImagesToCollection({ collectionId }: { collectionId: string }) {
  const fileInputRef = useRef<HTMLInputElement>(null);

  const { form, handleSubmitWithAction, resetFormAndAction } = useHookFormAction(
    uploadImageToCollectionAction,
    zodResolver(uploadImageToCollectionSchema),
    {
      actionProps: {
        onSuccess: () => {
          window.alert("Logged in successfully!");
          resetFormAndAction();
        },
      },
      formProps: {
        defaultValues: {
          collectionId,
        },
      },
      errorMapProps: {},
    },
  );

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    form.setValue("imageFiles", e.target.files);
    handleSubmitWithAction();
  };

  return (
    <form onSubmit={handleSubmitWithAction} noValidate>
      <input {...form.register("collectionId")} type="hidden" />
      <input
        type="file"
        className="hidden"
        {...form.register("imageFiles")}
        multiple
        ref={fileInputRef}
        accept="image/*"
        onChange={onChange}
      />
      <label htmlFor="imageFile">
        <button
          onClick={() => fileInputRef.current?.click()}
          type="button"
          className="ml-3 inline-flex items-center rounded-md bg-indigo-500 px-3 py-2 text-sm font-semibold text-white shadow-xs hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
        >
          Upload Images
        </button>
      </label>
    </form>
  );
}
