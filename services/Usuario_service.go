package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
)


type UsuarioService struct {
	Repository *repositories.UsuarioRepository
}


func NewUsuarioService(repository *repositories.UsuarioRepository,) *UsuarioService {

	return &UsuarioService{
		Repository: repository,
	}
}


// Crear usuario
func (s *UsuarioService) Create(usuario *models.Usuario,) error {
	

	if usuario.Nombre == "" {
		return  errors.New("el nombre es obligatorio")
	}

	if usuario.Correo == "" {
		return  errors.New("el correo es obligatorio")
	}

	if usuario.Password == "" {
		return  errors.New("la contraseña es obligatorio")
	}

	if usuario.Rol == "" {
		usuario.Rol = "cliente"
	}


	return s.Repository.Create(usuario)
}


// Obtener todos los usuarios
func (s *UsuarioService) FindAll() ([]models.Usuario,error,) {

	return s.Repository.FindAll()
}


// Obtener usuario por ID
func (s *UsuarioService) FindByID(id string,) (*models.Usuario, error) {

	return s.Repository.FindByID(id)
}


// Actualizar usuario
func (s *UsuarioService) Update(id string,usuario *models.Usuario,) error {

	return s.Repository.Update(
		id,
		usuario,
	)
}


// Eliminar usuario
func (s *UsuarioService) Delete(id string,) error {

	return s.Repository.Delete(id)
}

func (s *UsuarioService) Login(correo, password string) (*models.Usuario, error) {
	if correo == "" || password == "" {
		return nil, errors.New("correo y contraseña son obligatorios")
	}

	return s.Repository.FindByID(correo)
}