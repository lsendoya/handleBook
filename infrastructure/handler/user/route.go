package user

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/user"
	userDB "github.com/lsendoya/handleBook/infrastructure/database/user"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	h := buildHandler(db)

	adminRoutes(e, h)
}

func buildHandler(db *gorm.DB) handler {
	storage := userDB.New(db)
	useCase := user.New(&storage)
	return newHandler(&useCase)
}

func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/users")
	g.POST("", h.Register)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
	g.GET("/:id", h.Get)
	g.GET("", h.List)
}
