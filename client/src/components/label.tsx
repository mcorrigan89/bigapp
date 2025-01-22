import React from "react";

export interface LabelProps
  extends React.LabelHTMLAttributes<HTMLLabelElement> {}

const Label = React.forwardRef<HTMLLabelElement, LabelProps>(
  ({ className, ...props }, ref) => {
    return <label className={"text-mono-1400"} ref={ref} {...props} />;
  },
);
Label.displayName = "label";

export { Label };
