package book

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/book"
	bookDB "github.com/lsendoya/handleBook/infrastructure/database/book"
	"github.com/lsendoya/handleBook/infrastructure/handler/middleware"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	h := buildHandler(db)
	m := middleware.New()

	adminRoutes(e, h, m.TokenAuthenticator, m.AdminAuthorizer)
	publicRoutes(e, h)
}

func buildHandler(db *gorm.DB) handler {
	storage := bookDB.New(db)
	useCase := book.New(storage)

	return newHandler(&useCase)

}

func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/admin/books", middlewares...)
	g.POST("", h.Add)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)

}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/books")
	g.GET("/:id", h.Get)
	g.GET("", h.List)
}
