import { Image as ImageType } from "@/api/gen/media/v1/image_pb";
import { env } from "@/env";
import Image from "next/image";

interface GridItem<T> {
  item: T;
  width: number;
  height: number;
  top: number;
}

const sumHeightInArray = (items: Array<GridItem<ImageType>>) => {
  if (items.length > 0) {
    return items[items.length - 1].top + items[items.length - 1].height;
  } else {
    return 0;
  }
};

const reorderItemArray = (columns: Array<Array<GridItem<ImageType>>>) => {
  return columns.reduce(
    (acc: { column: number; height: number }, curr, idx) => {
      const height = sumHeightInArray(curr);
      if (height < acc.height) {
        return { column: idx, height };
      } else {
        return acc;
      }
    },
    { column: 0, height: sumHeightInArray(columns[0]) },
  );
};

function layout(images: ImageType[], numberOfColumns: number) {
  const columnWidth = 100;
  const columns: Array<Array<GridItem<ImageType>>> = [];
  for (let i = 0; i < numberOfColumns; i++) {
    columns.push([]);
  }

  for (let i = 0; i < images.length; i++) {
    const image = images[i];
    const { column, height } = reorderItemArray(columns);
    const processedItem = {
      id: image.id,
      item: { ...image },
      top: height,
      left: (column % numberOfColumns) * columnWidth + (column % numberOfColumns),
      width: columnWidth,
      height: (image.height / image.width) * columnWidth,
    };
    columns[column].push(processedItem);
  }

  return columns;
}

interface Props {
  images: ImageType[];
  numberOfColumns?: number;
}

export function GalleryComponent({ images, numberOfColumns = 3 }: Props) {
  const imageColumns = layout(images, numberOfColumns);
  const columnWidth = `${100 / numberOfColumns}%`;
  return (
    <div className="flex w-[calc(100%-3rem)] text-white">
      <div className="flex gap-4" style={{ alignItems: "flex-start" }}>
        {imageColumns.map((column, columnIndex) => (
          <div key={columnIndex} style={{ width: columnWidth }} className="shrink-0">
            <div className="flex flex-col gap-4">
              {column.map((image, imageIndex) => (
                <div key={imageIndex} className="w-full overflow-hidden rounded-lg bg-gray-200">
                  <Image
                    src={env.NEXT_PUBLIC_SERVER_URL + image.item.url}
                    width={image.item.width}
                    height={image.item.height}
                    alt={`Image ${imageIndex + 1}`}
                    className="h-auto w-full object-cover"
                  />
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
