package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Renta struct {

	ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UsuarioID          bson.ObjectID `bson:"usuarioId" json:"usuarioId"`
	VideojuegoID       bson.ObjectID `bson:"videojuegoId" json:"videojuegoId"`
	//CategoriaID        bson.ObjectID `bson:"categoriaId" json:"categoriaId"`
	CategoriaID        string        `bson:"categoriaId" json:"categoriaId"`
	FechaRenta         string `bson:"fechaRenta" json:"fechaRenta"`
	PeriodoRenta       int       `bson:"periodoRenta" json:"periodoRenta"`
	FechaEntrega       string `bson:"fechaEntrega" json:"fechaEntrega"`
	Estado             string `bson:"estado" json:"estado"`
	Activo             bool   `bson:"activo" json:"activo"`
	FechaCreacion      string `bson:"fechaCreacion" json:"fechaCreacion"`
	FechaActualizacion string `bson:"fechaActualizacion" json:"fechaActualizacion"`
	NombreUsuario      string `bson:"nombreUsuario" json:"nombreUsuario"`
	NombreVideojuego   string `bson:"nombreVideojuego" json:"nombreVideojuego"`
}
