package user

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/user"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
	"github.com/lsendoya/handleBook/model"
	"time"
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
	m := model.User{}

	if err := c.Bind(&m); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	if m.Name == "" || m.Email == "" || m.Password == "" {
		return c.JSON(h.response.BadRequest(fmt.Errorf("invalid fields of user")))
	}

	if err := h.useCase.Register(&m); err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.Register()", err))
	}

	return c.JSON(h.response.Created(m))
}

func (h handler) List(c echo.Context) error {
	users, err := h.useCase.List()
	if err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.List()", err))
	}

	return c.JSON(h.response.OK(users))
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
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	if err := c.Bind(&m); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	m.UpdatedAt = time.Now()

	errUpdate := h.useCase.Update(id, &m)
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
