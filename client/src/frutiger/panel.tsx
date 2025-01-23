import React from "react";

export const FrutigerPanel = ({ children }: { children: React.ReactNode }) => {
  return (
    <div
      className="
      relative
      overflow-hidden
      rounded-xl
      border
      border-white/30
      bg-white/10
      p-6
      shadow-lg
      backdrop-blur-md
    "
    >
      <div
        className="
        absolute
        inset-0
        bg-gradient-to-br
        from-blue-200/30
        via-white/20
        to-blue-100/20
      "
      />
      <div className="relative">{children}</div>
    </div>
  );
};
