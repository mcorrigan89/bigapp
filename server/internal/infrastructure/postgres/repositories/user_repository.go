package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
)

type postgresUserRepository struct {
}

func NewPostgresUserRepository() *postgresUserRepository {
	return &postgresUserRepository{}
}

func (repo *postgresUserRepository) GetUserByID(ctx context.Context, querier *models.Queries, userId uuid.UUID) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row.User), nil
}

func (repo *postgresUserRepository) GetUserByEmail(ctx context.Context, querier *models.Queries, email string) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row.User), nil
}
