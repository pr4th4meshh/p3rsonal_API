package profile

import (
	"strings"

	models "github.com/pr4th4meshh/p3rsonal_API/internal/model"
	"github.com/pr4th4meshh/p3rsonal_API/internal/utils"
)

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

func (s *Service) Search(query string) map[string]any {
	query = strings.ToLower(strings.TrimSpace(query))

	var experiences []models.Experience
	var projects []models.Project
	var skills []string

	for _, item := range s.profile.Experience {
		if utils.Matches(
			query,
			item.Role,
			item.Company,
			item.Location,
			strings.Join(item.Stack, " "),
			strings.Join(item.Highlights, " "),
		) {
			experiences = append(experiences, item)
		}
	}

	for _, item := range s.profile.Projects {
		if utils.Matches(
			query,
			item.Name,
			item.Subtitle,
			strings.Join(item.Technologies, " "),
			strings.Join(item.Highlights, " "),
		) {
			projects = append(projects, item)
		}
	}

	skillCategories := [][]string{
		s.profile.Skills.Languages,
		s.profile.Skills.Databases,
		s.profile.Skills.Frontend,
		s.profile.Skills.Backend,
		s.profile.Skills.Mobile,
		s.profile.Skills.Testing,
		s.profile.Skills.DevOps,
		s.profile.Skills.Tools,
	}

	for _, category := range skillCategories {
		for _, skill := range category {
			if strings.Contains(strings.ToLower(skill), query) {
				skills = append(skills, skill)
			}
		}
	}

	return map[string]any{
		"query": query,
		"data": map[string]any{
			"experience": experiences,
			"projects":   projects,
			"skills":     skills,
		},
		"count": len(experiences) + len(projects) + len(skills),
	}
}
