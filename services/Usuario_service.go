package services

import (
	"errors"
	"gamerentapi/models"
	"gamerentapi/repositories"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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

    if !strings.Contains(usuario.Correo, "@") {
        return errors.New("el correo no tiene un formato válido")
    }

    if usuario.FechaRegistro.IsZero() {
        usuario.FechaRegistro = time.Now()
    }

    return s.UsuarioRepository.Create(usuario)
}

func (s *UsuarioService) FindAll() ([]*models.Usuario, error) {
    return s.UsuarioRepository.FindAll()
}

func (s *UsuarioService) FindByID(id bson.ObjectID) (*models.Usuario, error) {
    return s.UsuarioRepository.FindByID(id)
}

func (s *UsuarioService) Update(id bson.ObjectID, usuario *models.Usuario) error {
    return s.UsuarioRepository.Update(id, usuario)
}

func (s *UsuarioService) Delete(id bson.ObjectID) error {
    return s.UsuarioRepository.Delete(id)
}
