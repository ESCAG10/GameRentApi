package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Videojuego struct {

	ID                 bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo             string        `bson:"titulo" json:"titulo"`
	Plataforma         string        `bson:"plataforma" json:"plataforma"`
	Descripcion        string        `bson:"descripcion" json:"descripcion"`
	Desarrollador      string        `bson:"desarrollador" json:"desarrollador"`
	Editor             string        `bson:"editor" json:"editor"`
	PrecioRenta        float64       `bson:"precioRenta" json:"precioRenta"`
	Stock              int           `bson:"stock" json:"stock"`
	Activo             bool          `bson:"activo" json:"activo"`
	CategoriaID        bson.ObjectID `bson:"categoriaId" json:"categoriaId"`
	FechaCreacion      time.Time     `bson:"fechaCreacion" json:"fechaCreacion"`
	FechaActualizacion time.Time     `bson:"fechaActualizacion" json:"fechaActualizacion"`

}