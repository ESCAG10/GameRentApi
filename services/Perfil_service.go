package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type PerfilService struct {
    PerfilRepository *repositories.PerfilRepository
}

func NewPerfilService(repo *repositories.PerfilRepository) *PerfilService {
    return &PerfilService{PerfilRepository: repo}
}

func (s *PerfilService) Create(perfil *models.Perfil) error {
    if perfil.Nombre == "" {
        return errors.New("el nombre es obligatorio")
    }
    if perfil.Telefono == "" {
        return errors.New("el teléfono es obligatorio")
    }
    return s.PerfilRepository.Create(perfil)
}

func (s *PerfilService) FindAll() ([]*models.Perfil, error) {
    return s.PerfilRepository.FindAll()
}

func (s *PerfilService) FindByID(id bson.ObjectID) (*models.Perfil, error) {
    return s.PerfilRepository.FindByID(id)
}

func (s *PerfilService) Update(id bson.ObjectID, perfil *models.Perfil) error {
    return s.PerfilRepository.Update(id, perfil)
}

func (s *PerfilService) Delete(id bson.ObjectID) error {
    return s.PerfilRepository.Delete(id)
}
