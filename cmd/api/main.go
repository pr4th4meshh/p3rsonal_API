package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pr4th4meshh/p3rsonal_API/internal/api"
	"github.com/pr4th4meshh/p3rsonal_API/internal/handlers"
	"github.com/pr4th4meshh/p3rsonal_API/internal/profile"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	profileService := profile.NewService(profile.ProfileData())
	handler := handlers.New(profileService)

	api.RegisterRoutes(e, handler)

	e.GET("/health", func(ec *echo.Context) error {
		return ec.JSON(http.StatusOK, map[string]string{"message": "pr4th4meshh api healthy"})
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
