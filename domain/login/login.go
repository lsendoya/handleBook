package login

import "github.com/lsendoya/handleBook/model"

type UseCase interface {
	Login(email, password, jwtSecretKey string) (*model.User, string, error)
}

type UseCaseUser interface {
	Login(email, password string) (*model.User, error)
}
