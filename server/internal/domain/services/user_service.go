package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/repositories"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
)

type UserService interface {
	GetUserByID(ctx context.Context, querier *models.Queries, userId uuid.UUID) (*entities.UserEntity, error)
	GetUserByEmail(ctx context.Context, querier *models.Queries, email string) (*entities.UserEntity, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserByID(ctx context.Context, querier *models.Queries, userId uuid.UUID) (*entities.UserEntity, error) {
	user, err := s.userRepo.GetUserByID(ctx, querier, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, querier *models.Queries, email string) (*entities.UserEntity, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, querier, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
