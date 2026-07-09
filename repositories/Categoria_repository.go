package repositories

import (
	"context"
	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoriaRepository struct {
	Collection *mongo.Collection
}

func NewCategoriaRepository(collection *mongo.Collection) *CategoriaRepository {
	return &CategoriaRepository{
		Collection: collection,
	}
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) error {
	_, err := r.Collection.InsertOne(context.Background(), categoria)
	return err
}

func (r *CategoriaRepository) FindAll() ([]*models.Categoria, error) {
	cursor, err := r.Collection.Find(context.Background(), nil)
	if err != nil {
		return []*models.Categoria{}, err
	}
	defer cursor.Close(context.Background())

	var categorias []*models.Categoria
	if err = cursor.All(context.Background(), &categorias); err != nil {
		return []*models.Categoria{}, err
	}

	return categorias, nil
}

func (r *CategoriaRepository) FindByID(id string,) (*models.Categoria, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var categoria models.Categoria

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&categoria)

	if err != nil {
		return nil, err
	}

	return &categoria, nil
}

func (r *CategoriaRepository) Update(id string, categoria *models.Categoria) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Collection.UpdateByID(context.Background(), objectID, bson.M{"$set": categoria})
	return err
}

func (r *CategoriaRepository) Delete(id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}