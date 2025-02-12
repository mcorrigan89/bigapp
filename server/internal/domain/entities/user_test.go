package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mcorrigan89/bigapp/server/internal/infrastructure/postgres/models"
	"github.com/stretchr/testify/assert"
)

func TestNewUserEntity(t *testing.T) {

	t.Run("create entity from model", func(t *testing.T) {
		userId := uuid.New()
		givenName := "John"
		familyName := "Doe"
		email := "johndoe@gmail.com"
		handle := "johndoe"

		userModel := models.User{
			ID:            userId,
			GivenName:     &givenName,
			FamilyName:    &familyName,
			Email:         email,
			EmailVerified: true,
			UserHandle:    handle,
			Claimed:       true,
			AvatarID:      nil,
		}

		userEntity := NewUserEntity(userModel, nil)

		assert.Equal(t, userEntity.ID, userId)
		assert.Equal(t, *userEntity.GivenName, givenName)
		assert.Equal(t, *userEntity.FamilyName, familyName)
		assert.Equal(t, userEntity.Email, email)
		assert.Equal(t, userEntity.EmailVerified, true)
		assert.Equal(t, userEntity.Claimed, true)
		assert.Equal(t, userEntity.Handle, handle)
	})

	t.Run("create entity from model with avatar", func(t *testing.T) {
		userId := uuid.New()
		givenName := "John"
		familyName := "Doe"
		email := "johndoe@gmail.com"
		handle := "johndoe"

		userModel := models.User{
			ID:            userId,
			GivenName:     &givenName,
			FamilyName:    &familyName,
			Email:         email,
			EmailVerified: true,
			UserHandle:    handle,
			Claimed:       true,
			AvatarID:      nil,
		}

		imageId := uuid.New()
		avatar := models.Image{
			ID:         imageId,
			BucketName: "image",
			ObjectID:   "image.jpg",
			Height:     100,
			Width:      100,
			FileSize:   1000,
		}

		avatarEntity := NewImageEntity(avatar)

		userEntity := NewUserEntity(userModel, avatarEntity)

		assert.Equal(t, userEntity.ID, userId)
		assert.Equal(t, *userEntity.GivenName, givenName)
		assert.Equal(t, *userEntity.FamilyName, familyName)
		assert.Equal(t, userEntity.Email, email)
		assert.Equal(t, userEntity.EmailVerified, true)
		assert.Equal(t, userEntity.Claimed, true)
		assert.Equal(t, userEntity.Handle, handle)
		assert.Equal(t, userEntity.Avatar, avatarEntity)
	})
}

func TestUserEntityFullName(t *testing.T) {
	t.Run("user entity full name", func(t *testing.T) {
		userId := uuid.New()
		givenName := "John"
		familyName := "Doe"
		email := "johndoe@gmail.com"
		handle := "johndoe"

		userModel := models.User{
			ID:            userId,
			GivenName:     &givenName,
			FamilyName:    &familyName,
			Email:         email,
			EmailVerified: true,
			Claimed:       true,
			UserHandle:    handle,
			AvatarID:      nil,
		}

		userEntity := NewUserEntity(userModel, nil)

		assert.Equal(t, *userEntity.FullName(), "John Doe")
	})
}
