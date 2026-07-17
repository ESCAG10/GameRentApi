package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Perfil struct {
    ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
    UsuarioID          bson.ObjectID `bson:"usuarioId" json:"usuarioId"`
    Nombre             string        `bson:"nombre" json:"nombre"`
    Direccion          string        `bson:"direccion" json:"direccion"`
    Telefono           string        `bson:"telefono" json:"telefono"`
    NivelCliente       string        `bson:"nivelCliente" json:"nivelCliente"`
    FotoPerfil         string        `bson:"fotoPerfil" json:"fotoPerfil"`
    FechaCreacion      time.Time     `bson:"fechaCreacion" json:"fechaCreacion"`
    FechaActualizacion time.Time     `bson:"fechaActualizacion" json:"fechaActualizacion"`
}
