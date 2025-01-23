package application

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/simple_auth/server/internal/application/commands"
	"github.com/mcorrigan89/simple_auth/server/internal/application/queries"
	"github.com/mcorrigan89/simple_auth/server/internal/common"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/external"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/services"
	"github.com/mcorrigan89/simple_auth/server/internal/infrastructure/postgres/models"

	"github.com/rs/zerolog"
)

type ImageApplicationService interface {
	GetImageByID(ctx context.Context, query queries.ImageByIDQuery) (*entities.ImageEntity, error)
	GetImageDataByID(ctx context.Context, query queries.ImageDataByIDQuery) ([]byte, string, error)
	UploadImage(ctx context.Context, cmd commands.CreateNewImageCommand) (*entities.ImageEntity, error)
}

type imageApplicationService struct {
	config            *common.Config
	wg                *sync.WaitGroup
	logger            *zerolog.Logger
	db                *pgxpool.Pool
	queries           models.Querier
	imageService      services.ImageService
	imageMediaService external.ImageMediaService
}

func NewImageApplicationService(db *pgxpool.Pool, wg *sync.WaitGroup, cfg *common.Config, logger *zerolog.Logger, imageService services.ImageService, imageMediaService external.ImageMediaService) *imageApplicationService {
	dbQueries := models.New(db)
	return &imageApplicationService{
		db:                db,
		config:            cfg,
		wg:                wg,
		logger:            logger,
		queries:           dbQueries,
		imageService:      imageService,
		imageMediaService: imageMediaService,
	}
}

func (app *imageApplicationService) GetImageByID(ctx context.Context, query queries.ImageByIDQuery) (*entities.ImageEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")

	image, err := app.imageService.GetImageByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get image by ID")
		return nil, err
	}

	return image, nil
}

func (app *imageApplicationService) GetImageDataByID(ctx context.Context, query queries.ImageDataByIDQuery) ([]byte, string, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")

	imageEntity, err := app.imageService.GetImageByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get image by ID")
		return nil, "", err
	}

	imageData, contentType, err := app.imageMediaService.GetImageDataByID(ctx, imageEntity, "small")
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get image data by ID")
		return nil, "", err
	}

	return imageData, contentType, nil
}

func (app *imageApplicationService) UploadImage(ctx context.Context, cmd commands.CreateNewImageCommand) (*entities.ImageEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Uploading image")

	imageData, err := app.imageMediaService.UploadImage(ctx, "image", cmd.ObjectID, cmd.File, cmd.Size)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to upload image")
		return nil, err
	}

	image, err := app.imageService.CreateImage(ctx, app.queries, imageData)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create image")
		return nil, err
	}

	return image, nil
}
