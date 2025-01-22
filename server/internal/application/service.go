package application

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/services"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
	"github.com/rs/zerolog"
)

type UserApplicationService interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.UserEntity, error)
}

type userApplicationService struct {
	config      *common.Config
	wg          *sync.WaitGroup
	logger      *zerolog.Logger
	db          *pgxpool.Pool
	userService services.UserService
}

func NewUserApplicationService(db *pgxpool.Pool, wg *sync.WaitGroup, cfg *common.Config, logger *zerolog.Logger, userService services.UserService) *userApplicationService {
	return &userApplicationService{
		db:          db,
		config:      cfg,
		wg:          wg,
		logger:      logger,
		userService: userService,
	}
}

func (app *userApplicationService) GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.UserEntity, error) {
	queries := models.New(app.db)

	user, err := app.userService.GetUserByID(ctx, queries, userID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	return user, nil
}

func (app *userApplicationService) GetUserByEmail(ctx context.Context, email string) (*entities.UserEntity, error) {
	queries := models.New(app.db)

	user, err := app.userService.GetUserByEmail(ctx, queries, email)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by email")
		return nil, err
	}

	return user, nil
}
