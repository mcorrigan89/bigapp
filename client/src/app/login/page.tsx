"use client"; // form.tsx

import { useForm } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import { useActionState } from "react";
import { loginAction } from "./action";
import { loginSchema } from "./schema";

export default function LoginPage() {
  const [lastResult, action] = useActionState(loginAction, undefined);
  const [form, fields] = useForm({
    // Sync the result of last submission
    lastResult,

    // Reuse the validation logic on the client
    onValidate({ formData }) {
      return parseWithZod(formData, { schema: loginSchema });
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
        <button>Login</button>
      </form>
    </div>
  );
}
