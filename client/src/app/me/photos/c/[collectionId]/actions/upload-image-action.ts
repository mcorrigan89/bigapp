"use server";
import { actionClient } from "@/lib/safe-actions";
import { uploadImageToCollectionSchema } from "./upload-image-schema";
import { uploadImagesToCollection } from "@/api/client";
import { revalidatePath } from "next/cache";

export const uploadImageToCollectionAction = actionClient
  .schema(uploadImageToCollectionSchema)
  .action(async ({ parsedInput: { collectionId, imageFiles } }) => {
    const files = imageFiles as File[];

    await uploadImagesToCollection({
      collectionId,
      files: files,
    });
    revalidatePath(`/me/photos/c/[collectionId]`);
  });
