package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type RentaService struct {
	Repository *repositories.RentaRepository
	VideojuegoRepository *repositories.VideojuegoRepository
}

func NewRentaService(repository *repositories.RentaRepository, videojuegoRepository *repositories.VideojuegoRepository) *RentaService {
	return &RentaService{
		Repository: repository,
		VideojuegoRepository: videojuegoRepository,
	}
}

// Crear renta
func (s *RentaService) Create(renta *models.Renta) error {
	_, err := s.VideojuegoRepository.FindByID(renta.VideojuegoID.Hex())
	if err != nil {
		return err
	}

	return s.Repository.Create(renta)
}

// Obtener todas
func (s *RentaService) FindAll() ([]models.Renta, error) {
	return s.Repository.FindAll()
}

// Obtener por ID
func (s *RentaService) FindByID(id string) (*models.Renta, error) {
	return s.Repository.FindByID(id)
}

// Actualizar
func (s *RentaService) Update(id string, renta *models.Renta) error {
	return s.Repository.Update(id, renta)
}

// Eliminar
func (s *RentaService) Delete(id string) error {
	return s.Repository.Delete(id)
}
