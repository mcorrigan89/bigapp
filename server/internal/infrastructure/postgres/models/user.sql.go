// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, given_name, family_name, email, email_verified) 
VALUES ($1, $2, $3, $4, $5::boolean) RETURNING id, given_name, family_name, email, email_verified, avatar_id, created_at, updated_at, version
`

type CreateUserParams struct {
	ID            uuid.UUID `json:"id"`
	GivenName     *string   `json:"given_name"`
	FamilyName    *string   `json:"family_name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.GivenName,
		arg.FamilyName,
		arg.Email,
		arg.EmailVerified,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GivenName,
		&i.FamilyName,
		&i.Email,
		&i.EmailVerified,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createUserSession = `-- name: CreateUserSession :one
INSERT INTO user_session (user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING id, user_id, impersonator_id, token, expires_at, user_expired, created_at, updated_at, version
`

type CreateUserSessionParams struct {
	UserID    uuid.UUID `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateUserSession(ctx context.Context, arg CreateUserSessionParams) (UserSession, error) {
	row := q.db.QueryRow(ctx, createUserSession, arg.UserID, arg.Token, arg.ExpiresAt)
	var i UserSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ImpersonatorID,
		&i.Token,
		&i.ExpiresAt,
		&i.UserExpired,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const expireUserSession = `-- name: ExpireUserSession :exec
UPDATE user_session SET user_expired = TRUE WHERE user_session.id = $1
`

func (q *Queries) ExpireUserSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, expireUserSession, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT users.id, users.given_name, users.family_name, users.email, users.email_verified, users.avatar_id, users.created_at, users.updated_at, users.version FROM users
WHERE users.email = $1
`

type GetUserByEmailRow struct {
	User User `json:"user"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.User.ID,
		&i.User.GivenName,
		&i.User.FamilyName,
		&i.User.Email,
		&i.User.EmailVerified,
		&i.User.AvatarID,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
		&i.User.Version,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT users.id, users.given_name, users.family_name, users.email, users.email_verified, users.avatar_id, users.created_at, users.updated_at, users.version FROM users
WHERE users.id = $1
`

type GetUserByIDRow struct {
	User User `json:"user"`
}

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.User.ID,
		&i.User.GivenName,
		&i.User.FamilyName,
		&i.User.Email,
		&i.User.EmailVerified,
		&i.User.AvatarID,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
		&i.User.Version,
	)
	return i, err
}

const getUserBySessionToken = `-- name: GetUserBySessionToken :one
SELECT users.id, users.given_name, users.family_name, users.email, users.email_verified, users.avatar_id, users.created_at, users.updated_at, users.version, user_session.id, user_session.user_id, user_session.impersonator_id, user_session.token, user_session.expires_at, user_session.user_expired, user_session.created_at, user_session.updated_at, user_session.version FROM users
JOIN user_session ON users.id = user_session.user_id
WHERE user_session.token = $1
`

type GetUserBySessionTokenRow struct {
	User        User        `json:"user"`
	UserSession UserSession `json:"user_session"`
}

func (q *Queries) GetUserBySessionToken(ctx context.Context, token string) (GetUserBySessionTokenRow, error) {
	row := q.db.QueryRow(ctx, getUserBySessionToken, token)
	var i GetUserBySessionTokenRow
	err := row.Scan(
		&i.User.ID,
		&i.User.GivenName,
		&i.User.FamilyName,
		&i.User.Email,
		&i.User.EmailVerified,
		&i.User.AvatarID,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
		&i.User.Version,
		&i.UserSession.ID,
		&i.UserSession.UserID,
		&i.UserSession.ImpersonatorID,
		&i.UserSession.Token,
		&i.UserSession.ExpiresAt,
		&i.UserSession.UserExpired,
		&i.UserSession.CreatedAt,
		&i.UserSession.UpdatedAt,
		&i.UserSession.Version,
	)
	return i, err
}

const setAvatarImage = `-- name: SetAvatarImage :one
UPDATE users SET avatar_id = $1 WHERE id = $2 RETURNING id, given_name, family_name, email, email_verified, avatar_id, created_at, updated_at, version
`

type SetAvatarImageParams struct {
	ImageID *uuid.UUID `json:"image_id"`
	UserID  uuid.UUID  `json:"user_id"`
}

func (q *Queries) SetAvatarImage(ctx context.Context, arg SetAvatarImageParams) (User, error) {
	row := q.db.QueryRow(ctx, setAvatarImage, arg.ImageID, arg.UserID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GivenName,
		&i.FamilyName,
		&i.Email,
		&i.EmailVerified,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET given_name = $1, family_name = $2 WHERE id = $3 RETURNING id, given_name, family_name, email, email_verified, avatar_id, created_at, updated_at, version
`

type UpdateUserParams struct {
	GivenName  *string   `json:"given_name"`
	FamilyName *string   `json:"family_name"`
	UserID     uuid.UUID `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.GivenName, arg.FamilyName, arg.UserID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GivenName,
		&i.FamilyName,
		&i.Email,
		&i.EmailVerified,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}
