package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Usuario struct {
	ID       bson.ObjectID `bson:"id" json:"_id,omitempty"`
	Nombre   string `bson:"nombre" json:"nombre"`
	Correo   string `bson:"correo" json:"correo"`
	FechaRegistro string `bson:"fecha_registro" json:"fecha_registro"`
}