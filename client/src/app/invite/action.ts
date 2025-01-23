"use server";

import { redirect } from "next/navigation";
import { parseWithZod } from "@conform-to/zod";
import { inviteSchema } from "./schema";
import { inviteUser } from "@/api/client";
import { ErrorCode } from "@/api/gen/user/v1/user_pb";
import { ErrorHandler, handleServiceCall } from "@/api/handlers";

const handlers: ErrorHandler = {
  [ErrorCode.EMAIL_EXISTS]: (msg) => console.error(msg),
};

export async function inviteAction(prevState: unknown, formData: FormData) {
  const submission = parseWithZod(formData, {
    schema: inviteSchema,
  });

  if (submission.status !== "success") {
    return submission.reply();
  }

  const response = await handleServiceCall(
    inviteUser({
      email: submission.value.email,
    }),
    handlers,
  );
  if (response.error) {
    return submission.reply({
      formErrors: [response.error.message],
    });
  }

  redirect("/");
}
