package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioRepository struct {
	Collection *mongo.Collection
}

func NewUsuarioRepository(collection *mongo.Collection) *UsuarioRepository {
	return &UsuarioRepository{
		Collection: collection,
	}
}

// Crear usuario
func (r *UsuarioRepository) Create(usuario *models.Usuario) error {

	ahora := time.Now()

	usuario.FechaRegistro = ahora
	usuario.FechaCreacion = ahora
	usuario.FechaActualizacion = ahora

	_, err := r.Collection.InsertOne(
		context.Background(),
		usuario,
	)

	return err
}


// Obtener todos los usuarios
func (r *UsuarioRepository) FindAll() ([]models.Usuario, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var usuarios []models.Usuario

	err = cursor.All(
		context.Background(),
		&usuarios,
	)

	return usuarios, err
}


// Obtener usuario por ID
func (r *UsuarioRepository) FindByID(id string) (*models.Usuario, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}


	var usuario models.Usuario


	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&usuario)


	if err != nil {
		return nil, err
	}


	return &usuario, nil
}


// Actualizar usuario
func (r *UsuarioRepository) Update(id string,usuario *models.Usuario,) error {


	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}


	// Recuperar datos actuales
	existente, err := r.FindByID(id)

	if err != nil {
		return err
	}


	// Mantener fecha original
	usuario.FechaRegistro = existente.FechaRegistro


	_, err = r.Collection.UpdateOne(
		context.Background(),

		bson.M{
			"_id": objectID,
		},

		bson.M{
			"$set": bson.M{

				"nombre": usuario.Nombre,

				"correo": usuario.Correo,

				"fechaRegistro": usuario.FechaRegistro,

				"fechaActualizacion": usuario.FechaActualizacion,

				"passwordHash": usuario.PasswordHash,

				"rol": usuario.Rol,

				"activo": usuario.Activo,
			},
		},
	)


	return err
}


// Eliminar usuario
func (r *UsuarioRepository) Delete(id string) error {


	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}


	_, err = r.Collection.DeleteOne(
		context.Background(),

		bson.M{
			"_id": objectID,
		},
	)


	return err
}