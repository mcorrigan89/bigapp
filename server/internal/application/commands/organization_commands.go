package commands

import (
	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/rs/xid"
)

type CreateNewOrganizationCommand struct {
	Name   string    `json:"name" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

func (cmd *CreateNewOrganizationCommand) ToDomain() (*entities.OrganizationEntity, uuid.UUID) {
	return &entities.OrganizationEntity{
		ID:     uuid.New(),
		Name:   cmd.Name,
		Handle: xid.New().String(),
	}, cmd.UserID
}
