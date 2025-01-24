package handlers

import (
	"context"
	"errors"
	"time"

	"connectrpc.com/connect"

	"github.com/google/uuid"

	"github.com/mcorrigan89/bigapp/server/internal/application"
	"github.com/mcorrigan89/bigapp/server/internal/application/commands"
	"github.com/mcorrigan89/bigapp/server/internal/application/queries"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	commonv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/common/v1"
	imagev1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/media/v1"
	userv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/user/v1"

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

	var avatar *imagev1.Image
	if userEntity.Avatar != nil {
		avatar = &imagev1.Image{
			Id:     userEntity.Avatar.ID.String(),
			Url:    userEntity.Avatar.UrlSlug(),
			Width:  userEntity.Avatar.Width,
			Height: userEntity.Avatar.Height,
			Size:   userEntity.Avatar.Size,
		}
	}

	res := connect.NewResponse(&userv1.GetUserByIdResponse{
		User: &userv1.User{
			Id:         userEntity.ID.String(),
			GivenName:  userEntity.GivenName,
			FamilyName: userEntity.FamilyName,
			Email:      userEntity.Email,
			Avatar:     avatar,
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

	var avatar *imagev1.Image
	if userEntity.Avatar != nil {
		avatar = &imagev1.Image{
			Id:     userEntity.Avatar.ID.String(),
			Url:    userEntity.Avatar.UrlSlug(),
			Width:  userEntity.Avatar.Width,
			Height: userEntity.Avatar.Height,
			Size:   userEntity.Avatar.Size,
		}
	}

	res := connect.NewResponse(&userv1.GetUserBySessionTokenResponse{
		User: &userv1.User{
			Id:         userEntity.ID.String(),
			GivenName:  userEntity.GivenName,
			FamilyName: userEntity.FamilyName,
			Email:      userEntity.Email,
			Avatar:     avatar,
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

	var res *connect.Response[userv1.CreateUserResponse]

	userSessionEntity, err := rpc.userApplicationService.CreateUser(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error sending email login")
		switch err {
		case entities.ErrEmailInUse:
			res = connect.NewResponse(&userv1.CreateUserResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_EMAIL_EXISTS,
					Message: "Email is not available",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	var avatar *imagev1.Image
	if userSessionEntity.User.Avatar != nil {
		avatar = &imagev1.Image{
			Id:     userSessionEntity.User.Avatar.ID.String(),
			Url:    userSessionEntity.User.Avatar.UrlSlug(),
			Width:  userSessionEntity.User.Avatar.Width,
			Height: userSessionEntity.User.Avatar.Height,
			Size:   userSessionEntity.User.Avatar.Size,
		}
	}

	res = connect.NewResponse(&userv1.CreateUserResponse{
		User: &userv1.User{
			Id:         userSessionEntity.User.ID.String(),
			GivenName:  userSessionEntity.User.GivenName,
			FamilyName: userSessionEntity.User.FamilyName,
			Email:      userSessionEntity.User.Email,
			Avatar:     avatar,
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

	var avatar *imagev1.Image
	if userSessionEntity.User.Avatar != nil {
		avatar = &imagev1.Image{
			Id:     userSessionEntity.User.Avatar.ID.String(),
			Url:    userSessionEntity.User.Avatar.UrlSlug(),
			Width:  userSessionEntity.User.Avatar.Width,
			Height: userSessionEntity.User.Avatar.Height,
			Size:   userSessionEntity.User.Avatar.Size,
		}
	}

	res := connect.NewResponse(&userv1.LoginWithReferenceLinkResponse{
		User: &userv1.User{
			Id:         userSessionEntity.User.ID.String(),
			GivenName:  userSessionEntity.User.GivenName,
			FamilyName: userSessionEntity.User.FamilyName,
			Email:      userSessionEntity.User.Email,
			Avatar:     avatar,
		},
		Session: &userv1.UserSession{
			Token:     userSessionEntity.SessionToken,
			ExpiresAt: userSessionEntity.ExpiresAt().Format(time.RFC1123Z),
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) InviteUser(ctx context.Context, req *connect.Request[userv1.InviteUserRequest]) (*connect.Response[userv1.InviteUserResponse], error) {
	email := req.Msg.Email

	cmd := commands.InviteUserCommand{
		Email: email,
	}

	var res *connect.Response[userv1.InviteUserResponse]

	err := rpc.userApplicationService.InviteUser(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error sending email login")
		switch err {
		case entities.ErrEmailInUse:
			res = connect.NewResponse(&userv1.InviteUserResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_EMAIL_EXISTS,
					Message: "Email is not available",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		case entities.ErrUserClaimed:
			res = connect.NewResponse(&userv1.InviteUserResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_EMAIL_EXISTS,
					Message: "User is already claimed",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	res = connect.NewResponse(&userv1.InviteUserResponse{
		Status: "OK",
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) AcceptInviteReferenceLink(ctx context.Context, req *connect.Request[userv1.AcceptInviteReferenceLinkRequest]) (*connect.Response[userv1.AcceptInviteReferenceLinkResponse], error) {
	token := req.Msg.Token

	cmd := commands.AcceptInviteReferenceLinkCommand{
		ReferenceLinkToken: token,
	}

	userSessionEntity, err := rpc.userApplicationService.AcceptInviteReferenceLink(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error accepting email invite")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&userv1.AcceptInviteReferenceLinkResponse{
		User: &userv1.User{
			Id:         userSessionEntity.User.ID.String(),
			GivenName:  userSessionEntity.User.GivenName,
			FamilyName: userSessionEntity.User.FamilyName,
			Email:      userSessionEntity.User.Email,
			Avatar:     nil,
		},
		Session: &userv1.UserSession{
			Token:     userSessionEntity.SessionToken,
			ExpiresAt: userSessionEntity.ExpiresAt().Format(time.RFC1123Z),
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}
