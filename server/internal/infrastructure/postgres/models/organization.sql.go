// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: organization.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO organizations (id, organization_name, organization_handle) 
VALUES (
    $1, 
    $2,
    $3
) RETURNING id, organization_name, organization_handle, created_at, updated_at, version
`

type CreateOrganizationParams struct {
	ID                 uuid.UUID `json:"id"`
	OrganizationName   string    `json:"organization_name"`
	OrganizationHandle string    `json:"organization_handle"`
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error) {
	row := q.db.QueryRow(ctx, createOrganization, arg.ID, arg.OrganizationName, arg.OrganizationHandle)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.OrganizationName,
		&i.OrganizationHandle,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createOrganizationRole = `-- name: CreateOrganizationRole :one
INSERT INTO organization_roles (id, organization_id, role_name, role_description, role_type)
VALUES (
    $1, 
    $2, 
    $3, 
    $4, 
    $5
) RETURNING id, organization_id, role_name, role_description, role_type, created_at, updated_at, version
`

type CreateOrganizationRoleParams struct {
	ID              uuid.UUID `json:"id"`
	OrganizationID  uuid.UUID `json:"organization_id"`
	RoleName        string    `json:"role_name"`
	RoleDescription *string   `json:"role_description"`
	RoleType        string    `json:"role_type"`
}

func (q *Queries) CreateOrganizationRole(ctx context.Context, arg CreateOrganizationRoleParams) (OrganizationRole, error) {
	row := q.db.QueryRow(ctx, createOrganizationRole,
		arg.ID,
		arg.OrganizationID,
		arg.RoleName,
		arg.RoleDescription,
		arg.RoleType,
	)
	var i OrganizationRole
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.RoleName,
		&i.RoleDescription,
		&i.RoleType,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createOrganizationUser = `-- name: CreateOrganizationUser :one
INSERT INTO organization_users (id, organization_id, user_id, role_id)
VALUES (
    $1, 
    $2, 
    $3, 
    $4
) RETURNING id, organization_id, user_id, role_id, created_at, updated_at, version
`

type CreateOrganizationUserParams struct {
	ID             uuid.UUID `json:"id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	UserID         uuid.UUID `json:"user_id"`
	RoleID         uuid.UUID `json:"role_id"`
}

func (q *Queries) CreateOrganizationUser(ctx context.Context, arg CreateOrganizationUserParams) (OrganizationUser, error) {
	row := q.db.QueryRow(ctx, createOrganizationUser,
		arg.ID,
		arg.OrganizationID,
		arg.UserID,
		arg.RoleID,
	)
	var i OrganizationUser
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.UserID,
		&i.RoleID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const getOrganizationByID = `-- name: GetOrganizationByID :one
SELECT organizations.id, organizations.organization_name, organizations.organization_handle, organizations.created_at, organizations.updated_at, organizations.version FROM organizations
WHERE organizations.id = $1
`

type GetOrganizationByIDRow struct {
	Organization Organization `json:"organization"`
}

func (q *Queries) GetOrganizationByID(ctx context.Context, id uuid.UUID) (GetOrganizationByIDRow, error) {
	row := q.db.QueryRow(ctx, getOrganizationByID, id)
	var i GetOrganizationByIDRow
	err := row.Scan(
		&i.Organization.ID,
		&i.Organization.OrganizationName,
		&i.Organization.OrganizationHandle,
		&i.Organization.CreatedAt,
		&i.Organization.UpdatedAt,
		&i.Organization.Version,
	)
	return i, err
}
