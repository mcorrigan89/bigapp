package commands

import (
	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
)

type CreateNewUserCommand struct {
	Email      string  `json:"email" validate:"required,email"`
	GivenName  *string `json:"firstName" validate:"-"`
	FamilyName *string `json:"lastName" validate:"-"`
}

func (cmd *CreateNewUserCommand) ToDomain() *entities.UserEntity {
	return &entities.UserEntity{
		ID:         uuid.New(),
		Email:      cmd.Email,
		GivenName:  cmd.GivenName,
		FamilyName: cmd.FamilyName,
	}
}

type RequestEmailLoginCommand struct {
	Email string `json:"email" validate:"required,email"`
}

func (cmd *RequestEmailLoginCommand) ToDomain() string {
	return cmd.Email
}

type LoginWithReferenceLinkCommand struct {
	ReferenceLinkToken string `json:"token" validate:"required"`
}

func (cmd *LoginWithReferenceLinkCommand) ToDomain() string {
	return cmd.ReferenceLinkToken
}
