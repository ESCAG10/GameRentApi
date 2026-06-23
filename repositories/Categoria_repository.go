package repositories

import "gamerentapi/models"

type CategoriaRepository struct {
}

func NewCategoriaRepository() *CategoriaRepository {
	return &CategoriaRepository{}
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) error {
	return nil
}

func (r *CategoriaRepository) FindAll() ([]*models.Categoria, error) {
	return []*models.Categoria{}, nil
}