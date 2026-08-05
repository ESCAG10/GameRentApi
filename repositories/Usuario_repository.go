package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gamerentapi/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioRepository struct {
	Collection *mongo.Collection
}

func NewUsuarioRepository(collection *mongo.Collection) *UsuarioRepository {
	return &UsuarioRepository{
		Collection: collection,
	}
}

// Crear usuario
func (r *UsuarioRepository) Create(usuario *models.Usuario) error {

    ahora := time.Now()

    usuario.FechaRegistro = ahora
    usuario.FechaCreacion = ahora
    usuario.FechaActualizacion = ahora


	res, err := r.Collection.InsertOne(context.Background(), usuario)
	if err != nil {
    return err
}

// Aquí res.InsertedID es el ObjectID generado por MongoDB
fmt.Println("Nuevo ID:", res.InsertedID)

usuario.ID = res.InsertedID.(bson.ObjectID)
//usuario.IDString = usuario.ID.Hex() // convierte ObjectID a string

    return err
}


// Obtener todos los usuarios
func (r *UsuarioRepository) FindAll() ([]models.Usuario, error) {

	cursor, err := r.Collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var usuarios []models.Usuario

	err = cursor.All(
		context.Background(),
		&usuarios,
	)

	return usuarios, err
}


// Obtener usuario por ID
func (r *UsuarioRepository) FindByID(id string) (*models.Usuario, error) {

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}


	var usuario models.Usuario


	err = r.Collection.FindOne(
		context.Background(),
		bson.M{
			"_id": objectID,
		},
	).Decode(&usuario)


	if err != nil {
		return nil, err
	}


	return &usuario, nil
}


// Actualizar usuario
func (r *UsuarioRepository) Update(id string,usuario *models.Usuario,) error {


	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}


	// Recuperar datos actuales
	existente, err := r.FindByID(id)

	if err != nil {
		return err
	}


	// Mantener fechas originales
	usuario.FechaRegistro = existente.FechaRegistro
	usuario.FechaCreacion = existente.FechaCreacion
	
	// Actualizar únicamente la fecha de modificación
	usuario.FechaActualizacion = time.Now()


	_, err = r.Collection.UpdateOne(
	context.Background(),
	bson.M{
		"_id": objectID,
	},
	bson.M{
		"$set": bson.M{
			"nombre":              usuario.Nombre,
			"correo":              usuario.Correo,
			"password":            usuario.Password,
			"rol":                 usuario.Rol,
			"activo":              usuario.Activo,
			"fechaRegistro":       usuario.FechaRegistro,
			"fechaCreacion":       usuario.FechaCreacion,
			"fechaActualizacion":  usuario.FechaActualizacion,
		},
	},
)


	return err
}


// Eliminar usuario
func (r *UsuarioRepository) Delete(id string) error {


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

// Login
func (r *UsuarioRepository) Login(correo, password string) (*models.Usuario, error) {

var usuario models.Usuario

err := r.Collection.FindOne(
context.Background(),
bson.M{
"correo": correo,
"password": password,
},
).Decode(&usuario)

if err != nil {
return nil, err
}

if usuario.Password != password{
	return nil, errors.New("Contraseña incorrecta")}

return &usuario, nil
}

func (r *UsuarioRepository) FindByCorreo(correo string) (*models.Usuario, error) {

	var usuario models.Usuario

	err := r.Collection.FindOne(
		context.Background(),
		bson.M{
			"correo": correo,
		},
	).Decode(&usuario)

	if err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (r *UsuarioRepository) ExisteCorreo(correo string) (bool, error) {

	count, err := r.Collection.CountDocuments(
		context.Background(),
		bson.M{
			"correo": correo,
		},
	)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}