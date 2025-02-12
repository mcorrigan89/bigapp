package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type OrganizationRepository interface {
	GetOrganizationByID(ctx context.Context, querier models.Querier, organizationID uuid.UUID) (*entities.OrganizationEntity, error)
	CreateOrganization(ctx context.Context, querier models.Querier, organization *entities.OrganizationEntity) (*entities.OrganizationEntity, error)
	CreateOrganizationRole(ctx context.Context, querier models.Querier, organizationRole *entities.OrganizationRoleEntity) (*entities.OrganizationRoleEntity, error)
	CreateOrganizationUser(ctx context.Context, querier models.Querier, organizationUser *entities.OrganizationUserEntity) (*entities.OrganizationUserEntity, error)
}
