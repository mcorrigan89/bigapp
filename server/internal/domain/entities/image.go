package entities

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/infrastructure/postgres/models"
)

var (
	ErrImageNotFound = fmt.Errorf("image not found")
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
