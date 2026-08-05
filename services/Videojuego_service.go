package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type VideojuegoService struct {
	Repository *repositories.VideojuegoRepository
}

func NewVideojuegoService(repository *repositories.VideojuegoRepository) *VideojuegoService {
	return &VideojuegoService{
		Repository: repository,
	}
}

// Crear videojuego
func (s *VideojuegoService) Create(videojuego *models.Videojuego) error {

	if videojuego.Titulo == "" {
		return errors.New("el título es obligatorio")
	}
	
	if videojuego.PrecioRenta <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}
	
	if videojuego.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}

	return s.Repository.Create(videojuego)
}

// Obtener todos
func (s *VideojuegoService) FindAll() ([]models.Videojuego, error) {
	return s.Repository.FindAll()
}


func (s *VideojuegoService) FindActivos() ([]models.Videojuego, error) {
	return s.Repository.FindActivos()
}
// Obtener por ID
func (s *VideojuegoService) FindByID(id string) (*models.Videojuego, error) {
	return s.Repository.FindByID(id)
}

// Actualizar
func (s *VideojuegoService) Update(id string, videojuego *models.Videojuego) error {
	return s.Repository.Update(id, videojuego)
}

// Eliminar
func (s *VideojuegoService) Delete(id string) error {
	return s.Repository.Delete(id)

}