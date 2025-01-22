package entities

import (
	"github.com/google/uuid"
)

type EmailEntity struct {
	ID        uuid.UUID
	ToEmail   string
	FromEmail string
	Subject   string
	Body      string
}

type EmailEntityArgs struct {
	ID        uuid.UUID
	ToEmail   string
	FromEmail string
	Subject   string
	Body      string
}

func NewEmailEntity(args EmailEntityArgs) *EmailEntity {
	return &EmailEntity{
		ID:        args.ID,
		ToEmail:   args.ToEmail,
		FromEmail: args.FromEmail,
		Subject:   args.Subject,
		Body:      args.Body,
	}
}
