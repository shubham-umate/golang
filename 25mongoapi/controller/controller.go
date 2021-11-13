package controller

import "go.mongodb.org/mongo-driver/mongo"

const connectionString = "localhost:27017"
const dbName = "netflix-go"
const colName = "netflix-go"

var collection *mongo.Collection

// connect with mongo

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)
}
