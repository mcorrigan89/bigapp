package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, querier models.Querier, userId uuid.UUID) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, querier models.Querier, email string) (*entities.UserEntity, error)
	GetUserContextBySessionToken(ctx context.Context, querier models.Querier, sessionToken string) (*entities.UserContextEntity, error)
	CreateUser(ctx context.Context, querier models.Querier, user *entities.UserEntity) (*entities.UserEntity, error)
}
