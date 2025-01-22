package handlers

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"

	"github.com/google/uuid"

	"github.com/mcorrigan89/simple_auth/server/internal/application"
	"github.com/mcorrigan89/simple_auth/server/internal/application/commands"
	"github.com/mcorrigan89/simple_auth/server/internal/application/queries"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	userv1 "github.com/mcorrigan89/simple_auth/server/internal/interfaces/rpc/gen/user/v1"

	"github.com/rs/zerolog"
)

type userServiceV1 struct {
	logger                 *zerolog.Logger
	userApplicationService application.UserApplicationService
}

func NewUserServiceV1(logger *zerolog.Logger, userApplicationService application.UserApplicationService) *userServiceV1 {
	return &userServiceV1{
		logger:                 logger,
		userApplicationService: userApplicationService,
	}
}

func (rpc *userServiceV1) GetUserById(ctx context.Context, req *connect.Request[userv1.GetUserByIdRequest]) (*connect.Response[userv1.GetUserByIdResponse], error) {
	userID := req.Msg.Id

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error parsing user ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.UserByIDQuery{
		ID: userUUID,
	}

	userEntity, err := rpc.userApplicationService.GetUserByID(ctx, query)
	if err != nil {
		if err == entities.ErrUserNotFound {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		rpc.logger.Err(err).Ctx(ctx).Msg("Error getting user by ID")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&userv1.GetUserByIdResponse{
		User: &userv1.User{
			Id:         userEntity.ID.String(),
			GivenName:  userEntity.GivenName,
			FamilyName: userEntity.FamilyName,
			Email:      userEntity.Email,
		},
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) GetUserBySessionToken(ctx context.Context, req *connect.Request[userv1.GetUserBySessionTokenRequest]) (*connect.Response[userv1.GetUserBySessionTokenResponse], error) {
	token := req.Msg.Token
	if token == "" {
		err := errors.New("token is required")
		rpc.logger.Err(err).Ctx(ctx).Msg("Token is required")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.UserBySessionTokenQuery{
		SessionToken: token,
	}

	userEntity, err := rpc.userApplicationService.GetUserBySessionToken(ctx, query)
	if err != nil {
		if err == entities.ErrUserNotFound {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		rpc.logger.Err(err).Ctx(ctx).Msg("Error getting user by session token")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&userv1.GetUserBySessionTokenResponse{
		User: &userv1.User{
			Id:         userEntity.ID.String(),
			GivenName:  userEntity.GivenName,
			FamilyName: userEntity.FamilyName,
			Email:      userEntity.Email,
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) CreateUser(ctx context.Context, req *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.CreateUserResponse], error) {
	email := req.Msg.Email
	givenName := req.Msg.GivenName
	familyName := req.Msg.FamilyName

	cmd := commands.CreateNewUserCommand{
		Email:      email,
		GivenName:  givenName,
		FamilyName: familyName,
	}

	userSessionEntity, err := rpc.userApplicationService.CreateUser(ctx, cmd)
	if err != nil {
		if err == entities.ErrUserNotFound {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		rpc.logger.Err(err).Ctx(ctx).Msg("Error getting user by session token")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&userv1.CreateUserResponse{
		User: &userv1.User{
			Id:         userSessionEntity.User.ID.String(),
			GivenName:  userSessionEntity.User.GivenName,
			FamilyName: userSessionEntity.User.FamilyName,
			Email:      userSessionEntity.User.Email,
		},
		Session: &userv1.UserSession{
			Token:     userSessionEntity.SessionToken,
			ExpiresAt: userSessionEntity.ExpiresAt().Format(time.RFC1123Z),
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) CreateLoginEmail(ctx context.Context, req *connect.Request[userv1.CreateLoginEmailRequest]) (*connect.Response[userv1.CreateLoginEmailResponse], error) {
	email := req.Msg.Email

	cmd := commands.RequestEmailLoginCommand{
		Email: email,
	}

	err := rpc.userApplicationService.RequestEmailLogin(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error sending email login")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&userv1.CreateLoginEmailResponse{
		Status: "OK",
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) LoginWithReferenceLink(ctx context.Context, req *connect.Request[userv1.LoginWithReferenceLinkRequest]) (*connect.Response[userv1.LoginWithReferenceLinkResponse], error) {

	token := req.Msg.Token

	cmd := commands.LoginWithReferenceLinkCommand{
		ReferenceLinkToken: token,
	}

	userSessionEntity, err := rpc.userApplicationService.LoginWithReferenceLink(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error getting user by session token")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&userv1.LoginWithReferenceLinkResponse{
		User: &userv1.User{
			Id:         userSessionEntity.User.ID.String(),
			GivenName:  userSessionEntity.User.GivenName,
			FamilyName: userSessionEntity.User.FamilyName,
			Email:      userSessionEntity.User.Email,
		},
		Session: &userv1.UserSession{
			Token:     userSessionEntity.SessionToken,
			ExpiresAt: userSessionEntity.ExpiresAt().Format(time.RFC1123Z),
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}
