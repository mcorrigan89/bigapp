package application

import (
	"context"
	"sync"

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

type OrganizationApplicationService interface {
	GetOrganizationByID(ctx context.Context, query queries.OrganizationByIDQuery) (*entities.OrganizationEntity, error)
	CreateOrganization(ctx context.Context, command commands.CreateNewOrganizationCommand) (*entities.OrganizationEntity, error)
}

type organizationModelApplicationService struct {
	config              *common.Config
	wg                  *sync.WaitGroup
	logger              *zerolog.Logger
	db                  *pgxpool.Pool
	queries             models.Querier
	organizationService services.OrganizationService
	userService         services.UserService
}

func NewOrganizationApplicationService(db *pgxpool.Pool, wg *sync.WaitGroup, cfg *common.Config, logger *zerolog.Logger, organizationService services.OrganizationService, userService services.UserService) *organizationModelApplicationService {
	dbQueries := models.New(db)
	return &organizationModelApplicationService{
		db:                  db,
		config:              cfg,
		wg:                  wg,
		logger:              logger,
		queries:             dbQueries,
		organizationService: organizationService,
		userService:         userService,
	}
}

func (app *organizationModelApplicationService) GetOrganizationByID(ctx context.Context, query queries.OrganizationByIDQuery) (*entities.OrganizationEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting organization by ID")

	user, err := app.organizationService.GetOrganizationByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get organization by ID")
		return nil, err
	}

	return user, nil
}

func (app *organizationModelApplicationService) CreateOrganization(ctx context.Context, command commands.CreateNewOrganizationCommand) (*entities.OrganizationEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Creating organization")
	tx, cancel, err := postgres.CreateTransaction(ctx, app.db)
	defer cancel()
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := models.New(app.db).WithTx(tx)

	organizationArgs, userID := command.ToDomain()

	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")
	userEntity, err := app.userService.GetUserByID(ctx, qtx, userID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	app.logger.Info().Ctx(ctx).Msg("Creating organization")
	organizationEntity, err := app.organizationService.CreateOrganization(ctx, app.queries, organizationArgs, userEntity)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create organization")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to commit transaction")
		return nil, err
	}

	return organizationEntity, nil
}
