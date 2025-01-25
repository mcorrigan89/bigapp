package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/domain/repositories"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type ImageService interface {
	GetImageByID(ctx context.Context, querier models.Querier, imageID uuid.UUID) (*entities.ImageEntity, error)
	CreateImage(ctx context.Context, querier models.Querier, image *entities.ImageEntity) (*entities.ImageEntity, error)
	GetCollectionByID(ctx context.Context, querier models.Querier, collectionID uuid.UUID) (*entities.CollectionEntity, error)
	GetCollectionByOwnerID(ctx context.Context, querier models.Querier, ownerID uuid.UUID) ([]*entities.CollectionEntity, error)
	CreateCollection(ctx context.Context, querier models.Querier, collection *entities.CollectionEntity) (*entities.CollectionEntity, error)
	AddImageToCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error)
}

type imageService struct {
	imageRepo repositories.ImageRepository
}

func NewImageService(imageRepo repositories.ImageRepository) *imageService {
	return &imageService{imageRepo: imageRepo}
}

func (s *imageService) GetImageByID(ctx context.Context, querier models.Querier, imageID uuid.UUID) (*entities.ImageEntity, error) {
	imageEntity, err := s.imageRepo.GetImageByID(ctx, querier, imageID)
	if err != nil {
		return nil, err
	}

	return imageEntity, nil
}

func (s *imageService) CreateImage(ctx context.Context, querier models.Querier, image *entities.ImageEntity) (*entities.ImageEntity, error) {
	imageEntity, err := s.imageRepo.CreateImage(ctx, querier, image)
	if err != nil {
		return nil, err
	}

	return imageEntity, nil
}

func (s *imageService) GetCollectionByID(ctx context.Context, querier models.Querier, imageID uuid.UUID) (*entities.CollectionEntity, error) {
	collectionEntity, err := s.imageRepo.GetCollectionByID(ctx, querier, imageID)
	if err != nil {
		return nil, err
	}

	return collectionEntity, nil
}

func (s *imageService) GetCollectionByOwnerID(ctx context.Context, querier models.Querier, ownerID uuid.UUID) ([]*entities.CollectionEntity, error) {
	collectionEntity, err := s.imageRepo.GetCollectionByOwnerID(ctx, querier, ownerID)
	if err != nil {
		return nil, err
	}

	return collectionEntity, nil
}
func (s *imageService) CreateCollection(ctx context.Context, querier models.Querier, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	collectionEntity, err := s.imageRepo.CreateCollection(ctx, querier, collection)
	if err != nil {
		return nil, err
	}

	return collectionEntity, nil
}

func (s *imageService) AddImageToCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	collectionEntity, err := s.imageRepo.AddImageToCollection(ctx, querier, image, collection)
	if err != nil {
		return nil, err
	}

	return collectionEntity, nil
}
