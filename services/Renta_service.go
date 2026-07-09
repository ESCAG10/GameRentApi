package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type RentaService struct {
	Repository *repositories.RentaRepository
}

func NewRentaService(repository *repositories.RentaRepository) *RentaService {
	return &RentaService{
		Repository: repository,
	}
}

func (s *RentaService) Create(renta *models.Renta) error {
	return s.Repository.Create(renta)
}

func (s *RentaService) FindAll() ([]models.Renta, error) {
	return s.Repository.FindAll()
}

func (s *RentaService) FindByID(id string) (*models.Renta, error) {
	return s.Repository.FindByID(id)
}

func (s *RentaService) Update(id string, renta *models.Renta) error {
	return s.Repository.Update(id, renta)
}

func (s *RentaService) Delete(id string) error {
	return s.Repository.Delete(id)
}