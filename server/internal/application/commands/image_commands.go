package commands

import (
	"io"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
)

type CreateNewAvatarImageCommand struct {
	UserID   uuid.UUID
	BucketID string
	ObjectID string
	File     io.Reader
	Size     int64
}

type ImageUploadData struct {
	ObjectID string
	File     io.Reader
	Size     int64
}

type UploadImagesToCollectionCommand struct {
	UserID       uuid.UUID
	CollectionID uuid.UUID
	BucketID     string
	Files        []ImageUploadData
}

type CreateNewCollectionCommand struct {
	OwnerID uuid.UUID
	Name    string
}

func (c CreateNewCollectionCommand) ToDomain() *entities.CollectionEntity {
	collection := entities.CollectionEntity{
		ID:      uuid.New(),
		Name:    c.Name,
		Public:  false,
		OwnerID: c.OwnerID,
	}
	return &collection
}
