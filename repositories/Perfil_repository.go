package repositories

import "gamerentapi/models"

type PerfilRepository struct {
}

func NewPerfilRepository() *PerfilRepository {
	return &PerfilRepository{}
}

func (r *PerfilRepository) Create(perfil *models.Perfil) error {
	return nil
}

func (r *PerfilRepository) FindAll() ([]*models.Perfil, error) {
	return []*models.Perfil{}, nil
}