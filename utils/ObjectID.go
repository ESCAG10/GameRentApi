package utils

import "go.mongodb.org/mongo-driver/v2/bson"

func ValidarObjectID(id string) bool {

	_, err := bson.ObjectIDFromHex(id)

	return err == nil
}