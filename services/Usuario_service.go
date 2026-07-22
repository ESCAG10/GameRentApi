package services

import (
	"gamerentapi/models"
	"gamerentapi/repositories"
)

type UsuarioService struct {
	Repository *repositories.UsuarioRepository
}

func NewUsuarioService(repository *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		Repository: repository,
	}
}

// Crear usuario
func (s *UsuarioService) Create(usuario *models.Usuario) error {
	return s.Repository.Create(usuario)
}

// Obtener todos los usuarios
func (s *UsuarioService) FindAll() ([]models.Usuario, error) {
	return s.Repository.FindAll()
}

// Obtener un usuario por ID
func (s *UsuarioService) FindByID(id string) (*models.Usuario, error) {
	return s.Repository.FindByID(id)
}

// Actualizar usuario
func (s *UsuarioService) Update(id string, usuario *models.Usuario) error {
	return s.Repository.Update(id, usuario)
}

// Eliminar usuario
func (s *UsuarioService) Delete(id string) error {
	return s.Repository.Delete(id)
}
