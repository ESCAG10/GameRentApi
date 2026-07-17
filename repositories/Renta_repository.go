package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RentaRepository struct {
    Collection *mongo.Collection
}

func NewRentaRepository(collection *mongo.Collection) *RentaRepository {
    return &RentaRepository{Collection: collection}
}

func (r *RentaRepository) Create(renta *models.Renta) error {
    renta.FechaCreacion = time.Now()
    renta.FechaActualizacion = time.Now()
    _, err := r.Collection.InsertOne(context.Background(), renta)
    return err
}

func (r *RentaRepository) FindAll() ([]*models.Renta, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var rentas []*models.Renta
    if err = cursor.All(context.Background(), &rentas); err != nil {
        return nil, err
    }
    return rentas, nil
}

func (r *RentaRepository) FindByID(id bson.ObjectID) (*models.Renta, error) {
    var renta models.Renta
    err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&renta)
    if err != nil {
        return nil, err
    }
    return &renta, nil
}

func (r *RentaRepository) Update(id bson.ObjectID, renta *models.Renta) error {
    renta.FechaActualizacion = time.Now()
    _, err := r.Collection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.M{"$set": renta},
    )
    return err
}

func (r *RentaRepository) Delete(id bson.ObjectID) error {
    _, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}