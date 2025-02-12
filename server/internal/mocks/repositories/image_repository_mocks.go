package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
	"github.com/stretchr/testify/mock"
)

type MockImageRepository struct {
	mock.Mock
}

func (m *MockImageRepository) GetImageByID(ctx context.Context, querier models.Querier, imageID uuid.UUID) (*entities.ImageEntity, error) {
	args := m.Called(ctx, querier, imageID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ImageEntity), args.Error(1)
}

func (m *MockImageRepository) CreateImage(ctx context.Context, querier models.Querier, image *entities.ImageEntity) (*entities.ImageEntity, error) {
	args := m.Called(ctx, querier, image)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ImageEntity), args.Error(1)
}

func (m *MockImageRepository) GetCollectionByID(ctx context.Context, querier models.Querier, id uuid.UUID) (*entities.CollectionEntity, error) {
	args := m.Called(ctx, querier, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.CollectionEntity), args.Error(1)
}

func (m *MockImageRepository) GetCollectionByOwnerID(ctx context.Context, querier models.Querier, ownerID uuid.UUID) ([]*entities.CollectionEntity, error) {
	args := m.Called(ctx, querier, ownerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.CollectionEntity), args.Error(1)
}

func (m *MockImageRepository) CreateCollection(ctx context.Context, querier models.Querier, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	args := m.Called(ctx, querier, collection)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.CollectionEntity), args.Error(1)
}

func (m *MockImageRepository) AddImageToCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	args := m.Called(ctx, querier, image, collection)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.CollectionEntity), args.Error(1)
}

func (m *MockImageRepository) RemoveImageFromCollection(ctx context.Context, querier models.Querier, image *entities.ImageEntity, collection *entities.CollectionEntity) (*entities.CollectionEntity, error) {
	args := m.Called(ctx, querier, image, collection)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.CollectionEntity), args.Error(1)
}
