package data

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB struct {
	database    string
	collection  string
	mongoClient *mongo.Client
}

func (m *MongoDB) Database() *mongo.Database {
	return m.mongoClient.Database(m.database)
}

func (m *MongoDB) Collection() *mongo.Collection {
	return m.Database().Collection(m.collection)
}
