package mgutil

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCollection(cli *mongo.Client, db, collection string) *mongo.Collection {
	return cli.Database(db).Collection(collection)
}
