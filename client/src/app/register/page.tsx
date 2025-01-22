"use client"; // form.tsx

import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import { useActionState } from "react";
import { createUserAction } from "./action";
import { createUserSchema } from "./schema";

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
    <div>
      <form id={form.id} onSubmit={form.onSubmit} action={action} noValidate>
        <div>
          <label>Email</label>
          <input
            type="email"
            key={fields.email.key}
            name={fields.email.name}
            defaultValue={fields.email.initialValue}
          />
          <div>{fields.email.errors}</div>
        </div>
        <div>
          <label>First Name</label>
          <input
            type="text"
            key={fields.givenName.key}
            name={fields.givenName.name}
            defaultValue={fields.givenName.initialValue}
          />
          <div>{fields.givenName.errors}</div>
        </div>
        <div>
          <label>Last Name</label>
          <input
            type="text"
            key={fields.familyName.key}
            name={fields.familyName.name}
            defaultValue={fields.familyName.initialValue}
          />
          <div>{fields.familyName.errors}</div>
        </div>
        <button>Register</button>
      </form>
    </div>
  );
}
