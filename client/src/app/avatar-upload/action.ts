"use server";

import { redirect } from "next/navigation";
import { parseWithZod } from "@conform-to/zod";
import { imageSchema } from "./schema";
import { uploadImage } from "@/api/client";

export async function uploadImageAction(
  prevState: unknown,
  formData: FormData,
) {
  const submission = parseWithZod(formData, {
    schema: imageSchema,
  });

  if (submission.status !== "success") {
    return submission.reply();
  }

  try {
    await uploadImage({
      file: submission.value.file,
    });
  } catch (err) {
    console.error(err);
    return submission.reply();
  } finally {
    redirect("/");
  }
}
