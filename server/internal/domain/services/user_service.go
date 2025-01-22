package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/repositories"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
	"github.com/rs/xid"
)

type UserService interface {
	GetUserByID(ctx context.Context, querier models.Querier, userId uuid.UUID) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, querier models.Querier, email string) (*entities.UserEntity, error)
	GetUserContextBySessionToken(ctx context.Context, querier models.Querier, sessionToken string) (*entities.UserContextEntity, error)
	CreateUser(ctx context.Context, querier models.Querier, user *entities.UserEntity) (*entities.UserEntity, error)
	CreateSession(ctx context.Context, querier models.Querier, user *entities.UserEntity) (*entities.UserContextEntity, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserByID(ctx context.Context, querier models.Querier, userId uuid.UUID) (*entities.UserEntity, error) {
	user, err := s.userRepo.GetUserByID(ctx, querier, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, querier models.Querier, email string) (*entities.UserEntity, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, querier, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserContextBySessionToken(ctx context.Context, querier models.Querier, sessionToken string) (*entities.UserContextEntity, error) {
	user, err := s.userRepo.GetUserContextBySessionToken(ctx, querier, sessionToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, querier models.Querier, user *entities.UserEntity) (*entities.UserEntity, error) {
	user, err := s.userRepo.CreateUser(ctx, querier, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateSession(ctx context.Context, querier models.Querier, user *entities.UserEntity) (*entities.UserContextEntity, error) {

	token := xid.New().String()
	expiresAt := time.Now().Add(time.Hour * 24 * 30)

	userSession, err := s.userRepo.CreateSession(ctx, querier, user, token, expiresAt)
	if err != nil {
		return nil, err
	}

	return userSession, nil
}
