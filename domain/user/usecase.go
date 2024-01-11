package user

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u *User) Register(user *model.User) error {
	user.BeforeCreate(user)
	return u.storage.Register(user)
}
func (u *User) List() (model.Users, error) {
	return u.storage.List()
}
func (u *User) Get(userId uuid.UUID) (*model.User, error) {
	return u.storage.Get(userId)
}
func (u *User) Update(userId uuid.UUID, payload interface{}) error {
	return u.storage.Update(userId, payload)
}
func (u *User) Delete(userId uuid.UUID) error {
	return u.storage.Delete(userId)
}
