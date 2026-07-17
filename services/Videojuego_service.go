package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type VideojuegoService struct {
    VideojuegoRepository *repositories.VideojuegoRepository
}

func NewVideojuegoService(repo *repositories.VideojuegoRepository) *VideojuegoService {
    return &VideojuegoService{VideojuegoRepository: repo}
}

func (s *VideojuegoService) Create(videojuego *models.Videojuego) error {
    if videojuego.Titulo == "" {
        return errors.New("el título es obligatorio")
    }
    if videojuego.PrecioRenta <= 0 {
        return errors.New("el precio de renta debe ser mayor a 0")
    }
    if videojuego.Stock < 0 {
        return errors.New("el stock no puede ser negativo")
    }
    return s.VideojuegoRepository.Create(videojuego)
}

func (s *VideojuegoService) FindAll() ([]*models.Videojuego, error) {
    return s.VideojuegoRepository.FindAll()
}

func (s *VideojuegoService) FindByID(id bson.ObjectID) (*models.Videojuego, error) {
    return s.VideojuegoRepository.FindByID(id)
}

func (s *VideojuegoService) Update(id bson.ObjectID, videojuego *models.Videojuego) error {
    return s.VideojuegoRepository.Update(id, videojuego)
}

func (s *VideojuegoService) Delete(id bson.ObjectID) error {
    return s.VideojuegoRepository.Delete(id)
}
