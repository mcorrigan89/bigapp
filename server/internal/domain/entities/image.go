package entities

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

var (
	ErrImageNotFound           = fmt.Errorf("image not found")
	ErrCollectionNotFound      = fmt.Errorf("image not found")
	ErrCollectionNotAuthorized = fmt.Errorf("collection not authorized")
)

type ImageEntity struct {
	ID         uuid.UUID
	BucketName string
	ObjectID   string
	Width      int32
	Height     int32
	Size       int32
}

func NewImageEntity(model models.Image) *ImageEntity {
	return &ImageEntity{
		ID:         model.ID,
		BucketName: model.BucketName,
		ObjectID:   model.ObjectID,
		Width:      model.Width,
		Height:     model.Height,
		Size:       model.FileSize,
	}
}

func (p *ImageEntity) UrlSlug() string {
	url := fmt.Sprintf("/image/%s", p.ID)
	return url
}

type CollectionEntity struct {
	ID      uuid.UUID
	Name    string
	OwnerID uuid.UUID
	Public  bool
	Images  []*ImageEntity
}

func NewCollectionEntity(model models.Collection, images []*ImageEntity) *CollectionEntity {
	return &CollectionEntity{
		ID:      model.ID,
		Name:    model.CollectionName,
		OwnerID: model.OwnerID,
		Public:  model.Public,
		Images:  images,
	}
}
