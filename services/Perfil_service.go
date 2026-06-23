package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type PerfilService struct {
	PerfilRepository *repositories.PerfilRepository
}

func NewPerfilService(repo *repositories.PerfilRepository) *PerfilService {
	return &PerfilService{
		PerfilRepository: repo,
	}
}

func (s *PerfilService) Create(perfil *models.Perfil) error {
	if perfil == nil {
		return errors.New("perfil cannot be nil")
	}
	return s.PerfilRepository.Create(perfil)
}