import { z } from "zod";

export const createCollectionSchema = z.object({
  collectionName: z.string().min(1, "Collection name is required"),
});
