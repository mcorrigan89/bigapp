-- name: GetOrganizationByID :one
SELECT sqlc.embed(organizations) FROM organizations
WHERE organizations.id = $1;

-- name: CreateOrganization :one
INSERT INTO organizations (id, organization_name, organization_handle) 
VALUES (
    sqlc.arg(id), 
    sqlc.arg(organization_name),
    sqlc.arg(organization_handle)
) RETURNING *;

-- name: CreateOrganizationRole :one
INSERT INTO organization_roles (id, organization_id, role_name, role_description, role_type)
VALUES (
    sqlc.arg(id), 
    sqlc.arg(organization_id), 
    sqlc.arg(role_name), 
    sqlc.arg(role_description), 
    sqlc.arg(role_type)
) RETURNING *;

-- name: CreateOrganizationUser :one
INSERT INTO organization_users (id, organization_id, user_id, role_id)
VALUES (
    sqlc.arg(id), 
    sqlc.arg(organization_id), 
    sqlc.arg(user_id), 
    sqlc.arg(role_id)
) RETURNING *;