package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Usuario struct {
	ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
	//IDString           string        `json:"id"`
	Nombre             string        `bson:"nombre" json:"nombre"`
	Correo             string        `bson:"correo" json:"correo"`
	Password           string        `bson:"password" json:"password"`
	Rol                string        `bson:"rol" json:"rol"`
	Activo             bool          `bson:"activo" json:"activo"`
	FechaRegistro      time.Time     `bson:"fechaRegistro" json:"fechaRegistro"`
	FechaCreacion      time.Time     `bson:"fechaCreacion" json:"fechaCreacion"`
	FechaActualizacion time.Time     `bson:"fechaActualizacion" json:"fechaActualizacion"`
}