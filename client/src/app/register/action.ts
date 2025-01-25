"use server";

import { returnValidationErrors } from "next-safe-action";
import { createUserSchema } from "./schema";
import { actionClient } from "@/lib/safe-actions";
import { createUser } from "@/api/client";
import { ErrorCode } from "@/api/gen/common/v1/errors_pb";
import { redirect } from "next/navigation";
import { cookies } from "next/headers";

export const createUserAction = actionClient
  .schema(createUserSchema)
  .action(async ({ parsedInput: { givenName, familyName, email } }) => {
    const response = await createUser({
      email: email,
      familyName: familyName,
      givenName: givenName,
    });

    if (response.error) {
      if (response.error.code === ErrorCode.EMAIL_EXISTS) {
        returnValidationErrors(createUserSchema, {
          email: {
            _errors: ["Email already exists"],
          },
        });
      }
    }

    if (!response.session?.token) {
      returnValidationErrors(createUserSchema, {
        _errors: ["No session token received"],
      });
    }

    const cookieStore = await cookies();
    cookieStore.set("x-session-token", response.session.token);

    redirect("/me");
  });
