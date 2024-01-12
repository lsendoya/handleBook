package loan

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/loan"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
	"github.com/lsendoya/handleBook/model"
	"time"
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
	var loanData model.Loan
	if err := c.Bind(&loanData); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	data, errRegister := h.useCase.Register(loanData)
	if err := h.response.ValidateErr(c, "h.useCase.Register()", errRegister); err != nil {
		return err
	}

	return c.JSON(h.response.Created(data))
}
func (h handler) List(c echo.Context) error {
	loans, err := h.useCase.List()
	if err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.List()", err))
	}
	return c.JSON(h.response.OK(loans))
}
func (h handler) UpdateStatus(c echo.Context) error {

	var m struct {
		Status    model.LoanStatus `json:"status"`
		UpdatedAt time.Time        `json:"updated_at"`
	}
	if err := c.Bind(&m); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	if m.Status == "" {
		return c.JSON(h.response.BadRequest(errors.New("the status is mandatory")))
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	m.UpdatedAt = time.Now()

	errUpdate := h.useCase.UpdateStatus(id, &m)
	if err := h.response.ValidateErr(c, "h.useCase.UpdateStatus()", errUpdate); err != nil {
		return err
	}

	return c.JSON(h.response.Updated(m))
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
