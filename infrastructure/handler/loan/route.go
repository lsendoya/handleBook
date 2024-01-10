package loan

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/loan"
	loanDB "github.com/lsendoya/handleBook/infrastructure/database/loan"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	h := buildHandler(db)

	adminRoutes(e, h)
}

func buildHandler(db *gorm.DB) handler {
	storage := loanDB.New(db)
	useCase := loan.New(storage)

	return newHandler(&useCase)
}

func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/loans")
	g.GET("", h.List)
	g.PUT("/:loanId", h.Update)
	g.POST("", h.Register)
}
