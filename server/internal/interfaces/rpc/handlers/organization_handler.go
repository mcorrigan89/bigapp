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
	organizationv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/organization/v1"

	"github.com/rs/zerolog"
)

type organizationServiceV1 struct {
	logger                         *zerolog.Logger
	organizationApplicationService application.OrganizationApplicationService
}

func NewOrganizationServiceV1(logger *zerolog.Logger, organizationApplicationService application.OrganizationApplicationService) *organizationServiceV1 {
	return &organizationServiceV1{
		logger:                         logger,
		organizationApplicationService: organizationApplicationService,
	}
}

func (rpc *organizationServiceV1) GetOrganizationById(ctx context.Context, req *connect.Request[organizationv1.GetOrganizationByIdRequest]) (*connect.Response[organizationv1.GetOrganizationByIdResponse], error) {
	organiationID := req.Msg.Id

	organizationUUID, err := uuid.Parse(organiationID)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error parsing organization ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.OrganizationByIDQuery{
		ID: organizationUUID,
	}

	var res *connect.Response[organizationv1.GetOrganizationByIdResponse]

	organizationEntity, err := rpc.organizationApplicationService.GetOrganizationByID(ctx, query)
	if err != nil {
		switch err {
		case entities.ErrOrganizationNotFound:
			res = connect.NewResponse(&organizationv1.GetOrganizationByIdResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_ORGANIZATION_NOT_FOUND,
					Message: "Organization not found",
				},
			})
			res.Header().Set("Organization-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	userDto := dto.OrganizationEntityToDto(organizationEntity)

	res = connect.NewResponse(&organizationv1.GetOrganizationByIdResponse{
		Organization: userDto,
	})
	res.Header().Set("Organization-Version", "v1")
	return res, nil
}

func (rpc *organizationServiceV1) CreateOrganization(ctx context.Context, req *connect.Request[organizationv1.CreateOrganizationRequest]) (*connect.Response[organizationv1.CreateOrganizationResponse], error) {

	userContextEntity := middleware.GetUserFromContext(ctx)
	if userContextEntity == nil {
		err := errors.New("user is not authenticated")
		rpc.logger.Err(err).Ctx(ctx).Msg("User is not authenticated")
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	command := commands.CreateNewOrganizationCommand{
		Name:   req.Msg.Name,
		UserID: userContextEntity.UserID,
	}

	var res *connect.Response[organizationv1.CreateOrganizationResponse]

	organizationEntity, err := rpc.organizationApplicationService.CreateOrganization(ctx, command)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	organizationDto := dto.OrganizationEntityToDto(organizationEntity)

	res = connect.NewResponse(&organizationv1.CreateOrganizationResponse{
		Organization: organizationDto,
	})
	res.Header().Set("Organization-Version", "v1")
	return res, nil
}
