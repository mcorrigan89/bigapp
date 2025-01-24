"use server";

import { redirect } from "next/navigation";
import { parseWithZod } from "@conform-to/zod";
import { loginSchema } from "./schema";
import { loginEmail } from "@/api/client";

export async function loginAction(prevState: unknown, formData: FormData) {
  const submission = parseWithZod(formData, {
    schema: loginSchema,
  });

  if (submission.status !== "success") {
    return submission.reply();
  }

  try {
    const response = await loginEmail({
      email: submission.value.email,
    });
    if (response.error) {
      return submission.reply({
        formErrors: [response.error.message],
      });
    }
  } catch (err) {
    console.error(err);
    return submission.reply();
  } finally {
    redirect("/");
  }
}
