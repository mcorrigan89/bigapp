// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package models

import (
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID             uuid.UUID  `json:"id"`
	CollectionName string     `json:"collection_name"`
	OwnerID        uuid.UUID  `json:"owner_id"`
	Public         bool       `json:"public"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Version        int32      `json:"version"`
}

type CollectionImage struct {
	CollectionID uuid.UUID `json:"collection_id"`
	ImageID      uuid.UUID `json:"image_id"`
	SortKey      string    `json:"sort_key"`
}

type Image struct {
	ID         uuid.UUID  `json:"id"`
	BucketName string     `json:"bucket_name"`
	ObjectID   string     `json:"object_id"`
	Height     int32      `json:"height"`
	Width      int32      `json:"width"`
	FileSize   int32      `json:"file_size"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Version    int32      `json:"version"`
	OwnerID    *uuid.UUID `json:"owner_id"`
}

type ReferenceLink struct {
	ID        uuid.UUID  `json:"id"`
	LinkID    uuid.UUID  `json:"link_id"`
	LinkType  string     `json:"link_type"`
	Token     string     `json:"token"`
	ExpiresAt time.Time  `json:"expires_at"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Version   int32      `json:"version"`
}

type SchemaMigration struct {
	Version int64 `json:"version"`
	Dirty   bool  `json:"dirty"`
}

type User struct {
	ID            uuid.UUID  `json:"id"`
	GivenName     *string    `json:"given_name"`
	FamilyName    *string    `json:"family_name"`
	Email         string     `json:"email"`
	EmailVerified bool       `json:"email_verified"`
	UserHandle    string     `json:"user_handle"`
	Claimed       bool       `json:"claimed"`
	AvatarID      *uuid.UUID `json:"avatar_id"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Version       int32      `json:"version"`
}

type UserSession struct {
	ID             uuid.UUID  `json:"id"`
	UserID         uuid.UUID  `json:"user_id"`
	ImpersonatorID *uuid.UUID `json:"impersonator_id"`
	Token          string     `json:"token"`
	ExpiresAt      time.Time  `json:"expires_at"`
	UserExpired    bool       `json:"user_expired"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Version        int32      `json:"version"`
}
