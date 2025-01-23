import { Image as ImageType } from "@/api/gen/media/v1/image_pb";
import { env } from "@/env";
import { cn } from "@/lib/utils";
import Image from "next/image";
import { cva, type VariantProps } from "class-variance-authority";

const avatarVariants = cva(
  "overflow-hidden rounded-full bg-slate-400 ring-1 ring-indigo-600",
  {
    variants: {
      size: {
        small: "h-8 w-8",
        medium: "h-12 w-12",
        large: "h-16 w-16",
      },
    },
    defaultVariants: {
      size: "medium",
    },
  },
);

interface AvatarProps extends VariantProps<typeof avatarVariants> {
  avatar?: ImageType;
}

export function Avatar({ avatar, size }: AvatarProps) {
  if (avatar) {
    return (
      <div className={cn(avatarVariants({ size }))}>
        <Image
          alt={"profile picture"}
          src={env.NEXT_PUBLIC_SERVER_URL + avatar.url}
          width={avatar.width}
          height={avatar.height}
        />
      </div>
    );
  } else {
    return <div className={cn(avatarVariants({ size }))} />;
  }
}
