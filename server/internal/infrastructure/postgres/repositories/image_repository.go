package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type postgresImageRepository struct {
}

func NewPostgresImageRepository() *postgresImageRepository {
	return &postgresImageRepository{}
}

func (repo *postgresImageRepository) GetImageByID(ctx context.Context, querier models.Querier, imageID uuid.UUID) (*entities.ImageEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetImageByID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	return entities.NewImageEntity(row.Image), nil
}

func (repo *postgresImageRepository) CreateImage(ctx context.Context, querier models.Querier, image *entities.ImageEntity) (*entities.ImageEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateImage(ctx, models.CreateImageParams{
		ID:         image.ID,
		BucketName: image.BucketName,
		ObjectID:   image.ObjectID,
		Width:      image.Width,
		Height:     image.Height,
		FileSize:   image.Size,
	})
	if err != nil {
		return nil, err
	}

	return entities.NewImageEntity(row), nil
}

func (repo *postgresImageRepository) GetCollectionByID(ctx context.Context, querier models.Querier, id uuid.UUID) (*entities.CollectionEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetCollectionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	imageRows, err := querier.GetCollectionImagesByCollectionID(ctx, row.Collection.ID)
	if err != nil {
		return nil, err
	}

	imageEntities := make([]*entities.ImageEntity, 0, len(imageRows))
	for _, imageRow := range imageRows {
		imageEntities = append(imageEntities, entities.NewImageEntity(imageRow.Image))
	}

	return entities.NewCollectionEntity(row.Collection, imageEntities), nil
}
func (repo *postgresImageRepository) GetCollectionByOwnerID(ctx context.Context, querier models.Querier, ownerID uuid.UUID) ([]*entities.CollectionEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	rows, err := querier.GetCollectionByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	collectionEntities := make([]*entities.CollectionEntity, 0, len(rows))

	for _, row := range rows {
		imageRows, err := querier.GetCollectionImagesByCollectionID(ctx, row.Collection.ID)
		if err != nil {
			return nil, err
		}

		imageEntities := make([]*entities.ImageEntity, 0, len(imageRows))
		for _, imageRow := range imageRows {
			imageEntities = append(imageEntities, entities.NewImageEntity(imageRow.Image))
		}

		collectionEntities = append(collectionEntities, entities.NewCollectionEntity(row.Collection, imageEntities))
	}

	return collectionEntities, nil
}

func (repo *postgresImageRepository) CreateCollection(ctx context.Context, querier models.Querier, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateCollection(ctx, models.CreateCollectionParams{
		ID:             collection.ID,
		CollectionName: collection.Name,
		OwnerID:        collection.OwnerID,
		Public:         collection.Public,
	})
	if err != nil {
		return nil, err
	}

	return entities.NewCollectionEntity(row, nil), nil
}

func (repo *postgresImageRepository) AddImageToCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	_, err := querier.AddImageToCollection(ctx, models.AddImageToCollectionParams{
		ImageID:      image.ID,
		CollectionID: collection.ID,
	})
	if err != nil {
		return nil, err
	}

	return repo.GetCollectionByID(ctx, querier, collection.ID)
}

func (repo *postgresImageRepository) RemoveImageFromCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	err := querier.RemoveImageFromCollection(ctx, models.RemoveImageFromCollectionParams{
		ImageID:      image.ID,
		CollectionID: collection.ID,
	})
	if err != nil {
		return nil, err
	}

	return repo.GetCollectionByID(ctx, querier, collection.ID)
}
