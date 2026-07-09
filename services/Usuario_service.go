package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type UsuarioService struct {
	UsuarioRepository *repositories.UsuarioRepository
}

func NewUsuarioService(repo *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		UsuarioRepository: repo,
	}
}

func (s *UsuarioService) Create(usuario *models.Usuario) error {

	if usuario.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if usuario.Correo == "" {
		return errors.New("el correo es obligatorio")
	}

	return s.UsuarioRepository.Create(usuario)
}

func (s *UsuarioService) FindAll() ([]*models.Usuario, error) {
	return s.UsuarioRepository.FindAll()
}

func (s *UsuarioService) FindByID(id string) (*models.Usuario, error) {
	return s.UsuarioRepository.FindByID(id)
}

func (s *UsuarioService) Update(id string, usuario *models.Usuario) error {
	return s.UsuarioRepository.Update(id, usuario)
}

func (s *UsuarioService) Delete(id string) error {
	return s.UsuarioRepository.Delete(id)
}