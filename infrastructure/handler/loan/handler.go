package loan

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/loan"
	"github.com/lsendoya/handleBook/infrastructure/handler/response"
	"github.com/lsendoya/handleBook/model"
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
	var loanData struct {
		UserID uuid.UUID `json:"user_id"`
		BookID uuid.UUID `json:"book_id"`
	}
	if err := c.Bind(&loanData); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	if err := h.useCase.Register(loanData); err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.Register()", err))
	}

	return c.JSON(h.response.Created(loanData))
}
func (h handler) List(c echo.Context) error {
	loans, err := h.useCase.List()
	if err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.List()", err))
	}
	return c.JSON(h.response.OK(loans))
}
func (h handler) Update(c echo.Context) error {
	var mdl = model.Loan{}
	if err := c.Bind(&mdl); err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(h.response.BadRequest(err))
	}

	err = h.useCase.Update(id, &mdl)
	if err != nil {
		return c.JSON(h.response.InternalServerError(c, "h.useCase.Update()", err))
	}

	return c.JSON(h.response.Updated(mdl))
}
