package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type RentaService struct {
	RentaRepository *repositories.RentaRepository
}

func NewRentaService(repo *repositories.RentaRepository) *RentaService {
	return &RentaService{
		RentaRepository: repo,
	}
}

func (s *RentaService) Create(renta *models.Renta) error {

	if renta.UsuarioID.IsZero() {
		return errors.New("usuario requerido")
	}

	if renta.VideojuegoID.IsZero() {
		return errors.New("videojuego requerido")
	}

	return s.RentaRepository.Create(renta)
}