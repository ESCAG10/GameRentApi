package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Perfil struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string        `bson:"nombre" json:"nombre"`
	Direccion   string        `bson:"direccion" json:"direccion"`
	Telefono    string        `bson:"telefono" json:"telefono"`
	NivelCliente string        `bson:"nivelCliente" json:"nivelCliente"`
}