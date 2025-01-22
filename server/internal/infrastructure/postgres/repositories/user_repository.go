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

func (repo *postgresUserRepository) GetUserByID(ctx context.Context, querier models.Querier, userId uuid.UUID) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row.User), nil
}

func (repo *postgresUserRepository) GetUserByEmail(ctx context.Context, querier models.Querier, email string) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row.User), nil
}

func (repo *postgresUserRepository) GetUserContextBySessionToken(ctx context.Context, querier models.Querier, sessionToken string) (*entities.UserContextEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserBySessionToken(ctx, sessionToken)
	if err != nil {
		return nil, err
	}

	return entities.NewUserContextEntity(row.User, row.UserSession), nil
}

func (repo *postgresUserRepository) CreateUser(ctx context.Context, querier models.Querier, user *entities.UserEntity) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateUser(ctx, models.CreateUserParams{
		ID:            user.ID,
		GivenName:     user.GivenName,
		FamilyName:    user.FamilyName,
		Email:         user.Email,
		EmailVerified: false,
		AvatarUrl:     nil,
	})
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row), nil
}
