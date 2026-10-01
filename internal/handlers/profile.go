package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pr4th4meshh/p3rsonal_API/internal/profile"
)

type Handler struct {
	profile *profile.Service
}

func New(service *profile.Service) *Handler {
	return &Handler{
		profile: service,
	}
}

func (h *Handler) Me(c *echo.Context) error {
	data := h.profile.GetProfile()
	return c.JSON(http.StatusOK, data)
}
