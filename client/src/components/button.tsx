import React from "react";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "@/lib/utils";

const buttonVariants = cva(
  "rounded text-sm font-medium text-white hover:cursor-pointer",
  {
    variants: {
      variant: {
        default: "bg-indigo-700 hover:bg-indigo-600 active:bg-indigo-800",
        positive: "bg-teal-700 hover:bg-teal-600 active:bg-teal-800",
        negative: "bg-rose-700 hover:bg-rose-600 active:bg-rose-800",
      },
      size: {
        small: "px-2 py-1",
        default: "px-2.5 py-1.5",
        large: "px-4 py-2",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  },
);

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  asChild?: boolean;
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, ...props }, ref) => {
    return (
      <button
        className={cn(buttonVariants({ variant, size, className }))}
        ref={ref}
        {...props}
      />
    );
  },
);
Button.displayName = "Button";

export { Button, buttonVariants };
