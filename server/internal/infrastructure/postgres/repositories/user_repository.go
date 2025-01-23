package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type postgresUserRepository struct {
}

func NewPostgresUserRepository() *postgresUserRepository {
	return &postgresUserRepository{}
}

func (repo *postgresUserRepository) GetUserByID(ctx context.Context, querier models.Querier, userID uuid.UUID) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	avatarImageEntity, err := repo.getAvatarImageEntity(ctx, querier, &row.User)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row.User, avatarImageEntity), nil
}

func (repo *postgresUserRepository) GetUserByEmail(ctx context.Context, querier models.Querier, email string) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	avatarImageEntity, err := repo.getAvatarImageEntity(ctx, querier, &row.User)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row.User, avatarImageEntity), nil
}

func (repo *postgresUserRepository) GetUserContextBySessionToken(ctx context.Context, querier models.Querier, sessionToken string) (*entities.UserContextEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetUserBySessionToken(ctx, sessionToken)
	if err != nil {
		return nil, err
	}

	avatarImageEntity, err := repo.getAvatarImageEntity(ctx, querier, &row.User)
	if err != nil {
		return nil, err
	}

	userEntity := entities.NewUserEntity(row.User, avatarImageEntity)

	return entities.NewUserContextEntity(userEntity, row.UserSession), nil
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
	})
	if err != nil {
		return nil, err
	}

	avatarImageEntity, err := repo.getAvatarImageEntity(ctx, querier, &row)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row, avatarImageEntity), nil
}

func (repo *postgresUserRepository) CreateSession(ctx context.Context, querier models.Querier, user *entities.UserEntity, token string, expiresAt time.Time) (*entities.UserContextEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateUserSession(ctx, models.CreateUserSessionParams{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return nil, err
	}

	return entities.NewUserContextEntity(user, row), nil
}

func (repo *postgresUserRepository) SetAvatarImage(ctx context.Context, querier models.Querier, image *entities.ImageEntity, user *entities.UserEntity) (*entities.UserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.SetAvatarImage(ctx, models.SetAvatarImageParams{
		UserID:  user.ID,
		ImageID: &image.ID,
	})
	if err != nil {
		return nil, err
	}

	avatarImageEntity, err := repo.getAvatarImageEntity(ctx, querier, &row)
	if err != nil {
		return nil, err
	}

	return entities.NewUserEntity(row, avatarImageEntity), nil
}

func (repo *postgresUserRepository) getAvatarImageEntity(ctx context.Context, querier models.Querier, user *models.User) (*entities.ImageEntity, error) {
	if user.AvatarID != nil {
		avatarRow, err := querier.GetImageByID(ctx, *user.AvatarID)
		if err != nil {
			return nil, err
		}

		return entities.NewImageEntity(avatarRow.Image), nil
	}

	return nil, nil
}
