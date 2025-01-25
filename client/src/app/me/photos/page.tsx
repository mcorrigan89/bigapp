import { getCollectiondByOwnerToken } from "@/api/client";
import { GalleryPreview } from "@/components/gallery/gallery-preview";
import Link from "next/link";
import CreateCollectionModal from "./actions/create-collection-modal";

export default async function PhotosPage() {
  const response = await getCollectiondByOwnerToken();

  const collections = response?.collections ?? [];

  return (
    <>
      <div className="m-12 md:flex md:items-center md:justify-between">
        <div className="min-w-0 flex-1">
          <h2 className="text-2xl/7 font-bold text-white sm:truncate sm:text-3xl sm:tracking-tight">My Photos</h2>
        </div>
        <div className="mt-4 flex md:mt-0 md:ml-4">
          <CreateCollectionModal />
        </div>
      </div>
      <div className="grid grid-cols-4 gap-8">
        {collections.map((collection) => {
          return (
            <div key={collection.id}>
              <Link href={`/me/photos/c/${collection.id}`} className="text-white">
                <div className="text-white">{collection.name}</div>
                <GalleryPreview images={collection.images} />
              </Link>
            </div>
          );
        })}
      </div>
    </>
  );
}
