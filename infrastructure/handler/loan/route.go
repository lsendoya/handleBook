package loan

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/loan"
	bookDB "github.com/lsendoya/handleBook/infrastructure/database/book"
	loanDB "github.com/lsendoya/handleBook/infrastructure/database/loan"
	userDB "github.com/lsendoya/handleBook/infrastructure/database/user"
	"github.com/lsendoya/handleBook/infrastructure/handler/middleware"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	h := buildHandler(db)
	m := middleware.New()
	adminRoutes(e, h, m.TokenAuthenticator, m.AdminAuthorizer)
	privateRoute(e, h, m.TokenAuthenticator)
}

func buildHandler(db *gorm.DB) handler {
	loanStorage := loanDB.New(db)
	userStorage := userDB.New(db)
	bookStorage := bookDB.New(db)

	useCase := loan.New(loanStorage, bookStorage, &userStorage)

	return newHandler(&useCase)
}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/admin/loans", middlewares...)
	g.GET("", h.List)
	g.PUT("/:id", h.UpdateStatus)
	g.GET("/:id", h.Get)
}

func privateRoute(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/private/loans", middlewares...)
	g.POST("", h.Register)
}
