import React from "react";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "@/lib/utils";

const buttonVariants = cva(
  "rounded font-medium text-white hover:cursor-pointer",
  {
    variants: {
      variant: {
        default: "bg-primary-800 hover:bg-primary-700 active:bg-primary-900",
        positive:
          "bg-positive-800 hover:bg-positive-700 active:bg-positive-900",
        negative:
          "bg-negative-800 hover:bg-negative-700 active:bg-negative-900",
      },
      size: {
        default: "h-10 px-4",
        small: "h-8 px-3",
        large: "h-12 px-6",
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
