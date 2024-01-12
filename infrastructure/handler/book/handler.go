package book

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/book"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
	"github.com/lsendoya/handleBook/model"
	"time"
)

type handler struct {
	useCase  book.UseCase
	response response.Response
}

func newHandler(uc book.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Add(c echo.Context) error {
	m := model.Book{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	if m.Author == "" || m.Title == "" || m.PublicationYear < 0 || m.AvailableQuantity < 0 || m.Publisher == "" {
		return c.JSON(h.response.BadRequest(fmt.Errorf("invalid fields of book")))
	}

	if err := h.useCase.Add(&m); err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.Add()", err))
	}

	return c.JSON(h.response.Created(m))
}

func (h handler) List(c echo.Context) error {
	books, err := h.useCase.List()
	if err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.List()", err))
	}
	return c.JSON(h.response.OK(books))
}

func (h handler) Get(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	mdl, errGet := h.useCase.Get(id)
	if err := h.response.ValidateErr(c, "h.useCase.Get()", errGet); err != nil {
		return err
	}

	return c.JSON(h.response.OK(mdl))
}
func (h handler) Update(c echo.Context) error {
	var m struct {
		Title             string    `json:"title"`
		Author            string    `json:"author"`
		PublicationYear   int       `json:"publication_year"`
		AvailableQuantity int       `json:"available_quantity"`
		UpdatedAt         time.Time `json:"updated_at"`
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	m.UpdatedAt = time.Now()

	if err := c.Bind(&m); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	errUpdate := h.useCase.Update(id, m)
	if err := h.response.ValidateErr(c, "h.useCase.UpdateStatus()", errUpdate); err != nil {
		return err
	}
	return c.JSON(h.response.Updated(m))
}
func (h handler) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	errDelete := h.useCase.Delete(id)
	if err := h.response.ValidateErr(c, "h.useCase.Delete", errDelete); err != nil {
		return err
	}

	return c.JSON(h.response.Deleted(id))
}
