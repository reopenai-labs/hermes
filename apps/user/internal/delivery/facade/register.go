package facade

import (
	"google.golang.org/grpc"
	"qiyibyte.com/hermes/apps/user/facade/userapi"
	"qiyibyte.com/hermes/apps/user/internal/repository"
)

func Register(server *grpc.Server) {
	userapi.RegisterUserProfileServer(server, &UserProfileFacade{repository: repository.UserProfile()})
}
