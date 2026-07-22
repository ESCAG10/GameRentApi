package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Usuario struct {
    ID              bson.ObjectID `bson:"_id,omitempty" json:"id"`
    Nombre          string        `bson:"nombre" json:"nombre"`
    Correo          string        `bson:"correo" json:"correo"`
    FechaRegistro   string     `bson:"fechaRegistro" json:"fechaRegistro"`
    PasswordHash    string        `bson:"passwordHash" json:"-"`
    Rol             string        `bson:"rol" json:"rol"`
    Activo          bool          `bson:"activo" json:"activo"`
}