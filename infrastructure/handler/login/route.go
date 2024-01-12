package login

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/domain/login"
	"github.com/lsendoya/handleBook/domain/user"
	userDB "github.com/lsendoya/handleBook/infrastructure/database/user"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	h := buildHandler(db)

	publicRoutes(e, h)
}

func buildHandler(db *gorm.DB) handler {
	storage := userDB.New(db)
	useCaseUser := user.New(&storage)
	useCase := login.New(&useCaseUser)

	return newHandler(&useCase)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/login")
	g.POST("", h.Login)
}
