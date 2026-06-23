package repositories

import "gamerentapi/models"

type VideojuegoRepository struct {
}

func NewVideojuegoRepository() *VideojuegoRepository {
	return &VideojuegoRepository{}
}

func (r *VideojuegoRepository) Create(videojuego *models.Videojuego) error {
	return nil
}

func (r *VideojuegoRepository) FindAll() ([]*models.Videojuego, error) {
	return []*models.Videojuego{}, nil
}