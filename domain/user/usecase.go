package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{s}
}

func (u *User) Register(user *model.User) error {
	user.BeforeCreate(user)

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	user.Password = string(password)

	return u.storage.Register(user)
}
func (u *User) List() (model.Users, error) {
	return u.storage.List()
}
func (u *User) Get(userId uuid.UUID) (*model.User, error) {
	m, err := u.storage.Get(userId)

	if err != nil {
		return nil, err
	}
	m.Password = ""

	return m, nil
}
func (u *User) Update(userId uuid.UUID, payload interface{}) error {
	return u.storage.Update(userId, payload)
}
func (u *User) Delete(userId uuid.UUID) error {
	return u.storage.Delete(userId)
}
func (u *User) GetByEmail(email string) (*model.User, error) {
	return u.storage.GetByEmail(email)
}
func (u *User) Login(email, password string) (*model.User, error) {
	m, err := u.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "user.GetByEmail()", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("%s %w", "bcrypt.CompareHashAndPassword()", err)
	}

	m.Password = ""

	return m, nil
}
