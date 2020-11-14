package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) isExist(coll, field string, value interface{}) bool {
	collection := r.repo.Collection(coll)
	filter := bson.D{{Key: field, Value: value}}
	var res interface{}
	err := collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return false
	}
	return true
}
