package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type VideojuegoService struct {
	VideojuegoRepository *repositories.VideojuegoRepository
}

func NewVideojuegoService(repo *repositories.VideojuegoRepository) *VideojuegoService {
	return &VideojuegoService{
		VideojuegoRepository: repo,
	}
}

func (s *VideojuegoService) Create(videojuego *models.Videojuego) error {

	if videojuego.Titulo == "" {
		return errors.New("el título es obligatorio")
	}

	if videojuego.PrecioRenta <= 0 {
		return errors.New("el precio debe ser mayor a cero")
	}

	return s.VideojuegoRepository.Create(videojuego)
}