package book

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/book"
	bookDB "github.com/lsendoya/handleBook/infrastructure/database/book"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	h := buildHandler(db)

	adminRoutes(e, h)
	publicRoutes(e, h)
}

func buildHandler(db *gorm.DB) handler {
	storage := bookDB.New(db)
	useCase := book.New(storage)

	return newHandler(&useCase)

}

func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/books")
	g.POST("", h.Add)
	g.PUT("/:bookId", h.Update)
	g.DELETE("/:bookId", h.Delete)

}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/books")
	g.GET("/:bookId", h.Get)
	g.GET("", h.List)
}
