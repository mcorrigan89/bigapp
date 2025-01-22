package application

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/simple_auth/server/internal/application/commands"
	"github.com/mcorrigan89/simple_auth/server/internal/application/queries"
	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/services"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
	"github.com/rs/zerolog"
)

type UserApplicationService interface {
	GetUserByID(ctx context.Context, query queries.UserByIDQuery) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, query queries.UserByEmailQuery) (*entities.UserEntity, error)
	GetUserBySessionToken(ctx context.Context, query queries.UserBySessionTokenQuery) (*entities.UserEntity, error)
	CreateUser(ctx context.Context, cmd commands.CreateNewUserCommand) (*entities.UserEntity, error)
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

func (app *userApplicationService) GetUserByID(ctx context.Context, query queries.UserByIDQuery) (*entities.UserEntity, error) {
	queries := models.New(app.db)

	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")

	user, err := app.userService.GetUserByID(ctx, queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	return user, nil
}

func (app *userApplicationService) GetUserByEmail(ctx context.Context, query queries.UserByEmailQuery) (*entities.UserEntity, error) {
	queries := models.New(app.db)

	app.logger.Info().Ctx(ctx).Msg("Getting user by email")

	user, err := app.userService.GetUserByEmail(ctx, queries, query.Email)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by email")
		return nil, err
	}

	return user, nil
}

func (app *userApplicationService) GetUserBySessionToken(ctx context.Context, query queries.UserBySessionTokenQuery) (*entities.UserEntity, error) {
	queries := models.New(app.db)

	app.logger.Info().Ctx(ctx).Msg("Getting user by sessionToken")

	userContext, err := app.userService.GetUserContextBySessionToken(ctx, queries, query.SessionToken)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by sessionToken")
		return nil, err
	}

	return userContext.User, nil
}

func (app *userApplicationService) CreateUser(ctx context.Context, cmd commands.CreateNewUserCommand) (*entities.UserEntity, error) {
	queries := models.New(app.db)

	app.logger.Info().Ctx(ctx).Msg("Creating new user")

	userEntity := cmd.ToDomain()

	createdUser, err := app.userService.CreateUser(ctx, queries, userEntity)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create new user")
		return nil, err
	}

	return createdUser, nil
}
