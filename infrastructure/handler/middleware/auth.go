package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
	"github.com/lsendoya/handleBook/model"
	"log"
	"net/http"
	"os"
	"strings"
)

type Auth struct {
	r response.Response
}

func New() Auth {
	return Auth{}
}

func (a *Auth) TokenAuthenticator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := a.getTokenAuthorization(c.Request())
		if err != nil {
			return c.JSON(a.r.Unauthorized(errors.New("authentication required")))
		}

		isValid, claims := a.isTokenAuthValid(token)
		if !isValid {
			return c.JSON(a.r.Unauthorized(errors.New("invalid or expired token")))
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		return next(c)
	}
}

func (a *Auth) AdminAuthorizer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role")
		if role != model.Admin {
			return c.JSON(a.r.Forbidden())
		}
		return next(c)
	}
}

func (a *Auth) getTokenAuthorization(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if strings.Compare(token, "") == 0 {
		return "", errors.New("the authorization header is empty")
	}

	if strings.HasPrefix(token, "Bearer") {
		return token[:7], nil
	}

	return token, nil
}

func (a *Auth) isTokenAuthValid(token string) (bool, model.JWTCustomClaims) {
	claims, err := jwt.ParseWithClaims(token, &model.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return false, model.JWTCustomClaims{}
	}

	data, ok := claims.Claims.(*model.JWTCustomClaims)
	if !ok {
		log.Println("is not a jwt custom claims")
		return false, model.JWTCustomClaims{}
	}

	return true, *data

}
