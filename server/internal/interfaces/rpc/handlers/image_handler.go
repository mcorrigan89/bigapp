package handlers

import (
	"context"
	"errors"

	"connectrpc.com/connect"

	"github.com/google/uuid"

	"github.com/mcorrigan89/bigapp/server/internal/application"

	"github.com/mcorrigan89/bigapp/server/internal/application/commands"
	"github.com/mcorrigan89/bigapp/server/internal/application/queries"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/http/middleware"
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/dto"
	commonv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/common/v1"
	mediav1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/media/v1"

	"github.com/rs/zerolog"
)

type imageServiceV1 struct {
	logger                  *zerolog.Logger
	imageApplicationService application.ImageApplicationService
}

func NewImageServiceV1(logger *zerolog.Logger, imageApplicationService application.ImageApplicationService) *imageServiceV1 {
	return &imageServiceV1{
		logger:                  logger,
		imageApplicationService: imageApplicationService,
	}
}

func (rpc *imageServiceV1) GetImageById(ctx context.Context, req *connect.Request[mediav1.GetImageByIdRequest]) (*connect.Response[mediav1.GetImageByIdResponse], error) {
	imageID := req.Msg.ImageId

	imageUUID, err := uuid.Parse(imageID)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error parsing image ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.ImageByIDQuery{
		ID: imageUUID,
	}

	var res *connect.Response[mediav1.GetImageByIdResponse]

	imageEntity, err := rpc.imageApplicationService.GetImageByID(ctx, query)
	if err != nil {
		switch err {
		case entities.ErrUserNotFound:
			res = connect.NewResponse(&mediav1.GetImageByIdResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_IMAGE_NOT_FOUND,
					Message: "Image not found",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	imageDto := dto.ImageEntityToDto(imageEntity)

	res = connect.NewResponse(&mediav1.GetImageByIdResponse{
		Image: imageDto,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *imageServiceV1) GetCollectionById(ctx context.Context, req *connect.Request[mediav1.GetCollectionByIdRequest]) (*connect.Response[mediav1.GetCollectionByIdResponse], error) {
	collectionID := req.Msg.CollectionId

	collectionUUID, err := uuid.Parse(collectionID)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error parsing collection ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.CollectionByIDQuery{
		ID: collectionUUID,
	}

	var res *connect.Response[mediav1.GetCollectionByIdResponse]

	collectionEntity, err := rpc.imageApplicationService.GetCollectionByID(ctx, query)
	if err != nil {
		switch err {
		case entities.ErrCollectionNotFound:
			res = connect.NewResponse(&mediav1.GetCollectionByIdResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_IMAGE_NOT_FOUND,
					Message: "Collection not found",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	collectionDto := dto.CollectionEntityToDto(collectionEntity)

	res = connect.NewResponse(&mediav1.GetCollectionByIdResponse{
		Collection: collectionDto,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *imageServiceV1) GetCollectionByOwnerId(ctx context.Context, req *connect.Request[mediav1.GetCollectionByOwnerIdRequest]) (*connect.Response[mediav1.GetCollectionByOwnerIdResponse], error) {
	ownerID := req.Msg.OwnerId

	ownerUUID, err := uuid.Parse(ownerID)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error parsing owner ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.CollectionByOwnerIDQuery{
		OwnerID: ownerUUID,
	}

	var res *connect.Response[mediav1.GetCollectionByOwnerIdResponse]

	collectionEntities, err := rpc.imageApplicationService.GetCollectionByOwnerID(ctx, query)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error getting collection by owner ID")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	collectionDtos := make([]*mediav1.Collection, 0, len(collectionEntities))
	for _, collectionEntity := range collectionEntities {
		collectionDto := dto.CollectionEntityToDto(collectionEntity)
		collectionDtos = append(collectionDtos, collectionDto)
	}

	res = connect.NewResponse(&mediav1.GetCollectionByOwnerIdResponse{
		Collections: collectionDtos,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *imageServiceV1) GetCollectionByOwnerToken(ctx context.Context, req *connect.Request[mediav1.GetCollectionByOwnerTokenRequest]) (*connect.Response[mediav1.GetCollectionByOwnerTokenResponse], error) {
	token := req.Msg.Token

	query := queries.CollectionByOwnerTokenQuery{
		Token: token,
	}

	var res *connect.Response[mediav1.GetCollectionByOwnerTokenResponse]

	collectionEntities, err := rpc.imageApplicationService.GetCollectionByOwnerToken(ctx, query)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error getting collection by owner ID")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	collectionDtos := make([]*mediav1.Collection, 0, len(collectionEntities))
	for _, collectionEntity := range collectionEntities {
		collectionDto := dto.CollectionEntityToDto(collectionEntity)
		collectionDtos = append(collectionDtos, collectionDto)
	}

	res = connect.NewResponse(&mediav1.GetCollectionByOwnerTokenResponse{
		Collections: collectionDtos,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *imageServiceV1) CreateCollection(ctx context.Context, req *connect.Request[mediav1.CreateCollectionRequest]) (*connect.Response[mediav1.CreateCollectionResponse], error) {
	collectionName := req.Msg.CollectionName

	userContextEntity := middleware.GetUserFromContext(ctx)
	if userContextEntity == nil {
		err := errors.New("user is not authenticated")
		rpc.logger.Err(err).Ctx(ctx).Msg("User is not authenticated")
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	cmd := commands.CreateNewCollectionCommand{
		OwnerID: userContextEntity.User.ID,
		Name:    collectionName,
	}

	var res *connect.Response[mediav1.CreateCollectionResponse]

	collectionEntity, err := rpc.imageApplicationService.CreateCollection(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error creating collection")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res = connect.NewResponse(&mediav1.CreateCollectionResponse{
		Collection: dto.CollectionEntityToDto(collectionEntity),
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}
