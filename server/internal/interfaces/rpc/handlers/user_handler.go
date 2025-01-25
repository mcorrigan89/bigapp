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
	"github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/dto"
	commonv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/common/v1"
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

	var res *connect.Response[userv1.GetUserByIdResponse]

	userEntity, err := rpc.userApplicationService.GetUserByID(ctx, query)
	if err != nil {
		switch err {
		case entities.ErrUserNotFound:
			res = connect.NewResponse(&userv1.GetUserByIdResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_USER_NOT_FOUND,
					Message: "User not found",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	userDto := dto.UserEntityToDto(userEntity)

	res = connect.NewResponse(&userv1.GetUserByIdResponse{
		User: userDto,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) GetUserByHandle(ctx context.Context, req *connect.Request[userv1.GetUserByHandleRequest]) (*connect.Response[userv1.GetUserByHandleResponse], error) {
	handle := req.Msg.Handle
	if handle == "" {
		err := errors.New("handle is required")
		rpc.logger.Err(err).Ctx(ctx).Msg("Handle is required")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	query := queries.UserByHandleQuery{
		Handle: handle,
	}

	var res *connect.Response[userv1.GetUserByHandleResponse]

	userEntity, err := rpc.userApplicationService.GetUserByHandle(ctx, query)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error sending email login")
		switch err {
		case entities.ErrUserNotFound:
			res = connect.NewResponse(&userv1.GetUserByHandleResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_USER_NOT_FOUND,
					Message: "User not found",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	userDto := dto.UserEntityToDto(userEntity)

	res = connect.NewResponse(&userv1.GetUserByHandleResponse{
		User: userDto,
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

	userDto := dto.UserEntityToDto(userEntity)

	res := connect.NewResponse(&userv1.GetUserBySessionTokenResponse{
		User: userDto,
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

	userDto := dto.UserEntityToDto(userSessionEntity.User)

	res = connect.NewResponse(&userv1.CreateUserResponse{
		User: userDto,
		Session: &userv1.UserSession{
			Token:     userSessionEntity.SessionToken,
			ExpiresAt: userSessionEntity.ExpiresAt().Format(time.RFC1123Z),
		},
	})

	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (rpc *userServiceV1) UpdateUser(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.UpdateUserResponse], error) {
	Id := req.Msg.Id
	userUUID, err := uuid.Parse(Id)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error parsing user ID")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	email := req.Msg.Email
	givenName := req.Msg.GivenName
	familyName := req.Msg.FamilyName
	handle := req.Msg.Handle

	cmd := commands.UpdateUserCommand{
		ID:         userUUID,
		Email:      email,
		GivenName:  givenName,
		FamilyName: familyName,
		Handle:     handle,
	}

	var res *connect.Response[userv1.UpdateUserResponse]

	userEntity, err := rpc.userApplicationService.UpdateUser(ctx, cmd)
	if err != nil {
		rpc.logger.Err(err).Ctx(ctx).Msg("Error sending email login")
		switch err {
		case entities.ErrEmailInUse:
			res = connect.NewResponse(&userv1.UpdateUserResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_EMAIL_EXISTS,
					Message: "Email is not available",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		case entities.ErrHandleInUse:
			res = connect.NewResponse(&userv1.UpdateUserResponse{
				Error: &commonv1.ErrorDetails{
					Code:    commonv1.ErrorCode_ERROR_CODE_HANDLE_EXISTS,
					Message: "Handle is not available",
				},
			})
			res.Header().Set("User-Version", "v1")
			return res, nil
		default:
			return nil, connect.NewError(connect.CodeInternal, err)
		}
	}

	userDto := dto.UserEntityToDto(userEntity)

	res = connect.NewResponse(&userv1.UpdateUserResponse{
		User: userDto,
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

	userDto := dto.UserEntityToDto(userSessionEntity.User)

	res := connect.NewResponse(&userv1.LoginWithReferenceLinkResponse{
		User: userDto,
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
			FullName:   userSessionEntity.User.FullName(),
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
