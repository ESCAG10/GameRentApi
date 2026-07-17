package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Usuario struct {
    ID            bson.ObjectID `bson:"_id,omitempty" json:"id"`
    Nombre        string        `bson:"nombre" json:"nombre"`
    Correo        string        `bson:"correo" json:"correo"`
    FechaRegistro time.Time     `bson:"fecha_registro" json:"fecha_registro"`
    PasswordHash  string        `bson:"passwordHash" json:"-"`   // opcional, para login
    Rol           string        `bson:"rol" json:"rol"`          // opcional, para permisos
    Activo        bool          `bson:"activo" json:"activo"`    // opcional, estado del usuario
}
