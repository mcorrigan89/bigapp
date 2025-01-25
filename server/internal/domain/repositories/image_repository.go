package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type ImageRepository interface {
	GetImageByID(ctx context.Context, querier models.Querier, id uuid.UUID) (*entities.ImageEntity, error)
	CreateImage(ctx context.Context, querier models.Querier, image *entities.ImageEntity) (*entities.ImageEntity, error)
	GetCollectionByID(ctx context.Context, querier models.Querier, id uuid.UUID) (*entities.CollectionEntity, error)
	GetCollectionByOwnerID(ctx context.Context, querier models.Querier, ownerID uuid.UUID) ([]*entities.CollectionEntity, error)
	CreateCollection(ctx context.Context, querier models.Querier, collection *entities.CollectionEntity) (*entities.CollectionEntity, error)
	AddImageToCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error)
	RemoveImageFromCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error)
}
