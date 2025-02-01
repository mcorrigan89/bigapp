// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: ref_links.sql

package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createReferenceLink = `-- name: CreateReferenceLink :one
INSERT INTO reference_link (id, link_id, link_type, token, expires_at) 
VALUES ($1, $2, $3, $4, $5) RETURNING id, link_id, link_type, token, expires_at, created_at, updated_at, version
`

type CreateReferenceLinkParams struct {
	ID        uuid.UUID `json:"id"`
	LinkID    uuid.UUID `json:"link_id"`
	LinkType  string    `json:"link_type"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateReferenceLink(ctx context.Context, arg CreateReferenceLinkParams) (ReferenceLink, error) {
	row := q.db.QueryRow(ctx, createReferenceLink,
		arg.ID,
		arg.LinkID,
		arg.LinkType,
		arg.Token,
		arg.ExpiresAt,
	)
	var i ReferenceLink
	err := row.Scan(
		&i.ID,
		&i.LinkID,
		&i.LinkType,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const deleteReferenceLink = `-- name: DeleteReferenceLink :one
DELETE FROM reference_link WHERE reference_link.id = $1 RETURNING id, link_id, link_type, token, expires_at, created_at, updated_at, version
`

func (q *Queries) DeleteReferenceLink(ctx context.Context, id uuid.UUID) (ReferenceLink, error) {
	row := q.db.QueryRow(ctx, deleteReferenceLink, id)
	var i ReferenceLink
	err := row.Scan(
		&i.ID,
		&i.LinkID,
		&i.LinkType,
		&i.Token,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const getReferenceLinkByID = `-- name: GetReferenceLinkByID :one
SELECT reference_link.id, reference_link.link_id, reference_link.link_type, reference_link.token, reference_link.expires_at, reference_link.created_at, reference_link.updated_at, reference_link.version FROM reference_link
WHERE reference_link.id = $1
`

type GetReferenceLinkByIDRow struct {
	ReferenceLink ReferenceLink `json:"reference_link"`
}

func (q *Queries) GetReferenceLinkByID(ctx context.Context, id uuid.UUID) (GetReferenceLinkByIDRow, error) {
	row := q.db.QueryRow(ctx, getReferenceLinkByID, id)
	var i GetReferenceLinkByIDRow
	err := row.Scan(
		&i.ReferenceLink.ID,
		&i.ReferenceLink.LinkID,
		&i.ReferenceLink.LinkType,
		&i.ReferenceLink.Token,
		&i.ReferenceLink.ExpiresAt,
		&i.ReferenceLink.CreatedAt,
		&i.ReferenceLink.UpdatedAt,
		&i.ReferenceLink.Version,
	)
	return i, err
}

const getReferenceLinkByToken = `-- name: GetReferenceLinkByToken :one
SELECT reference_link.id, reference_link.link_id, reference_link.link_type, reference_link.token, reference_link.expires_at, reference_link.created_at, reference_link.updated_at, reference_link.version FROM reference_link
WHERE reference_link.token = $1
`

type GetReferenceLinkByTokenRow struct {
	ReferenceLink ReferenceLink `json:"reference_link"`
}

func (q *Queries) GetReferenceLinkByToken(ctx context.Context, token string) (GetReferenceLinkByTokenRow, error) {
	row := q.db.QueryRow(ctx, getReferenceLinkByToken, token)
	var i GetReferenceLinkByTokenRow
	err := row.Scan(
		&i.ReferenceLink.ID,
		&i.ReferenceLink.LinkID,
		&i.ReferenceLink.LinkType,
		&i.ReferenceLink.Token,
		&i.ReferenceLink.ExpiresAt,
		&i.ReferenceLink.CreatedAt,
		&i.ReferenceLink.UpdatedAt,
		&i.ReferenceLink.Version,
	)
	return i, err
}
