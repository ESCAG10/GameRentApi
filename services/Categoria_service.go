package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type CategoriaService struct {
	CategoriaRepository *repositories.CategoriaRepository
}

func NewCategoriaService(repo *repositories.CategoriaRepository) *CategoriaService {
	return &CategoriaService{
		CategoriaRepository: repo,
	}
}

func (s *CategoriaService) Create(categoria *models.Categoria) error {

	if categoria.Nombre == "" {
		return errors.New("el nombre de la categoría es obligatorio")
	}

	return s.CategoriaRepository.Create(categoria)
}