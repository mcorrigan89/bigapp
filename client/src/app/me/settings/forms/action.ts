"use server";

import { returnValidationErrors } from "next-safe-action";
import { updateUserSchema } from "./schema";
import { actionClient } from "@/lib/safe-actions";
import { updateUser, uploadAvatarImage } from "@/api/client";
import { ErrorCode } from "@/api/gen/common/v1/errors_pb";
import { revalidatePath } from "next/cache";

export const updateUserAction = actionClient
  .schema(updateUserSchema)
  .action(async ({ parsedInput: { userId, givenName, familyName, email, handle, imageFile } }) => {
    if (imageFile?.size && imageFile.size > 0) {
      await uploadAvatarImage({
        file: imageFile,
      });
    }

    const response = await updateUser({
      id: userId,
      handle: handle,
      email: email,
      familyName: familyName,
      givenName: givenName,
    });

    if (response.error) {
      if (response.error.code === ErrorCode.EMAIL_EXISTS) {
        returnValidationErrors(updateUserSchema, {
          email: {
            _errors: ["Email already exists"],
          },
        });
      }
      if (response.error.code === ErrorCode.USER_HANDLE_EXISTS) {
        returnValidationErrors(updateUserSchema, {
          handle: {
            _errors: ["Handle already taken"],
          },
        });
      }
    }
    revalidatePath("/me/settings");
  });
