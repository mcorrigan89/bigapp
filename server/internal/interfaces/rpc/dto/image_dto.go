package dto

import (
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	imagev1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/media/v1"
)

func ImageEntityToDto(imageEntity *entities.ImageEntity) *imagev1.Image {
	image := &imagev1.Image{
		Id:     imageEntity.ID.String(),
		Url:    imageEntity.UrlSlug(),
		Width:  imageEntity.Width,
		Height: imageEntity.Height,
		Size:   imageEntity.Size,
	}

	return image
}

func CollectionEntityToDto(collectionEntity *entities.CollectionEntity) *imagev1.Collection {
	images := make([]*imagev1.Image, 0, len(collectionEntity.Images))
	for _, imageEntity := range collectionEntity.Images {
		image := ImageEntityToDto(imageEntity)
		images = append(images, image)
	}

	collection := &imagev1.Collection{
		Id:     collectionEntity.ID.String(),
		Name:   collectionEntity.Name,
		Images: images,
	}

	return collection
}
