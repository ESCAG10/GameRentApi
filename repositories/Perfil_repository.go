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

func (r *PerfilRepository) Create(perfil *models.Perfil) error {

	_, err := r.Collection.InsertOne(
		context.Background(),
		perfil,
	)

	return err
}

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

func (r *PerfilRepository) Update(id string, perfil *models.Perfil,) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{
				"nombre": perfil.Nombre,
				"telefono": perfil.Telefono,
				"direccion": perfil.Direccion,
				"nivelcliente": perfil.NivelCliente,
			},
		},
	)

	return err
}

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