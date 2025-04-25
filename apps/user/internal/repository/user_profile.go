package repository

import (
	"context"
	"gorm.io/gorm"
	"qiyibyte.com/hermes/apps/user/internal/domain/entity"
	"qiyibyte.com/hermes/pkg/infra/database"
	"sync"
)

var userProfileInit sync.Once
var userProfileRepository *UserProfileRepository

func UserProfile() *UserProfileRepository {
	userProfileInit.Do(func() {
		userProfileRepository = &UserProfileRepository{
			db: database.Instance(),
		}
	})
	return userProfileRepository
}

type UserProfileRepository struct {
	db *gorm.DB
}

func (Self *UserProfileRepository) GetByUsername(ctx context.Context, username string) (*entity.UserProfile, error) {
	var profile entity.UserProfile
	result := Self.db.WithContext(ctx).
		Where(&entity.UserProfile{Username: username}).
		Find(&profile)
	if result.RowsAffected > 0 {
		return &profile, result.Error
	}
	return nil, result.Error
}
