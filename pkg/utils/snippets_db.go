package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterId(id string) (primitive.M, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.M{}, err
	}

	filter := primitive.M{"_id": oid}
	return filter, nil
}
