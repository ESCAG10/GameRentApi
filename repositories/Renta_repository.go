package repositories

import (
	"context"
	"errors"
	"gamerentapi/config"
	"gamerentapi/models"

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

	existe, err := r.ExisteRentaActiva(
		renta.UsuarioID,
		renta.VideojuegoID,
	)

	if err != nil {
		return err
	}

	if existe {
		return errors.New("ya tienes este videojuego rentado")
	}

	_, err = r.Collection.InsertOne(
		context.Background(),
		renta,
	)

	if err != nil {
		return err
	}

	_, err = config.Database.Collection("videojuego").UpdateOne(
		context.Background(),
		bson.M{
			"_id": renta.VideojuegoID,
		},
		bson.M{
			"$inc": bson.M{
				"stock": -1,
			},
		},
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

	if err != nil {
		return nil, err
	}

	usuariosCollection := config.Database.Collection("usuario")

	videojuegosCollection := config.Database.Collection("videojuego")

	for i := range rentas {

		// Buscar nombre del usuario

		var usuario models.Usuario

		err := usuariosCollection.FindOne(
			context.Background(),
			bson.M{
				"_id": rentas[i].UsuarioID,
			},
		).Decode(&usuario)

		if err == nil {

			rentas[i].NombreUsuario = usuario.Nombre

		}

		// Buscar nombre del videojuego

		var videojuego models.Videojuego

		err = videojuegosCollection.FindOne(
			context.Background(),
			bson.M{
				"_id": rentas[i].VideojuegoID,
			},
		).Decode(&videojuego)

		if err == nil {

			rentas[i].NombreVideojuego = videojuego.Titulo

		}

	}

	return rentas, nil
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
func (r *RentaRepository) Update(id string, renta *models.Renta) error {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	existente, err := r.FindByID(id)

	if err != nil {
		return err
	}

	renta.FechaCreacion = existente.FechaCreacion

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

// Eliminar renta y regresar stock
func (r *RentaRepository) Delete(id string) error {

objectID, err := bson.ObjectIDFromHex(id)

if err != nil {
return err
}


// Buscar la renta antes de eliminarla
var renta models.Renta

err = r.Collection.FindOne(
context.Background(),
bson.M{
"_id": objectID,
},
).Decode(&renta)


if err != nil {
return errors.New("renta no encontrada")
}



// Aumentar stock del videojuego
_, err = config.Database.Collection("videojuego").UpdateOne(
context.Background(),
bson.M{
"_id": renta.VideojuegoID,
},
bson.M{
"$inc": bson.M{
"stock": 1,
},
},
)


if err != nil {
return err
}



// Eliminar renta
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

return errors.New("no se pudo eliminar la renta")

}


return nil
}

// Verificar si el usuario ya tiene ese videojuego rentado
func (r *RentaRepository) ExisteRentaActiva(usuarioID, videojuegoID bson.ObjectID) (bool, error) {

	var renta models.Renta

	err := r.Collection.FindOne(
		context.Background(),
		bson.M{

			"usuarioId":    usuarioID,
			"videojuegoId": videojuegoID,
			"estado":       "Activo",
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

// Obtener rentas por usuario
func (r *RentaRepository) FindByUsuarioID(usuarioID string) ([]models.Renta, error) {

	objectID, err := bson.ObjectIDFromHex(usuarioID)

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
