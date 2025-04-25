package service

import (
	"context"
	"qiyibyte.com/hermes/apps/user/internal/domain/po"
	"qiyibyte.com/hermes/apps/user/internal/errcode"
	"qiyibyte.com/hermes/apps/user/internal/repository"
	"qiyibyte.com/hermes/pkg/biz"
	"sync"
)

var accessService *AccessService
var accessInit sync.Once

func Access() *AccessService {
	accessInit.Do(func() {
		accessService = &AccessService{
			repository: repository.UserProfile(),
		}
	})
	return accessService
}

type AccessService struct {
	repository *repository.UserProfileRepository
}

func (Self *AccessService) Login(ctx context.Context, request *po.LoginRequest) (*po.LoginResponse, error) {
	userProfile, err := Self.repository.GetByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}
	if userProfile == nil {
		return nil, biz.NewCodeError(errcode.UsernameNotExists)
	}
	return nil, nil
}
