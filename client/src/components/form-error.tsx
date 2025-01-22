import React from "react";

export interface FormErrorProps
  extends React.BaseHTMLAttributes<HTMLDivElement> {}

const FormError = React.forwardRef<HTMLDivElement, FormErrorProps>(
  ({ className, ...props }, ref) => {
    return (
      <div className={"px-1.5 text-sm text-negative-500"} ref={ref} {...props}>
        {props.children}
      </div>
    );
  },
);
FormError.displayName = "formError";

export { FormError };
