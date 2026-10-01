package api

import (
	"github.com/labstack/echo/v5"
	"github.com/pr4th4meshh/p3rsonal_API/internal/handlers"
)

func RegisterRoutes(e *echo.Echo, h *handlers.Handler) {
	e.GET("/me", h.Me)
}
