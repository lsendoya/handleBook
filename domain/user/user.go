package user

import (
	"github.com/google/uuid"
	"github.com/lsendoya/handleBook/model"
)

type UseCase interface {
	Register(user *model.User) error
	List() (model.Users, error)
	Get(userId uuid.UUID) (*model.User, error)
	Update(userId uuid.UUID, payload interface{}) error
	Delete(userId uuid.UUID) error
}

type Storage interface {
	Register(user *model.User) error
	List() (model.Users, error)
	Get(userId uuid.UUID) (*model.User, error)
	Update(userId uuid.UUID, payload interface{}) error
	Delete(userId uuid.UUID) error
}
