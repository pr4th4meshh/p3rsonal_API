package api

import (
	"github.com/labstack/echo/v5"
	"github.com/pr4th4meshh/p3rsonal_API/internal/handlers"
)

func RegisterRoutes(e *echo.Echo, h *handlers.Handler) {
	e.GET("/me", h.Me)
	e.GET("/contact", h.Contact)
	e.GET("/skills", h.Skills)
	e.GET("/stack", h.Stack)
	e.GET("/experience", h.Experience)
	e.GET("/projects", h.Projects)
	e.GET("/education", h.Education)
	e.GET("/stats", h.Stats)
	e.GET("/rules", h.Rules)
	e.GET("/search", h.Search)
}
