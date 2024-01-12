package login

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/constants"
	"github.com/lsendoya/handleBook/domain/login"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
	"github.com/lsendoya/handleBook/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type handler struct {
	useCase  login.UseCase
	response response.Response
}

func newHandler(uc login.UseCase) handler {
	return handler{
		useCase: uc,
	}
}

func (h handler) Login(c echo.Context) error {
	m := model.Login{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	if m.Email == "" || m.Password == "" {
		return c.JSON(h.response.BadRequest(errors.New("all fields are mandatory")))
	}

	user, token, err := h.useCase.Login(m.Email, m.Password, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return c.JSON(h.response.NotFound(err))
		}

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			resp := model.Response{
				Code:    constants.MessageBadRequest,
				Message: "wrong user or password",
			}
			return c.JSON(http.StatusBadRequest, resp)
		}
		return c.JSON(h.response.InternalServerError(c, "h.useCase.Login", err))
	}

	return c.JSON(h.response.OK(map[string]interface{}{"user": user, "token": token}))

}
