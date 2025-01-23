import { cn } from "@/lib/utils";
import React from "react";

export type LabelProps = React.LabelHTMLAttributes<HTMLLabelElement>;

const Label = React.forwardRef<HTMLLabelElement, LabelProps>(
  ({ className, ...props }, ref) => {
    return (
      <label className={cn("text-mono-1400", className)} ref={ref} {...props} />
    );
  },
);
Label.displayName = "label";

export { Label };
