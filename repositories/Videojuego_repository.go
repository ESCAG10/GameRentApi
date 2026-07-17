package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type VideojuegoRepository struct {
    Collection *mongo.Collection
}

func NewVideojuegoRepository(collection *mongo.Collection) *VideojuegoRepository {
    return &VideojuegoRepository{Collection: collection}
}

func (r *VideojuegoRepository) Create(videojuego *models.Videojuego) error {
    videojuego.FechaCreacion = time.Now()
    videojuego.FechaActualizacion = time.Now()
    _, err := r.Collection.InsertOne(context.Background(), videojuego)
    return err
}

func (r *VideojuegoRepository) FindAll() ([]*models.Videojuego, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var videojuegos []*models.Videojuego
    if err = cursor.All(context.Background(), &videojuegos); err != nil {
        return nil, err
    }
    return videojuegos, nil
}

func (r *VideojuegoRepository) FindByID(id bson.ObjectID) (*models.Videojuego, error) {
    var videojuego models.Videojuego
    err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&videojuego)
    if err != nil {
        return nil, err
    }
    return &videojuego, nil
}

func (r *VideojuegoRepository) Update(id bson.ObjectID, videojuego *models.Videojuego) error {
    videojuego.FechaActualizacion = time.Now()
    _, err := r.Collection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.M{"$set": videojuego},
    )
    return err
}

func (r *VideojuegoRepository) Delete(id bson.ObjectID) error {
    _, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}