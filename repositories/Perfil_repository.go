package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PerfilRepository struct {
    Collection *mongo.Collection
}

func NewPerfilRepository(collection *mongo.Collection) *PerfilRepository {
    return &PerfilRepository{Collection: collection}
}

func (r *PerfilRepository) Create(perfil *models.Perfil) error {
    perfil.FechaCreacion = time.Now()
    perfil.FechaActualizacion = time.Now()
    _, err := r.Collection.InsertOne(context.Background(), perfil)
    return err
}

func (r *PerfilRepository) FindAll() ([]*models.Perfil, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var perfiles []*models.Perfil
    if err = cursor.All(context.Background(), &perfiles); err != nil {
        return nil, err
    }
    return perfiles, nil
}

func (r *PerfilRepository) FindByID(id bson.ObjectID) (*models.Perfil, error) {
    var perfil models.Perfil
    err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&perfil)
    if err != nil {
        return nil, err
    }
    return &perfil, nil
}

func (r *PerfilRepository) Update(id bson.ObjectID, perfil *models.Perfil) error {
    perfil.FechaActualizacion = time.Now()
    _, err := r.Collection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.M{"$set": perfil},
    )
    return err
}

func (r *PerfilRepository) Delete(id bson.ObjectID) error {
    _, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}