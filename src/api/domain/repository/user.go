package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/uma-co82/Shupple-api/src/api/domain"
	"github.com/uma-co82/Shupple-api/src/api/domain/user"
)

type UserRepository interface {
	GetByUid(uid string) (*user.User, error)
	GetUserInformationByRelatedUser(*user.User) error
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) GetByUid(uid string) (*user.User, error) {
	var person user.User
	if err := u.conn.First(&person, "uid=?", uid).Error; err != nil {
		return &person, domain.RaiseDBError()
	}
	return &person, nil
}

func (u *userRepository) GetUserInformationByRelatedUser(person *user.User) error {
	if err := u.conn.Model(person).Related(person.UserInformation, "UserInformation").Error; err != nil {
		return domain.RaiseDBError()
	}
	return nil
}
