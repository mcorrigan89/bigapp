import { z } from "zod";

export const uploadImageToCollectionSchema = z.object({
  collectionId: z.string().uuid(),
  imageFiles: typeof window !== "undefined" ? z.instanceof(FileList) : z.any().refine((f) => f as FileList),
});
