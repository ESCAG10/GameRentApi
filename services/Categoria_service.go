package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type CategoriaService struct {
	Repository *repositories.CategoriaRepository
}

func NewCategoriaService(repository *repositories.CategoriaRepository) *CategoriaService {
	return &CategoriaService{
		Repository: repository,
	}
}

// Crear categoría
func (s *CategoriaService) Create(categoria *models.Categoria) error {
	return s.Repository.Create(categoria)
}

// Obtener todas
func (s *CategoriaService) FindAll() ([]models.Categoria, error) {
	return s.Repository.FindAll()
}

// Obtener por ID
func (s *CategoriaService) FindByID(id string) (*models.Categoria, error) {
	return s.Repository.FindByID(id)
}

// Actualizar
func (s *CategoriaService) Update(id string, categoria *models.Categoria) error {
	return s.Repository.Update(id, categoria)
}

// Eliminar
func (s *CategoriaService) Delete(id string) error {
	return s.Repository.Delete(id)
}