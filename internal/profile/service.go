package profile

import models "github.com/pr4th4meshh/p3rsonal_API/internal/model"

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
