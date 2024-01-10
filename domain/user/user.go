package user

import "github.com/lsendoya/handleBook/model"

type UseCase interface {
	Register(user *model.User) error
	List() (model.Users, error)
	Get(userId string) (*model.User, error)
	Update(userId string, payload interface{}) error
	Delete(userId string) error
}

type Storage interface {
	Register(user *model.User) error
	List() (model.Users, error)
	Get(userId string) (*model.User, error)
	Update(userId string, payload interface{}) error
	Delete(userId string) error
}
