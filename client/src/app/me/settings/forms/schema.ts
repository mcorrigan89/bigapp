import { z } from "zod";
import { zfd } from "zod-form-data";

const MAX_FILE_SIZE = 5_000_000;

export const updateUserSchema = z.object({
  imageFile: zfd
    .file()
    .refine((file) => file.size !== 0, "File is required")
    .refine((file) => file.size < MAX_FILE_SIZE, "Max size is 5MB.")
    .optional(),
  userId: z.string().uuid(),
  givenName: z.string().min(1, "Given name is required"),
  familyName: z.string().min(1, "Family name is required"),
  email: z.string().email("Invalid email"),
  handle: z.string().min(6, "Handle is required"),
});
