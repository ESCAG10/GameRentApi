package repositories

import (
	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioRepository struct {
	Collection *mongo.Collection
}

func NewUsuarioRepository(collection *mongo.Collection) *UsuarioRepository {
	return &UsuarioRepository{
		Collection: collection,
	}
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) error {
	return nil
}

func (r *UsuarioRepository) FindAll() ([]*models.Usuario, error) {
	return []*models.Usuario{}, nil
}