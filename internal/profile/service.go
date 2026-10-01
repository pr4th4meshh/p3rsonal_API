package profile

import models "github.com/pr4th4meshh/p3rsonal_API/internal/model"

// api.GET("/profile", h.profile)
// api.GET("/contact", h.contact)
// api.GET("/skills", h.skills)
// api.GET("/skills/:category", h.skillCategory)
// api.GET("/stack", h.stack)
// api.GET("/experience", h.experience)
// api.GET("/experience/:slug", h.experienceBySlug)
// api.GET("/projects", h.projects)
// api.GET("/projects/:slug", h.projectBySlug)
// api.GET("/education", h.education)
// api.GET("/stats", h.stats)
// api.GET("/rules", h.rules)
// api.GET("/search", h.search)

type Service struct {
	profile models.Profile
}

func NewService(profile models.Profile) *Service {
	return &Service{
		profile: profile,
	}
}

func (s *Service) GetProfile() models.Profile {
	return s.profile
}

func (s *Service) GetContact() models.Contact {
	return s.profile.Contact
}

func (s *Service) GetSkills() models.Skills {
	return s.profile.Skills
}

func (s *Service) GetExperience() []models.Experience {
	return s.profile.Experience
}

func (s *Service) GetProjects() []models.Project {
	return s.profile.Projects
}

func (s *Service) GetEducation() []models.Education {
	return s.profile.Education
}

func (s *Service) GetRules() []string {
	return s.profile.Rules
}
