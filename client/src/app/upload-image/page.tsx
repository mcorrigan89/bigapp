"use client";

import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import { useActionState } from "react";
import { uploadImageAction } from "./action";
import { imageSchema } from "./schema";
import { Button } from "@/components/button";
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { FormError } from "@/components/form-error";

export default function ImageUploadPage() {
  const [lastResult, action] = useActionState(uploadImageAction, undefined);
  const [form, fields] = useForm({
    // Sync the result of last submission
    lastResult,

    // Reuse the validation logic on the client
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: imageSchema });
    },

    // Validate the form on blur event triggered
    shouldValidate: "onBlur",
    shouldRevalidate: "onInput",
  });

  return (
    <div className="flex h-screen flex-col items-center justify-center">
      <form
        id={form.id}
        onSubmit={form.onSubmit}
        action={action}
        noValidate
        className="flex flex-col gap-4"
      >
        <div className="flex flex-col gap-1 text-mono-1400">
          <Label htmlFor="image" className="text-md">
            Image
          </Label>
          <Input
            id="image"
            type="file"
            key={fields.file.key}
            name={fields.file.name}
            data-error={fields.file.errors && fields.file.errors.length > 0}
          />
          {fields.file.errors?.map((error) => (
            <FormError key={error}>{error}</FormError>
          ))}
        </div>
        <Button variant={"default"}>Login</Button>
      </form>
    </div>
  );
}
