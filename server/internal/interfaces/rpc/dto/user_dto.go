package dto

import (
	"github.com/mcorrigan89/bigapp/server/internal/domain/entities"
	imagev1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/media/v1"
	userv1 "github.com/mcorrigan89/bigapp/server/internal/interfaces/rpc/gen/user/v1"
)

func UserEntityToDto(userEntity *entities.UserEntity) *userv1.User {
	var avatar *imagev1.Image
	if userEntity.Avatar != nil {
		avatar = &imagev1.Image{
			Id:     userEntity.Avatar.ID.String(),
			Url:    userEntity.Avatar.UrlSlug(),
			Width:  userEntity.Avatar.Width,
			Height: userEntity.Avatar.Height,
			Size:   userEntity.Avatar.Size,
		}
	}

	user := &userv1.User{
		Id:         userEntity.ID.String(),
		GivenName:  userEntity.GivenName,
		FamilyName: userEntity.FamilyName,
		FullName:   userEntity.FullName(),
		Email:      userEntity.Email,
		Avatar:     avatar,
		Handle:     userEntity.Handle,
	}

	return user
}
