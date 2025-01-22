package dto

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mcorrigan89/simple_auth/server/internal/domain/entities"
)

type UserDto struct {
	ID         uuid.UUID `json:"id"`
	GivenName  *string   `json:"given_name"`
	FamilyName *string   `json:"family_name"`
	Email      string    `json:"email"`
}

func NewUserDtoFromEntity(entity *entities.UserEntity) *UserDto {
	return &UserDto{
		ID:         entity.ID,
		GivenName:  entity.GivenName,
		FamilyName: entity.FamilyName,
		Email:      entity.Email,
	}
}

func (dto *UserDto) ToJson() ([]byte, error) {
	return json.Marshal(dto)
}
