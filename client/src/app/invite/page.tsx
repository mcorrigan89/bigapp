"use client";

import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import { useActionState } from "react";
import { inviteAction } from "./action";
import { inviteSchema } from "./schema";
import { Button } from "@/components/button";
import { Input } from "@/components/input";
import { Label } from "@/components/label";
import { FormError } from "@/components/form-error";

export default function InvitePage() {
  const [lastResult, action] = useActionState(inviteAction, undefined);
  const [form, fields] = useForm({
    // Sync the result of last submission
    lastResult,

    // Reuse the validation logic on the client
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: inviteSchema });
    },

    // Validate the form on blur event triggered
    shouldValidate: "onBlur",
    shouldRevalidate: "onInput",
  });

  return (
    <>
      <div className="flex h-screen flex-col items-center justify-center">
        <form
          id={form.id}
          onSubmit={form.onSubmit}
          action={action}
          noValidate
          className="flex flex-col gap-4"
        >
          <div className="text-mono-1400 flex flex-col gap-1">
            <Label htmlFor="email" className="text-md">
              Email
            </Label>
            <Input
              id="email"
              type="email"
              key={fields.email.key}
              name={fields.email.name}
              defaultValue={fields.email.initialValue}
              data-error={fields.email.errors && fields.email.errors.length > 0}
            />
            {fields.email.errors?.map((error) => (
              <FormError key={error}>{error}</FormError>
            ))}
          </div>
          {form.errors?.map((error) => (
            <FormError key={error}>{error}</FormError>
          ))}
          <Button variant={"default"}>Login</Button>
        </form>
      </div>
    </>
  );
}
