package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	"github.com/mcorrigan89/bigapp/server/internal/domain/repositories"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
)

type OrganizationService interface {
	GetOrganizationByID(ctx context.Context, querier models.Querier, organizationID uuid.UUID) (*entities.OrganizationEntity, error)
	CreateOrganization(ctx context.Context, querier models.Querier, organization *entities.OrganizationEntity, user *entities.UserEntity) (*entities.OrganizationEntity, error)
}

type organizationService struct {
	orgRepo repositories.OrganizationRepository
}

func NewOrganizationService(orgRepo repositories.OrganizationRepository) *organizationService {
	return &organizationService{orgRepo: orgRepo}
}

func (s *organizationService) GetOrganizationByID(ctx context.Context, querier models.Querier, organizationID uuid.UUID) (*entities.OrganizationEntity, error) {
	organization, err := s.orgRepo.GetOrganizationByID(ctx, querier, organizationID)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (s *organizationService) CreateOrganization(ctx context.Context, querier models.Querier, organization *entities.OrganizationEntity, userEntity *entities.UserEntity) (*entities.OrganizationEntity, error) {
	organization, err := s.orgRepo.CreateOrganization(ctx, querier, organization)
	if err != nil {
		return nil, err
	}

	ownerRoleDescription := "Owner of the organization"
	ownerRole, err := s.orgRepo.CreateOrganizationRole(ctx, querier, &entities.OrganizationRoleEntity{
		ID:              uuid.New(),
		RoleName:        "Owner",
		RoleDescription: &ownerRoleDescription,
		RoleType:        entities.OrganizationOwnerRole,
	})
	if err != nil {
		return nil, err
	}

	adminRoleDescription := "Admin in the organization"
	_, err = s.orgRepo.CreateOrganizationRole(ctx, querier, &entities.OrganizationRoleEntity{
		ID:              uuid.New(),
		RoleName:        "Admin",
		RoleDescription: &adminRoleDescription,
		RoleType:        entities.OrganizationAdminRole,
	})
	if err != nil {
		return nil, err
	}

	memberRoleDescription := "Member in the organization"
	_, err = s.orgRepo.CreateOrganizationRole(ctx, querier, &entities.OrganizationRoleEntity{
		ID:              uuid.New(),
		RoleName:        "Member",
		RoleDescription: &memberRoleDescription,
		RoleType:        entities.OrganizationMemberRole,
	})

	if err != nil {
		return nil, err
	}

	_, err = s.orgRepo.CreateOrganizationUser(ctx, querier, &entities.OrganizationUserEntity{
		ID:             uuid.New(),
		UserID:         userEntity.ID,
		RoleID:         ownerRole.ID,
		OrganizationID: organization.ID,
	})
	if err != nil {
		return nil, err
	}

	return organization, nil
}
