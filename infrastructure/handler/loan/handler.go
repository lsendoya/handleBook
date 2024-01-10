package loan

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/loan"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
)

type handler struct {
	useCase  loan.UseCase
	response response.Response
}

func newHandler(uc loan.UseCase) handler {
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
func (h handler) Update(c echo.Context) error {
	return nil
}
