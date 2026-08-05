package repositories

import (
	"context"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PerfilRepository struct {
	Collection *mongo.Collection
}

func NewPerfilRepository(collection *mongo.Collection) *PerfilRepository {
	return &PerfilRepository{
		Collection: collection,
	}
}

// Crear perfil
func (r *PerfilRepository) Create(perfil *models.Perfil) error {

	// Fechas automáticas
	//ahora := time.Now()

	//perfil.FechaCreacion = ahora
	//perfil.FechaActualizacion = ahora

	_, err := r.Collection.InsertOne(
		context.Background(),
		perfil,
	)

	return err
}

// Obtener todos los perfiles
func (r *PerfilRepository) FindAll() ([]models.Perfil, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var perfiles []models.Perfil

	err = cursor.All(
		context.Background(),
		&perfiles,
	)

	return perfiles, err
}

// Obtener un perfil por ID
func (r *PerfilRepository) FindByID(id string) (*models.Perfil, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var perfil models.Perfil

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&perfil)

	if err != nil {
		return nil, err
	}

	return &perfil, nil
}

// Actualizar perfil
func (r *PerfilRepository) Update(id string, perfil *models.Perfil,) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	// Recuperar el perfil existente
	existente, err := r.FindByID(id)

	if err != nil {
		return err
	}

	// Mantener la fecha de creación
	perfil.FechaCreacion = existente.FechaCreacion

	// Actualizar la fecha de modificación
	//perfil.FechaActualizacion = time.Now()

	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{
				"usuarioId":          perfil.UsuarioID,
				"nombre":             perfil.Nombre,
				"direccion":          perfil.Direccion,
				"telefono":           perfil.Telefono,
				"nivelCliente":       perfil.NivelCliente,
				"fotoPerfil":         perfil.FotoPerfil,
				"fechaCreacion":      perfil.FechaCreacion,
				"fechaActualizacion": perfil.FechaActualizacion,
			},
		},
	)

	return err
}

// Eliminar perfil
func (r *PerfilRepository) Delete(id string) error {

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

func (r *PerfilRepository) FindByUsuarioID(usuarioID string) (*models.Perfil, error) {

	objectID, err := bson.ObjectIDFromHex(usuarioID)

	if err != nil {
		return nil, err
	}

	var perfil models.Perfil

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"usuarioId": objectID,
		},
	).Decode(&perfil)

	if err != nil {
		return nil, err
	}

	return &perfil, nil
}