package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Categoria struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string        `bson:"nombre" json:"nombre"`
	Descripcion string        `bson:"descripcion" json:"descripcion"`
}