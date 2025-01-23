import React from "react";

export const FrutigerButton = ({ children }: { children: React.ReactNode }) => {
  return (
    <button
      className="
      relative
      overflow-hidden
      rounded-full
      border
      border-white/30
      bg-gradient-to-b
      from-blue-300/80
      to-blue-500/80
      px-6
      py-2
      font-sans
      text-white
      shadow-lg
      backdrop-blur-sm
      transition-all
      duration-300
      hover:from-blue-200/90
      hover:to-blue-400/90
      hover:shadow-blue-300/50
      active:scale-95
    "
    >
      <div className="absolute inset-0 rounded-full bg-gradient-to-t from-white/5 to-white/30" />
      <span className="relative">{children}</span>
    </button>
  );
};
