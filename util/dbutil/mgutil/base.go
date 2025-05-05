package mgutil

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne inserts a single document into the specified collection.
// This function IGNORES the result of the insert operation.
func InsertOne(ctx context.Context, collection *mongo.Collection, data any, opts ...*options.InsertOneOptions) error {
	_, err := collection.InsertOne(ctx, data, opts...)
	if err != nil {
		return err
	}
	return nil
}

// FindOne finds one document in the specified collection that matches the filter.
// result is a Pointer to the struct that will hold the result.
func FindOne(ctx context.Context, collection *mongo.Collection, filter any, result any, opts ...*options.FindOneOptions) (any, error) {
	res := collection.FindOne(ctx, filter, opts...)
	if err := res.Err(); err != nil {
		return nil, err
	}

	err := res.Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteOne deletes one document in the specified collection that matches the filter.
func DeleteOne(ctx context.Context, collection *mongo.Collection, filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return collection.DeleteOne(ctx, filter, opts...)
}

// UpdateOne finds one document in the specified collection that matches the filter.
// ! NOTE: the parameter `data` MUST contain UPDATE OPERATOR(beginning with '$') to update the document.UpdateOne WILL NOT add it.
func UpdateOne(ctx context.Context, collection *mongo.Collection, filter any, data any, opts ...*options.UpdateOptions) error {
	_, err := collection.UpdateOne(ctx, filter, data, opts...)
	if err != nil {
		return err
	}

	return nil
}

func NewUpdateOptions() *options.UpdateOptions {
	return &options.UpdateOptions{}
}

func NewBsonM(kvPairs ...interface{}) *bson.M {
	ret := bson.M{}
	if len(kvPairs)%2 != 0 {
		panic("NewBsonM: arguments must be in key-value pairs")
	}

	for i := 0; i < len(kvPairs); i += 2 {
		key, ok := kvPairs[i].(string)
		if !ok {
			panic("NewBsonM: key must be a string")
		}
		ret[key] = kvPairs[i+1]
	}
	return &ret
}

func NewBsonD(kvPairs ...interface{}) *bson.D {
	ret := bson.D{}
	if len(kvPairs)%2 != 0 {
		panic("NewBsonD: arguments must be in key-value pairs")
	}

	for i := 0; i < len(kvPairs); i += 2 {
		key, ok := kvPairs[i].(string)
		if !ok {
			panic("NewBsonD: key must be a string")
		}
		ret = append(ret, bson.E{Key: key, Value: kvPairs[i+1]})
	}
	return &ret
}

func NewFilter(kvPairs ...interface{}) any {
	return NewBsonM(kvPairs...)
}
