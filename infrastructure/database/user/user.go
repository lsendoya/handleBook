package user

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) User {
	return User{db: db}
}

func (u *User) Register(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (u *User) List() (model.Users, error) {
	var users model.Users
	if err := u.db.Omit("password").Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error retrieving users: %w", err)
	}
	return users, nil
}

func (u *User) Get(userId uuid.UUID) (*model.User, error) {
	var user model.User
	result := u.db.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) Update(userId uuid.UUID, payload interface{}) error {

	mdl, err := u.Get(userId)
	if err != nil {
		return err
	}

	if err := u.db.Model(mdl).Updates(payload).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(userId uuid.UUID) error {
	result := u.db.Delete(&model.User{}, "id = ?", userId)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u *User) GetByEmail(email string) (*model.User, error) {
	var user model.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}
