package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Videojuego struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo       string        `bson:"titulo" json:"titulo"`
	Plataforma   string        `bson:"plataforma" json:"plataforma"`
	PrecioRenta  float64       `bson:"precioRenta" json:"precioRenta"`
	Stock        int           `bson:"stock" json:"stock"`
	Activo       bool          `bson:"activo" json:"activo"`
	CategoriaID  bson.ObjectID `bson:"categoriaId" json:"categoriaId"`
}