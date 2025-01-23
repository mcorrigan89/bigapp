import { cn } from "@/lib/utils";
import React from "react";

export type FormErrorProps = React.BaseHTMLAttributes<HTMLDivElement>;

const FormError = React.forwardRef<HTMLDivElement, FormErrorProps>(
  ({ className, ...props }, ref) => {
    return (
      <div
        className={cn("px-1.5 text-sm text-rose-500", className)}
        ref={ref}
        {...props}
      >
        {props.children}
      </div>
    );
  },
);
FormError.displayName = "formError";

export { FormError };
