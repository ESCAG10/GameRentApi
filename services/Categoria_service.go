package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CategoriaService struct {
    CategoriaRepository *repositories.CategoriaRepository
}

func NewCategoriaService(repo *repositories.CategoriaRepository) *CategoriaService {
    return &CategoriaService{CategoriaRepository: repo}
}

func (s *CategoriaService) Create(categoria *models.Categoria) error {
    if categoria.Nombre == "" {
        return errors.New("el nombre es obligatorio")
    }
    return s.CategoriaRepository.Create(categoria)
}

func (s *CategoriaService) FindAll() ([]*models.Categoria, error) {
    return s.CategoriaRepository.FindAll()
}

func (s *CategoriaService) FindByID(id bson.ObjectID) (*models.Categoria, error) {
    return s.CategoriaRepository.FindByID(id)
}

func (s *CategoriaService) Update(id bson.ObjectID, categoria *models.Categoria) error {
    return s.CategoriaRepository.Update(id, categoria)
}

func (s *CategoriaService) Delete(id bson.ObjectID) error {
    return s.CategoriaRepository.Delete(id)
}
