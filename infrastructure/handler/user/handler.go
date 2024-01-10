package user

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/user"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
)

type handler struct {
	useCase  user.UseCase
	response response.Response
}

func newHandler(uc user.UseCase) handler {
	return handler{
		useCase: uc,
	}
}

func (h handler) Register(c echo.Context) error {
	return nil
}
func (h handler) List(c echo.Context) error {
	return nil
}
func (h handler) Get(c echo.Context) error {
	return nil
}
func (h handler) Update(c echo.Context) error {
	return nil
}
func (h handler) Delete(c echo.Context) error {
	return nil
}
