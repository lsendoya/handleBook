package user

import "github.com/lsendoya/handleBook/model"

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u *User) Register(user *model.User) error {
	return nil
}
func (u *User) List() (model.Users, error) {
	return nil, nil
}
func (u *User) Get(userId string) (*model.User, error) {
	return &model.User{}, nil
}
func (u *User) Update(userId string, payload interface{}) error {
	return nil
}
func (u *User) Delete(userId string) error {
	return nil
}
