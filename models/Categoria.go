package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Categoria struct {
    ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
    Nombre             string        `bson:"nombre" json:"nombre"`
    Descripcion        string        `bson:"descripcion" json:"descripcion"`
    Activo             bool          `bson:"activo" json:"activo"`
    FechaCreacion      time.Time     `bson:"fechaCreacion" json:"fechaCreacion"`
    FechaActualizacion time.Time     `bson:"fechaActualizacion" json:"fechaActualizacion"`
}
