package entities

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

var (
	ErrOrganizationNotFound      = errors.New("organization not found")
	ErrorOrganizationHandleInUse = errors.New("organization handle in use")
)

type OrganizationEntity struct {
	ID     uuid.UUID
	Name   string
	Handle string
}

func NewOrganizationEntity(organizationModel models.Organization) *OrganizationEntity {
	return &OrganizationEntity{
		ID:     organizationModel.ID,
		Name:   organizationModel.OrganizationName,
		Handle: organizationModel.OrganizationHandle,
	}
}

type OrganizationUserEntity struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	RoleID         uuid.UUID
	OrganizationID uuid.UUID
}

var (
	OrganizationOwnerRole  = "owner"
	OrganizationAdminRole  = "admin"
	OrganizationMemberRole = "member"
)

type OrganizationRoleEntity struct {
	ID              uuid.UUID
	RoleName        string
	RoleDescription *string
	RoleType        string
}

func NewOrganizationRoleEntity(organizationRoleModel models.OrganizationRole) *OrganizationRoleEntity {
	return &OrganizationRoleEntity{
		ID:              organizationRoleModel.ID,
		RoleName:        organizationRoleModel.RoleName,
		RoleDescription: organizationRoleModel.RoleDescription,
		RoleType:        organizationRoleModel.RoleType,
	}
}
