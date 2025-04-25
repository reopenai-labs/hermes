package entity

const TableNameUserProfile = "users.user_profile"

type UserProfile struct {
	Id       uint64 `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
}

func (Self *UserProfile) TableName() string {
	return TableNameUserProfile
}
