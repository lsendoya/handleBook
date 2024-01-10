package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lsendoya/handleBook/infrastructure/handler/book"
	"github.com/lsendoya/handleBook/infrastructure/handler/loan"
	"github.com/lsendoya/handleBook/infrastructure/handler/user"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	book.NewRouter(e, db)
	loan.NewRouter(e, db)
	user.NewRouter(e, db)

}

func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"time":         time.Now().String(),
				"message":      "Hello World!",
				"service_name": "",
			},
		)
	})
}
