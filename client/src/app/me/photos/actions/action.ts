"use server";

import { createCollectionSchema } from "./schema";
import { actionClient } from "@/lib/safe-actions";
import { createCollection } from "@/api/client";
import { revalidatePath } from "next/cache";

export const createCollectionAction = actionClient
  .schema(createCollectionSchema)
  .action(async ({ parsedInput: { collectionName } }) => {
    await createCollection(collectionName);

    revalidatePath("/me/photos");
  });
