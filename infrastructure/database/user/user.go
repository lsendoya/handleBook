package user

import (
	"github.com/lsendoya/handleBook/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return User{db: db}
}

func (u User) Register(user *model.User) error {
	return nil
}
func (u User) List() (model.Users, error) {
	return nil, nil
}

func (u User) Get(userId string) (*model.User, error) {
	return &model.User{}, nil
}
func (u User) Update(userId string, payload interface{}) error {
	return nil
}
func (u User) Delete(userId string) error {
	return nil
}
