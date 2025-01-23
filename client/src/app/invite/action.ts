"use server";

import { redirect } from "next/navigation";
import { parseWithZod } from "@conform-to/zod";
import { inviteSchema } from "./schema";
import { inviteUser } from "@/api/client";

export async function inviteAction(prevState: unknown, formData: FormData) {
  const submission = parseWithZod(formData, {
    schema: inviteSchema,
  });

  if (submission.status !== "success") {
    return submission.reply();
  }

  try {
    const response = await inviteUser({
      email: submission.value.email,
    });
    if (response.status !== "OK") {
      return submission.reply();
    }
  } catch (err) {
    console.error(err);
    return submission.reply();
  } finally {
    redirect("/");
  }
}
