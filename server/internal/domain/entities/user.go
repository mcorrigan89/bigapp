package entities

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserEntity struct {
	ID         uuid.UUID
	GivenName  *string
	FamilyName *string
	Email      string
	Avatar     *ImageEntity
}

func NewUserEntity(userModel models.User, imageEntity *ImageEntity) *UserEntity {
	return &UserEntity{
		ID:         userModel.ID,
		GivenName:  userModel.GivenName,
		FamilyName: userModel.FamilyName,
		Email:      userModel.Email,
		Avatar:     imageEntity,
	}
}
