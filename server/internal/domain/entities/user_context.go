package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/models"
)

type UserContextEntity struct {
	SessionToken   string
	UserID         uuid.UUID
	User           *UserEntity
	userExpired    bool
	expiresAt      time.Time
	impersonatorID *uuid.UUID
}

func NewUserContextEntity(userModel models.User, userSessionModel models.UserSession) *UserContextEntity {
	return &UserContextEntity{
		SessionToken:   userSessionModel.Token,
		UserID:         userModel.ID,
		User:           NewUserEntity(userModel),
		userExpired:    userSessionModel.UserExpired,
		expiresAt:      userSessionModel.ExpiresAt,
		impersonatorID: userSessionModel.ImpersonatorID,
	}
}

func (uc *UserContextEntity) IsExpired() bool {
	return uc.userExpired || uc.expiresAt.Before(time.Now())
}
