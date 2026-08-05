package repositories

import (
	"context"
	"errors"
	"gamerentapi/models"

	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RentaRepository struct {
	Collection *mongo.Collection
}

func NewRentaRepository(collection *mongo.Collection) *RentaRepository {
	return &RentaRepository{
		Collection: collection,
	}
}

// Crear renta
func (r *RentaRepository) Create(renta *models.Renta) error {

	ahora := time.Now()

	if renta.FechaRenta.IsZero() {
		renta.FechaRenta = ahora
	}

	if renta.FechaEntrega.IsZero() {
		renta.FechaEntrega = ahora.AddDate(0, 0, renta.PeriodoRenta)
	}

	renta.FechaCreacion = ahora
	renta.FechaActualizacion = ahora

	_, err := r.Collection.InsertOne(
		context.Background(),
		renta,
	)

	return err
}

// Obtener todas las rentas
func (r *RentaRepository) FindAll() ([]models.Renta, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var rentas []models.Renta

	err = cursor.All(
		context.Background(),
		&rentas,
	)

	return rentas, err
}

// Obtener renta por ID
func (r *RentaRepository) FindByID(id string) (*models.Renta, error) {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var renta models.Renta

	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&renta)

	if err != nil {
		return nil, err
	}

	return &renta, nil
}

// Actualizar renta
func (r *RentaRepository) Update(id string, renta *models.Renta,) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	existente, err := r.FindByID(id)
	
	if err != nil {
		return err
	}
	
	renta.FechaRenta = existente.FechaRenta
	renta.FechaCreacion = existente.FechaCreacion
	renta.FechaActualizacion = time.Now()

	_, err = r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
		bson.M{
			"$set": bson.M{
				"usuarioId":          renta.UsuarioID,
				"videojuegoId":       renta.VideojuegoID,
				"categoriaId":        renta.CategoriaID,
				"fechaRenta":         renta.FechaRenta,
				"periodoRenta":       renta.PeriodoRenta,
				"fechaEntrega":       renta.FechaEntrega,
				"estado":             renta.Estado,
				"activo":             renta.Activo,
				"fechaCreacion":      renta.FechaCreacion,
				"fechaActualizacion": renta.FechaActualizacion,
			},
		},
	)

	return err
}

func (r *RentaRepository) Devolver(id string) error {

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
				"estado": "Devuelta",
				"fechaActualizacion": time.Now(),
			},
		},
	)

	return err
}

// Eliminar renta
func (r *RentaRepository) Delete(id string) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	result, err := r.Collection.DeleteOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("renta no encontrada")
	}

	return nil
}

func (r *RentaRepository) ExisteRentaActiva(usuarioID, videojuegoID bson.ObjectID) (bool, error) {

	var renta models.Renta

	err := r.Collection.FindOne(
		context.Background(),
		bson.M{
			"usuarioId":   usuarioID,
			"videojuegoId": videojuegoID,
			"estado":      "Activo",
		},
	).Decode(&renta)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil 

}

func (r *RentaRepository) FindByUsuarioID(id string) ([]models.Renta, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{
			"usuarioId": objectID,
		},
	)

	if err != nil {
		return nil, err
	}

	var rentas []models.Renta

	err = cursor.All(
		context.Background(),
		&rentas,
	)

	return rentas, err
}

func (r *VideojuegoRepository) IncrementarStock(id bson.ObjectID) error {

	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
		bson.M{
			"$inc": bson.M{
				"stock": 1,
			},
		},
	)

	return err
}