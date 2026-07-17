package repositories

import (
	"context"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoriaRepository struct {
    Collection *mongo.Collection
}

func NewCategoriaRepository(collection *mongo.Collection) *CategoriaRepository {
    return &CategoriaRepository{Collection: collection}
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) error {
    categoria.FechaCreacion = time.Now()
    categoria.FechaActualizacion = time.Now()
    _, err := r.Collection.InsertOne(context.Background(), categoria)
    return err
}

func (r *CategoriaRepository) FindAll() ([]*models.Categoria, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var categorias []*models.Categoria
    if err = cursor.All(context.Background(), &categorias); err != nil {
        return nil, err
    }
    return categorias, nil
}

func (r *CategoriaRepository) FindByID(id bson.ObjectID) (*models.Categoria, error) {
    var categoria models.Categoria
    err := r.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&categoria)
    if err != nil {
        return nil, err
    }
    return &categoria, nil
}

func (r *CategoriaRepository) Update(id bson.ObjectID, categoria *models.Categoria) error {
    categoria.FechaActualizacion = time.Now()
    _, err := r.Collection.UpdateOne(
        context.Background(),
        bson.M{"_id": id},
        bson.M{"$set": categoria},
    )
    return err
}

func (r *CategoriaRepository) Delete(id bson.ObjectID) error {
    _, err := r.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
    return err
}