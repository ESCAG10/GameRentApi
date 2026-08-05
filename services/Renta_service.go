package services

import (
	"errors"
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

	videojuego, err := s.VideojuegoRepository.FindByID(renta.VideojuegoID.Hex())
	if err != nil {
		return err
	}

	if videojuego.Stock <= 0 {
		return errors.New("el videojuego no tiene stock disponible")
	}

	existe, err := s.Repository.ExisteRentaActiva(
		renta.UsuarioID,
		renta.VideojuegoID,
	)

	if err != nil {
		return err
	}

	if existe {
		return errors.New("ya existe una renta activa para este videojuego")
	}

	if err := s.VideojuegoRepository.DecreaseStock(
		renta.VideojuegoID.Hex(),
	); err != nil {
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
func (s *RentaService) FindByUsuarioID(id string) ([]models.Renta, error) {
	return s.Repository.FindByUsuarioID(id)
}

// Actualizar renta
func (s *RentaService) Update(id string, renta *models.Renta) error {

	// Obtener videojuego asociado a la renta
	videojuego, err := s.VideojuegoRepository.FindByID(renta.VideojuegoID.Hex())
	if err != nil {
		return err
	}

	// Decrementar stock y actualizar videojuego
	if videojuego.Stock <= 0 {
		return errors.New("el videojuego no tiene stock disponible")
	}
	videojuego.Stock--

	if err := s.VideojuegoRepository.Update(videojuego.ID.Hex(), videojuego); err != nil {
		return err
	}

	return s.Repository.Update(id, renta)
}

// Eliminar renta
func (s *RentaService) Delete(id string) error {

	renta, err := s.Repository.FindByID(id)

	if err != nil {
		return err
	}

	if err := s.VideojuegoRepository.IncrementarStock(renta.VideojuegoID); err != nil {
		return err
	}

	return s.Repository.Delete(id)
}