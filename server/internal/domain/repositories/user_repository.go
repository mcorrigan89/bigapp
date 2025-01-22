package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, querier *models.Queries, userId uuid.UUID) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, querier *models.Queries, email string) (*entities.UserEntity, error)
}
