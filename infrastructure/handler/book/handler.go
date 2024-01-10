package book

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/book"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
)

type handler struct {
	useCase  book.UseCase
	response response.Response
}

func newHandler(uc book.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Add(c echo.Context) error {
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
