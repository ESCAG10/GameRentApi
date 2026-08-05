package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type RentaService struct {
	Repository           *repositories.RentaRepository
	VideojuegoRepository *repositories.VideojuegoRepository
}

func NewRentaService(
	repository *repositories.RentaRepository,
	videojuegoRepository *repositories.VideojuegoRepository,
) *RentaService {
	return &RentaService{
		Repository:           repository,
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

// Obtener todas las rentas
func (s *RentaService) FindAll() ([]models.Renta, error) {
	return s.Repository.FindAll()
}

// Obtener renta por ID (_id de la renta)
func (s *RentaService) FindByID(id string) (*models.Renta, error) {
	return s.Repository.FindByID(id)
}

// Obtener todas las rentas de un usuario
func (s *RentaService) FindByUsuarioID(usuarioID string) ([]models.Renta, error) {
	return s.Repository.FindByUsuarioID(usuarioID)
}

// Actualizar renta
func (s *RentaService) Update(id string, renta *models.Renta) error {
	return s.Repository.Update(id, renta)
}

// Eliminar renta
func (s *RentaService) Delete(id string) error {
	
	return s.Repository.Delete(id)
}