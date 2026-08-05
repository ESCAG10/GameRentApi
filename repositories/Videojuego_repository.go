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
	return &VideojuegoRepository{
		Collection: collection,
	}
}

// Crear videojuego
func (r *VideojuegoRepository) Create(videojuego *models.Videojuego) error {

	ahora := time.Now()

	videojuego.FechaCreacion = ahora
	videojuego.FechaActualizacion = ahora

	_, err := r.Collection.InsertOne(
		context.Background(),
		videojuego,
	)

	return err
}

// Obtener todos los videojuegos
func (r *VideojuegoRepository) FindAll() ([]models.Videojuego, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var videojuegos []models.Videojuego

	err = cursor.All(
		context.Background(),
		&videojuegos,
	)

	return videojuegos, err
}


func (r *VideojuegoRepository) FindActivos() ([]models.Videojuego, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{
			"activo": true,
		},
	)

	if err != nil {
		return nil, err
	}

	var videojuegos []models.Videojuego

	err = cursor.All(
		context.Background(),
		&videojuegos,
	)

	return videojuegos, err
}


// Obtener videojuego por ID
func (r *VideojuegoRepository) FindByID(id string) (*models.Videojuego, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var videojuego models.Videojuego

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&videojuego)

	if err != nil {
		return nil, err
	}

	return &videojuego, nil
}

// Actualizar videojuego
func (r *VideojuegoRepository) Update(id string,videojuego *models.Videojuego,) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	existente, err := r.FindByID(id)

	if err != nil {
		return err
	}

	videojuego.FechaCreacion = existente.FechaCreacion
	videojuego.FechaActualizacion = time.Now()

	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{
				"titulo":              videojuego.Titulo,
				"plataforma":          videojuego.Plataforma,
				"descripcion":         videojuego.Descripcion,
				"precioRenta":         videojuego.PrecioRenta,
				"stock":               videojuego.Stock,
				"activo":              videojuego.Activo,
				"categoriaId":         videojuego.CategoriaID,
				"desarrollador":       videojuego.Desarrollador,
				"editor":              videojuego.Editor,
				"fechaCreacion":       videojuego.FechaCreacion,
				"fechaActualizacion":  videojuego.FechaActualizacion,
			},
		},
	)

	return err
}

// Eliminar videojuego
func (r *VideojuegoRepository) Delete(id string) error {

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

func (r *VideojuegoRepository) DecreaseStock(id string) error {

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
			"$inc": bson.M{
				"stock": -1,
			},
		},
	)

	return err
}