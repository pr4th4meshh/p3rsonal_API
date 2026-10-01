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

func (h *Handler) Contact(c *echo.Context) error {
	data := h.profile.GetContact()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Skills(c *echo.Context) error {
	data := h.profile.GetSkills()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Stack(c *echo.Context) error {
	data := h.profile.GetSkills()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Experience(c *echo.Context) error {
	data := h.profile.GetExperience()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Projects(c *echo.Context) error {
	data := h.profile.GetProjects()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Education(c *echo.Context) error {
	data := h.profile.GetEducation()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Stats(c *echo.Context) error {
	data := h.profile.GetEducation()
	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Rules(c *echo.Context) error {
	data := h.profile.GetRules()
	return c.JSON(http.StatusOK, data)
}
