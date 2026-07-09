package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type PerfilService struct {
	Repository *repositories.PerfilRepository
}

func NewPerfilService(repository *repositories.PerfilRepository) *PerfilService {

	return &PerfilService{
		Repository: repository,
	}
}

func (s *PerfilService) Create(perfil *models.Perfil) error {
	return s.Repository.Create(perfil)
}

func (s *PerfilService) FindAll() ([]models.Perfil, error) {
	return s.Repository.FindAll()
}

func (s *PerfilService) FindByID(id string) (*models.Perfil, error) {
	return s.Repository.FindByID(id)
}

func (s *PerfilService) Update(id string, perfil *models.Perfil) error {
	return s.Repository.Update(id, perfil)
}

func (s *PerfilService) Delete(id string) error {
	return s.Repository.Delete(id)
}