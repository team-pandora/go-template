package database

import (
	"context"
	"fmt"

	"github.com/MichaelSimkin/go-template/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// FeatureCollection is the mongodb collection for feature.
var FeatureCollection *mongo.Collection

// InitMongo initializes the Mongo exported variable.
func InitMongo() {
	// Create mongodb client.
	mongoClient, err := NewMongoClient(config.Mongo.URI)
	if err != nil {
		panic(fmt.Errorf("failed creating mongodb client: %v", err))
	}

	// Get mongodb database.
	db, err := GetMongoDatabase(mongoClient, config.Mongo.URI)
	if err != nil {
		panic(fmt.Errorf("failed getting mongodb database: %v", err))
	}

	initFeatureCollection(db)
}

func initFeatureCollection(db *mongo.Database) {
	FeatureCollection = db.Collection(config.Mongo.FeatureCollectionName)
	CreateMongoCollectionIndex(FeatureCollection, mongo.IndexModel{
		Keys: bson.M{"data": "text"},
	})
	CreateMongoCollectionIndex(FeatureCollection, mongo.IndexModel{
		Keys: bson.M{"updatedAt": -1},
	})
}

// NewMongoClient creates a new mongodb client and connects to mongodb.
func NewMongoClient(connectionString string) (*mongo.Client, error) {
	// Create mongodb client.
	mongoOptions := options.Client().ApplyURI(connectionString)
	mongoClient, err := mongo.NewClient(mongoOptions)
	if err != nil {
		return nil, fmt.Errorf("failed creating mongodb client with connection string \"%s\": %v", connectionString, err)
	}

	// Connect client to mongodb.
	connectionTimeoutCtx, cancelConn := context.WithTimeout(context.Background(), config.Mongo.ConnectionTimeout)
	defer cancelConn()
	err = mongoClient.Connect(connectionTimeoutCtx)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to mongodb with connection string %s: %v", connectionString, err)
	}

	// Check the connection.
	pingTimeoutCtx, cancelPing := context.WithTimeout(context.Background(), config.Mongo.ClientPingTimeout)
	defer cancelPing()
	err = mongoClient.Ping(pingTimeoutCtx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed pinging to mongodb with connection string %s: %v", connectionString, err)
	}

	return mongoClient, nil
}

// GetMongoDatabase gets a mongodb database.
func GetMongoDatabase(mongoClient *mongo.Client, connectionString string) (*mongo.Database, error) {
	connString, err := connstring.ParseAndValidate(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed parsing connection string %s: %v", connectionString, err)
	}

	return mongoClient.Database(connString.Database), nil
}

// CreateMongoCollectionIndex creates a mongodb collection index.
func CreateMongoCollectionIndex(collection *mongo.Collection, indexModel mongo.IndexModel) {
	createIndexTimeoutCtx, cancelCreateIndex := context.WithTimeout(context.Background(), config.Mongo.CreateIndexTimeout)
	defer cancelCreateIndex()
	_, err := collection.Indexes().CreateOne(createIndexTimeoutCtx, indexModel)
	if err != nil {
		panic(fmt.Errorf("failed creating collection index: %v", err))
	}
}
