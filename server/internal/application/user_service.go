package application

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/bigapp/server/internal/application/commands"
	"github.com/mcorrigan89/bigapp/server/internal/application/queries"
	"github.com/mcorrigan89/bigapp/server/internal/common"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/domain/services"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"

	"github.com/rs/zerolog"
)

type UserApplicationService interface {
	GetUserByID(ctx context.Context, query queries.UserByIDQuery) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, query queries.UserByEmailQuery) (*entities.UserEntity, error)
	GetUserBySessionToken(ctx context.Context, query queries.UserBySessionTokenQuery) (*entities.UserEntity, error)
	CreateUser(ctx context.Context, cmd commands.CreateNewUserCommand) (*entities.UserContextEntity, error)
	RequestEmailLogin(ctx context.Context, cmd commands.RequestEmailLoginCommand) error
	LoginWithReferenceLink(ctx context.Context, cmd commands.LoginWithReferenceLinkCommand) (*entities.UserContextEntity, error)
}

type userApplicationService struct {
	config       *common.Config
	wg           *sync.WaitGroup
	logger       *zerolog.Logger
	db           *pgxpool.Pool
	queries      models.Querier
	userService  services.UserService
	emailService services.EmailService
}

func NewUserApplicationService(db *pgxpool.Pool, wg *sync.WaitGroup, cfg *common.Config, logger *zerolog.Logger, userService services.UserService, emailService services.EmailService) *userApplicationService {
	dbQueries := models.New(db)
	return &userApplicationService{
		db:           db,
		config:       cfg,
		wg:           wg,
		logger:       logger,
		queries:      dbQueries,
		userService:  userService,
		emailService: emailService,
	}
}

func (app *userApplicationService) GetUserByID(ctx context.Context, query queries.UserByIDQuery) (*entities.UserEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")

	user, err := app.userService.GetUserByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	return user, nil
}

func (app *userApplicationService) GetUserByEmail(ctx context.Context, query queries.UserByEmailQuery) (*entities.UserEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting user by email")

	user, err := app.userService.GetUserByEmail(ctx, app.queries, query.Email)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by email")
		return nil, err
	}

	return user, nil
}

func (app *userApplicationService) GetUserBySessionToken(ctx context.Context, query queries.UserBySessionTokenQuery) (*entities.UserEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting user by sessionToken")

	userContext, err := app.userService.GetUserContextBySessionToken(ctx, app.queries, query.SessionToken)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by sessionToken")
		return nil, err
	}

	return userContext.User, nil
}

func (app *userApplicationService) CreateUser(ctx context.Context, cmd commands.CreateNewUserCommand) (*entities.UserContextEntity, error) {
	tx, cancel, err := postgres.CreateTransaction(ctx, app.db)
	defer cancel()
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := models.New(app.db).WithTx(tx)

	app.logger.Info().Ctx(ctx).Msg("Creating new user")

	userEntity := cmd.ToDomain()

	createdUser, err := app.userService.CreateUser(ctx, qtx, userEntity)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create new user")
		return nil, err
	}

	userSession, err := app.userService.CreateSession(ctx, qtx, createdUser)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create new session")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to commit transaction")
		return nil, err
	}

	return userSession, nil
}

func (app *userApplicationService) RequestEmailLogin(ctx context.Context, cmd commands.RequestEmailLoginCommand) error {
	app.logger.Info().Ctx(ctx).Msg("Requesting email login")

	email := cmd.ToDomain()

	loginLink, err := app.userService.CreateLoginLink(ctx, app.queries, email)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create login link")
		return err
	}

	emailEntity := entities.EmailEntity{
		ID:        uuid.New(),
		ToEmail:   email,
		FromEmail: "mcorrigan89@gmail.com",
		Subject:   "Login to Simple Auth",
		Body:      fmt.Sprintf("Click the link to login: %s/authenticate?token=%s", app.config.CientURL, loginLink.Token),
	}

	_, err = app.emailService.SendEmail(ctx, app.queries, &emailEntity)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to send email")
		return err
	}

	return nil
}

func (app *userApplicationService) LoginWithReferenceLink(ctx context.Context, cmd commands.LoginWithReferenceLinkCommand) (*entities.UserContextEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Login with reference link")

	tx, cancel, err := postgres.CreateTransaction(ctx, app.db)
	defer cancel()
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := models.New(app.db).WithTx(tx)

	token := cmd.ToDomain()

	userSession, err := app.userService.LoginWithLink(ctx, qtx, token)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create login link")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to commit transaction")
		return nil, err
	}

	return userSession, nil
}
