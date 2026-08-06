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

// Crear categoría
func (r *CategoriaRepository) Create(categoria *models.Categoria) error {

	//ahora := time.Now()

	//categoria.FechaCreacion = ahora
	//categoria.FechaActualizacion = ahora

	_, err := r.Collection.InsertOne(
		context.Background(),
		categoria,
	)

	return err
}

// Obtener todas las categorías
func (r *CategoriaRepository) FindAll() ([]models.Categoria, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var categorias []models.Categoria

	err = cursor.All(
		context.Background(),
		&categorias,
	)

	return categorias, err
}

// Obtener categoría por ID
func (r *CategoriaRepository) FindByID(id string) (*models.Categoria, error) {

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

// Actualizar categoría
func (r *CategoriaRepository) Update(id string, categoria *models.Categoria,) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	existente, err := r.FindByID(id)

	if err != nil {
		return err
	}

	// Mantener la fecha de creación
	categoria.FechaCreacion = existente.FechaCreacion

	// Actualizar la fecha de modificación
	//categoria.FechaActualizacion = time.Now()

	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{
				"nombre":             categoria.Nombre,
				"descripcion":        categoria.Descripcion,
				"activo":             categoria.Activo,
				"fechaCreacion":      categoria.FechaCreacion,
				"fechaActualizacion": categoria.FechaActualizacion,
			},
		},
	)

	return err
}

// Eliminar categoría
func (r *CategoriaRepository) Delete(id string) error {

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