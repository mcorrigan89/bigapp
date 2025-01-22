"use client"; // form.tsx

import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import { useActionState } from "react";
import { createUserAction } from "./action";
import { createUserSchema } from "./schema";
import { Label } from "@/components/label";
import { Input } from "@/components/input";
import { FormError } from "@/components/form-error";
import { Button } from "@/components/button";

export default function RegisterPage() {
  const [lastResult, action] = useActionState(createUserAction, undefined);
  const [form, fields] = useForm({
    // Sync the result of last submission
    lastResult,

    // Reuse the validation logic on the client
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: createUserSchema });
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
        <div>
          <Label htmlFor="email">Email</Label>
          <Input
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
        <div>
          <Label htmlFor="givenName">First Name</Label>
          <Input
            type="text"
            key={fields.givenName.key}
            name={fields.givenName.name}
            defaultValue={fields.givenName.initialValue}
            data-error={
              fields.givenName.errors && fields.givenName.errors.length > 0
            }
          />
          <div>{fields.givenName.errors}</div>
        </div>
        <div>
          <Label htmlFor="familyName">Last Name</Label>
          <Input
            type="text"
            key={fields.familyName.key}
            name={fields.familyName.name}
            defaultValue={fields.familyName.initialValue}
            data-error={
              fields.familyName.errors && fields.familyName.errors.length > 0
            }
          />
          {fields.familyName.errors?.map((error) => (
            <FormError key={error}>{error}</FormError>
          ))}
        </div>
        <Button>Register</Button>
      </form>
    </div>
  );
}
