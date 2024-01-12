package login

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/lsendoya/handleBook/model"
	"time"
)

type Login struct {
	useCaseUser UseCaseUser
}

func New(uc UseCaseUser) Login {
	return Login{uc}
}

func (l *Login) Login(email, password, jwtSecretKey string) (*model.User, string, error) {
	user, err := l.useCaseUser.Login(email, password)
	if err != nil {
		return nil, "", fmt.Errorf("%s %w", "useCaseUser.Login()", err)
	}

	claims := model.JWTCustomClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, errSigned := token.SignedString([]byte(jwtSecretKey))
	if errSigned != nil {
		return nil, "", fmt.Errorf("%s %w", "token.SignedString()", err)
	}

	user.Password = ""

	return user, signed, nil
}
