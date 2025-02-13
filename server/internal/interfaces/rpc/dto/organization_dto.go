package dto

import (
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	organizationv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/organization/v1"
)

func OrganizationEntityToDto(organizationEntity *entities.OrganizationEntity) *organizationv1.Organization {

	organization := &organizationv1.Organization{
		Id:     organizationEntity.ID.String(),
		Name:   organizationEntity.Name,
		Handle: organizationEntity.Handle,
	}

	return organization
}
