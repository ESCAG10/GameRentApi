package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type RentaService struct {
    RentaRepository *repositories.RentaRepository
}

func NewRentaService(repo *repositories.RentaRepository) *RentaService {
    return &RentaService{RentaRepository: repo}
}

func (s *RentaService) Create(renta *models.Renta) error {
    if renta.PeriodoRenta <= 0 {
        return errors.New("el periodo de renta debe ser mayor a 0 días")
    }
    if renta.Estado == "" {
        renta.Estado = "activa"
    }
    return s.RentaRepository.Create(renta)
}

func (s *RentaService) FindAll() ([]*models.Renta, error) {
    return s.RentaRepository.FindAll()
}

func (s *RentaService) FindByID(id bson.ObjectID) (*models.Renta, error) {
    return s.RentaRepository.FindByID(id)
}

func (s *RentaService) Update(id bson.ObjectID, renta *models.Renta) error {
    return s.RentaRepository.Update(id, renta)
}

func (s *RentaService) Delete(id bson.ObjectID) error {
    return s.RentaRepository.Delete(id)
}
