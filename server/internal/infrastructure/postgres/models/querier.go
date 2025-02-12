// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddImageToCollection(ctx context.Context, arg AddImageToCollectionParams) (CollectionImage, error)
	CreateCollection(ctx context.Context, arg CreateCollectionParams) (Collection, error)
	CreateImage(ctx context.Context, arg CreateImageParams) (Image, error)
	CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error)
	CreateOrganizationRole(ctx context.Context, arg CreateOrganizationRoleParams) (OrganizationRole, error)
	CreateOrganizationUser(ctx context.Context, arg CreateOrganizationUserParams) (OrganizationUser, error)
	CreateReferenceLink(ctx context.Context, arg CreateReferenceLinkParams) (ReferenceLink, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserSession(ctx context.Context, arg CreateUserSessionParams) (UserSession, error)
	DeleteReferenceLink(ctx context.Context, id uuid.UUID) (ReferenceLink, error)
	ExpireUserSession(ctx context.Context, id uuid.UUID) error
	GetCollectionByID(ctx context.Context, id uuid.UUID) (GetCollectionByIDRow, error)
	GetCollectionByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]GetCollectionByOwnerIDRow, error)
	GetCollectionImagesByCollectionID(ctx context.Context, collectionID uuid.UUID) ([]GetCollectionImagesByCollectionIDRow, error)
	GetImageByID(ctx context.Context, id uuid.UUID) (GetImageByIDRow, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (GetOrganizationByIDRow, error)
	GetReferenceLinkByID(ctx context.Context, id uuid.UUID) (GetReferenceLinkByIDRow, error)
	GetReferenceLinkByToken(ctx context.Context, token string) (GetReferenceLinkByTokenRow, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
	GetUserByHandle(ctx context.Context, userHandle string) (GetUserByHandleRow, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIDRow, error)
	GetUserBySessionToken(ctx context.Context, token string) (GetUserBySessionTokenRow, error)
	RemoveImageFromCollection(ctx context.Context, arg RemoveImageFromCollectionParams) error
	SetAvatarImage(ctx context.Context, arg SetAvatarImageParams) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
