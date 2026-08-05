package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Renta struct {

	ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID          bson.ObjectID `bson:"usuarioId" json:"usuarioId"`
	VideojuegoID       bson.ObjectID `bson:"videojuegoId" json:"videojuegoId"`
	CategoriaID        bson.ObjectID `bson:"categoriaId" json:"categoriaId"`
	FechaRenta         time.Time `bson:"fechaRenta" json:"fechaRenta"`
	PeriodoRenta       int       `bson:"periodoRenta" json:"periodoRenta"`
	FechaEntrega       time.Time `bson:"fechaEntrega" json:"fechaEntrega"`
	Estado             string `bson:"estado" json:"estado"`
	Activo             bool   `bson:"activo" json:"activo"`
	FechaCreacion      time.Time `bson:"fechaCreacion" json:"fechaCreacion"`
	FechaActualizacion time.Time `bson:"fechaActualizacion" json:"fechaActualizacion"`
}
