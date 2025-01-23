import React from "react";
import { cva, type VariantProps } from "class-variance-authority";
import { cn } from "@/lib/utils";

const inputVariants = cva(
  cn(
    "w-full min-w-16 rounded-md border-0 px-4 py-1.5 font-light",
    "text-slate-900 ring-1 ring-slate-400 outline-0 ring-inset placeholder:text-slate-700",
    "focus:ring-2 focus:ring-indigo-600 focus:ring-inset",
    "hover:ring-1 hover:ring-indigo-600 hover:ring-inset",
    "data-[error=true]:ring-1 data-[error=true]:ring-rose-500 data-[error=true]:ring-inset focus:data-[error=true]:ring-2",
    "disabled:cursor-not-allowed disabled:bg-slate-200 disabled:hover:ring-slate-300",
  ),
  {
    variants: {},
    defaultVariants: {},
  },
);

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement>,
    VariantProps<typeof inputVariants> {}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className, ...props }, ref) => {
    return (
      <input
        className={cn(inputVariants({ className }))}
        ref={ref}
        {...props}
      />
    );
  },
);
Input.displayName = "Input";

export { Input, inputVariants };
