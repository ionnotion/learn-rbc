package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echo.New()

// Recieve and pass Handler with parameters
func Init() {
	// Define middlewares here
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		MaxAge:           24 * 60 * 60,
	}))

	// Define routes Here
}

func Start(address string) error {
	return e.Start(address)
}
