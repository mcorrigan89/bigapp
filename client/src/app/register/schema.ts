import { z } from "zod";

export const createUserSchema = z.object({
  email: z.string().email(),
  familyName: z.string().optional(),
  givenName: z.string().optional(),
});
