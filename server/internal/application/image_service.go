package application

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/bigapp/server/internal/application/commands"
	"github.com/mcorrigan89/bigapp/server/internal/application/queries"
	"github.com/mcorrigan89/bigapp/server/internal/common"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/domain/external"
	"github.com/mcorrigan89/bigapp/server/internal/domain/services"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"

	"github.com/rs/zerolog"
)

type ImageApplicationService interface {
	GetImageByID(ctx context.Context, query queries.ImageByIDQuery) (*entities.ImageEntity, error)
	GetImageDataByID(ctx context.Context, query queries.ImageDataByIDQuery) ([]byte, string, error)
	UploadAvatarImage(ctx context.Context, cmd commands.CreateNewAvatarImageCommand) (*entities.ImageEntity, error)
	GetCollectionByID(ctx context.Context, query queries.CollectionByIDQuery) (*entities.CollectionEntity, error)
	GetCollectionByOwnerID(ctx context.Context, query queries.CollectionByOwnerIDQuery) ([]*entities.CollectionEntity, error)
	GetCollectionByOwnerToken(ctx context.Context, query queries.CollectionByOwnerTokenQuery) ([]*entities.CollectionEntity, error)
	CreateCollection(ctx context.Context, cmd commands.CreateNewCollectionCommand) (*entities.CollectionEntity, error)
	UploadImagesToCollection(ctx context.Context, cmd commands.UploadImagesToCollectionCommand) (*entities.CollectionEntity, error)
}

type imageApplicationService struct {
	config            *common.Config
	wg                *sync.WaitGroup
	logger            *zerolog.Logger
	db                *pgxpool.Pool
	queries           models.Querier
	imageService      services.ImageService
	userService       services.UserService
	imageMediaService external.ImageMediaService
}

func NewImageApplicationService(db *pgxpool.Pool, wg *sync.WaitGroup, cfg *common.Config, logger *zerolog.Logger, imageService services.ImageService, userService services.UserService, imageMediaService external.ImageMediaService) *imageApplicationService {
	dbQueries := models.New(db)
	return &imageApplicationService{
		db:                db,
		config:            cfg,
		wg:                wg,
		logger:            logger,
		queries:           dbQueries,
		imageService:      imageService,
		userService:       userService,
		imageMediaService: imageMediaService,
	}
}

func (app *imageApplicationService) GetImageByID(ctx context.Context, query queries.ImageByIDQuery) (*entities.ImageEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting image by ID")

	image, err := app.imageService.GetImageByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get image by ID")
		return nil, err
	}

	return image, nil
}

func (app *imageApplicationService) GetImageDataByID(ctx context.Context, query queries.ImageDataByIDQuery) ([]byte, string, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting image data by ID")

	imageEntity, err := app.imageService.GetImageByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get image by ID")
		return nil, "", err
	}

	imageData, contentType, err := app.imageMediaService.GetImageDataByID(ctx, imageEntity, query.Rendition)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get image data by ID")
		return nil, "", err
	}

	return imageData, contentType, nil
}

func (app *imageApplicationService) UploadAvatarImage(ctx context.Context, cmd commands.CreateNewAvatarImageCommand) (*entities.ImageEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Uploading image")
	tx, cancel, err := postgres.CreateTransaction(ctx, app.db)
	defer cancel()
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := models.New(app.db).WithTx(tx)

	app.logger.Info().Ctx(ctx).Msg("Uploading image data")
	imageData, err := app.imageMediaService.UploadImage(ctx, cmd.BucketID, cmd.ObjectID, cmd.File, cmd.Size)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to upload image")
		return nil, err
	}

	app.logger.Info().Ctx(ctx).Msg("Creating image in database")
	image, err := app.imageService.CreateImage(ctx, qtx, imageData)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create image")
		return nil, err
	}

	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")
	user, err := app.userService.GetUserByID(ctx, app.queries, cmd.UserID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	app.logger.Info().Ctx(ctx).Msg("Setting image as avatar on user")
	_, err = app.userService.SetAvatarImage(ctx, qtx, image, user)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to set avatar image")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to commit transaction")
		return nil, err
	}

	return image, nil
}

func (app *imageApplicationService) GetCollectionByID(ctx context.Context, query queries.CollectionByIDQuery) (*entities.CollectionEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting collection by ID")

	collectionEntity, err := app.imageService.GetCollectionByID(ctx, app.queries, query.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get collection by ID")
		return nil, err
	}

	return collectionEntity, nil
}

func (app *imageApplicationService) GetCollectionByOwnerID(ctx context.Context, query queries.CollectionByOwnerIDQuery) ([]*entities.CollectionEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting collection by owner ID")

	collectionEntities, err := app.imageService.GetCollectionByOwnerID(ctx, app.queries, query.OwnerID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get collection by owner ID")
		return nil, err
	}

	return collectionEntities, nil
}

func (app *imageApplicationService) GetCollectionByOwnerToken(ctx context.Context, query queries.CollectionByOwnerTokenQuery) ([]*entities.CollectionEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Getting collection by owner token")

	userContext, err := app.userService.GetUserContextBySessionToken(ctx, app.queries, query.Token)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by sessionToken")
		return nil, err
	}

	collectionEntities, err := app.imageService.GetCollectionByOwnerID(ctx, app.queries, userContext.User.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get collection by owner ID")
		return nil, err
	}

	return collectionEntities, nil
}

func (app *imageApplicationService) UploadImagesToCollection(ctx context.Context, cmd commands.UploadImagesToCollectionCommand) (*entities.CollectionEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Uploading images to collection")
	tx, cancel, err := postgres.CreateTransaction(ctx, app.db)
	defer cancel()
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := models.New(app.db).WithTx(tx)

	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")
	userEntity, err := app.userService.GetUserByID(ctx, qtx, cmd.UserID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	app.logger.Info().Ctx(ctx).Msg("Getting collection by ID")
	collectionEntity, err := app.imageService.GetCollectionByID(ctx, qtx, cmd.CollectionID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get collection by ID")
		return nil, err
	}

	if collectionEntity.OwnerID != userEntity.ID {
		app.logger.Err(err).Ctx(ctx).Msg("User does not own collection")
		return nil, entities.ErrCollectionNotAuthorized
	}

	type uploadResult struct {
		image *entities.ImageEntity
		err   error
		name  string
	}

	results := make(chan uploadResult, len(cmd.Files))
	var wg sync.WaitGroup

	for _, file := range cmd.Files {
		wg.Add(1)
		go func(file commands.ImageUploadData) {
			defer wg.Done()
			image, err := app.imageMediaService.UploadImage(ctx, cmd.BucketID, file.ObjectID, file.File, file.Size)
			results <- uploadResult{
				image: image,
				err:   err,
				name:  file.ObjectID,
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var successCount int
	errors := make([]string, 0)
	uploadedImages := make([]*entities.ImageEntity, 0)

	if len(errors) > 0 {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to upload images")
		return nil, fmt.Errorf("failed to upload images: %v", errors)
	}

	for result := range results {
		if result.err != nil {
			errors = append(errors, fmt.Sprintf("%s: %v", result.name, result.err))
			continue
		}
		successCount++
		uploadedImages = append(uploadedImages, result.image)
	}

	for _, image := range uploadedImages {
		uploadedImage, err := app.imageService.CreateImage(ctx, qtx, image)
		if err != nil {
			app.logger.Err(err).Ctx(ctx).Msg("Failed to create image")
			return nil, err
		}

		_, err = app.imageService.AddImageToCollection(ctx, qtx, uploadedImage, collectionEntity)
		if err != nil {
			app.logger.Err(err).Ctx(ctx).Msg("Failed to add image to collection")
			return nil, err
		}
	}

	collectionEntity, err = app.imageService.GetCollectionByID(ctx, qtx, collectionEntity.ID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get collection by ID")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to commit transaction")
		return nil, err
	}

	return collectionEntity, nil
}

func (app *imageApplicationService) CreateCollection(ctx context.Context, cmd commands.CreateNewCollectionCommand) (*entities.CollectionEntity, error) {
	app.logger.Info().Ctx(ctx).Msg("Uploading images to collection")
	tx, cancel, err := postgres.CreateTransaction(ctx, app.db)
	defer cancel()
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create transaction")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := models.New(app.db).WithTx(tx)

	collectionEntity := cmd.ToDomain()

	app.logger.Info().Ctx(ctx).Msg("Getting user by ID")
	_, err = app.userService.GetUserByID(ctx, qtx, collectionEntity.OwnerID)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to get user by ID")
		return nil, err
	}

	collectionEntity, err = app.imageService.CreateCollection(ctx, qtx, collectionEntity)
	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to create collection")
		return nil, err
	}

	err = tx.Commit(ctx)

	if err != nil {
		app.logger.Err(err).Ctx(ctx).Msg("Failed to commit transaction")
		return nil, err
	}

	return collectionEntity, nil
}
