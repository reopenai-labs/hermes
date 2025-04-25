package facade

import (
	"context"
	"qiyibyte.com/hermes/apps/user/facade/userapi"
	"qiyibyte.com/hermes/apps/user/internal/repository"
)

type UserProfileFacade struct {
	userapi.UnimplementedUserProfileServer
	repository *repository.UserProfileRepository
}

func (Self *UserProfileFacade) QueryProfiles(context.Context, *userapi.QueryProfileRequest) (*userapi.UserProfileReply, error) {
	return nil, nil
}
