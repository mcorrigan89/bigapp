package queries

import (
	"github.com/google/uuid"
)

type UserByIDQuery struct {
	ID uuid.UUID
}

type UserByEmailQuery struct {
	Email string `validate:"required,email"`
}

type UserBySessionTokenQuery struct {
	SessionToken string
}
