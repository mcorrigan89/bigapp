import { Image as ImageType } from "@/api/gen/media/v1/image_pb";
import { env } from "@/env";
import Image from "next/image";

interface Props {
  images: ImageType[];
  numberOfColumns?: number;
}

export function GalleryPreview({ images }: Props) {
  return (
    <div className="flex aspect-square overflow-clip rounded-xl">
      <GalleryPreviewSwitch images={images} />
    </div>
  );
}
export function GalleryPreviewSwitch({ images }: Props) {
  if (images.length > 2) {
    return (
      <div className="grid h-full grid-cols-2 gap-2">
        <Image
          src={env.NEXT_PUBLIC_SERVER_URL + images[0].url}
          alt={"preview image"}
          width={images[0].width}
          height={images[0].height}
          className="cols-span-2 h-full w-full object-cover"
        />
        <div className="grid grid-rows-2 gap-2">
          <Image
            src={env.NEXT_PUBLIC_SERVER_URL + images[1].url}
            alt={"preview image"}
            width={images[0].width}
            height={images[0].height}
            className="h-full w-full object-cover"
          />
          <Image
            src={env.NEXT_PUBLIC_SERVER_URL + images[2].url}
            alt={"preview image"}
            width={images[0].width}
            height={images[0].height}
            className="h-full w-full object-cover"
          />
        </div>
      </div>
    );
  } else if (images.length > 1) {
    return (
      <div className="grid grid-cols-2 gap-2">
        <Image
          src={env.NEXT_PUBLIC_SERVER_URL + images[0].url}
          alt={"preview image"}
          width={images[0].width}
          height={images[0].height}
          className="h-full w-full object-cover"
        />
        <Image
          src={env.NEXT_PUBLIC_SERVER_URL + images[1].url}
          alt={"preview image"}
          width={images[0].width}
          height={images[0].height}
          className="h-full w-full object-cover"
        />
      </div>
    );
  } else if (images.length > 0) {
    return (
      <div className="relative">
        <Image
          src={env.NEXT_PUBLIC_SERVER_URL + images[0].url}
          alt={"preview image"}
          width={images[0].width}
          height={images[0].height}
          className="h-full w-full object-cover"
        />
      </div>
    );
  }
  return <div className="aspect-square grid-cols-2 overflow-clip rounded-xl">No Images</div>;
}
