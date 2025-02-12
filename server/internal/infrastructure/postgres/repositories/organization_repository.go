package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type postgresOrganizationRepository struct {
}

func NewPostgresOrganizationRepository() *postgresOrganizationRepository {
	return &postgresOrganizationRepository{}
}

func (repo *postgresOrganizationRepository) GetOrganizationByID(ctx context.Context, querier models.Querier, organizationID uuid.UUID) (*entities.OrganizationEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.GetOrganizationByID(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationEntity(row.Organization), nil
}

func (repo *postgresOrganizationRepository) CreateOrganization(ctx context.Context, querier models.Querier, organization *entities.OrganizationEntity) (*entities.OrganizationEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateOrganization(ctx, models.CreateOrganizationParams{
		ID:                 organization.ID,
		OrganizationName:   organization.Name,
		OrganizationHandle: organization.Handle,
	})
	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationEntity(row), nil
}

func (repo *postgresOrganizationRepository) CreateOrganizationRole(ctx context.Context, querier models.Querier, organizationRole *entities.OrganizationRoleEntity) (*entities.OrganizationRoleEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateOrganizationRole(ctx, models.CreateOrganizationRoleParams{
		ID:              organizationRole.ID,
		RoleName:        organizationRole.RoleName,
		RoleDescription: organizationRole.RoleDescription,
		RoleType:        organizationRole.RoleType,
	})
	if err != nil {
		return nil, err
	}

	return entities.NewOrganizationRoleEntity(row), nil
}

func (repo *postgresOrganizationRepository) CreateOrganizationUser(ctx context.Context, querier models.Querier, organizationUser *entities.OrganizationUserEntity) (*entities.OrganizationUserEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, postgres.DefaultTimeout)
	defer cancel()

	row, err := querier.CreateOrganizationUser(ctx, models.CreateOrganizationUserParams{
		ID:             organizationUser.ID,
		UserID:         organizationUser.UserID,
		RoleID:         organizationUser.RoleID,
		OrganizationID: organizationUser.OrganizationID,
	})
	if err != nil {
		return nil, err
	}

	return &entities.OrganizationUserEntity{
		ID:             row.ID,
		UserID:         row.UserID,
		RoleID:         row.RoleID,
		OrganizationID: row.OrganizationID,
	}, nil
}
