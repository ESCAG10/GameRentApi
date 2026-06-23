package repositories

import "gamerentapi/models"

type RentaRepository struct {
}

func NewRentaRepository() *RentaRepository {
	return &RentaRepository{}
}

func (r *RentaRepository) Create(renta *models.Renta) error {
	return nil
}

func (r *RentaRepository) FindAll() ([]*models.Renta, error) {
	return []*models.Renta{}, nil
}