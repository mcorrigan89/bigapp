"use server";

import { redirect } from "next/navigation";
import { parseWithZod } from "@conform-to/zod";
import { createUserSchema } from "./schema";
import { cookies } from "next/headers";
import { createUser } from "@/api/client";

export async function createUserAction(prevState: unknown, formData: FormData) {
  const submission = parseWithZod(formData, {
    schema: createUserSchema,
  });

  if (submission.status !== "success") {
    return submission.reply();
  }

  try {
    const response = await createUser({
      email: submission.value.email,
      givenName: submission.value.givenName,
      familyName: submission.value.familyName,
    });

    if (response.error) {
      return submission.reply({
        formErrors: [response.error.message],
      });
    }
    if (!response.session?.token) {
      return submission.reply({
        formErrors: ["No session token received"],
      });
    }
    const cookieStore = await cookies();
    cookieStore.set("x-session-token", response.session.token);
  } catch (err) {
    console.error(err);
    return submission.reply();
  } finally {
    redirect("/");
  }
}
