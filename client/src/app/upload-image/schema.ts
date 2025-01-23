import { z } from "zod";

const MAX_FILE_SIZE = 5_000_000;

export const imageSchema = z.object({
  file: z
    .instanceof(File)
    .refine((file) => file.size !== 0, "File is required")
    .refine((file) => file.size < MAX_FILE_SIZE, "Max size is 5MB."),
});
