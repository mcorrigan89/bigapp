import { getCollectiondById } from "@/api/client";
import { redirect } from "next/navigation";
import { UploadImagesToCollection } from "./actions/upload-image";
import { GalleryComponent } from "@/components/gallery/gallery-component";

interface Props {
  params: Promise<{
    collectionId: string;
  }>;
}

export default async function PhotosPage({ params }: Props) {
  const { collectionId } = await params;
  const response = await getCollectiondById(collectionId);
  const collection = response?.collection;
  if (!collection) {
    return redirect("/me/photos");
  }
  const images = collection.images;

  return (
    <>
      <div className="m-12 md:flex md:items-center md:justify-between">
        <div className="min-w-0 flex-1">
          <h2 className="text-2xl/7 font-bold text-white sm:truncate sm:text-3xl sm:tracking-tight">
            {collection.name}
          </h2>
        </div>
        <div className="mt-4 flex md:mt-0 md:ml-4">
          <UploadImagesToCollection collectionId={collectionId} />
        </div>
      </div>

      <div className="p-8">
        <GalleryComponent images={images} />
      </div>
    </>
  );
}
