package feature

import (
	"context"

	"github.com/MichaelSimkin/go-template/database"
	"github.com/MichaelSimkin/go-template/server/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// This exported Repository variable is used to access the repository methods.
var Repository = &repository{}

// The repository methods are defined on this struct.
type repository struct{}

func (repository) createDocument(ctx context.Context, document FeatureModel) (*FeatureModel, error) {
	result, err := database.FeatureCollection.InsertOne(ctx, document)
	if mongo.IsDuplicateKeyError(err) {
		return nil, errors.DuplicateKeyError
	}
	if err != nil {
		return nil, errors.NewUnknownMongoError(err)
	}

	// return the created document with the ID returned by MongoDB
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.InvalidMongoIDError
	}
	document.ID = insertedID

	return &document, nil
}

func (repository) getDocuments(ctx context.Context, filters string) ([]*FeatureModel, error) {
	// Create the filters object
	var searchFilters = bson.M{}
	err := bson.UnmarshalExtJSON([]byte(filters), true, &searchFilters)
	if err != nil {
		return nil, errors.NewInvalidFiltersError(err)
	}

	// Find the documents in the collection
	cursor, err := database.FeatureCollection.Find(ctx, searchFilters)
	if err == mongo.ErrNoDocuments {
		return []*FeatureModel{}, nil
	}
	if err != nil {
		return nil, errors.NewUnknownMongoError(err)
	}

	defer cursor.Close(ctx)

	return decodeMongoDocuments(ctx, cursor)
}

// decodeMongoDocuments decodes the documents returned by the MongoDB cursor.
func decodeMongoDocuments(ctx context.Context, cursor *mongo.Cursor) ([]*FeatureModel, error) {
	documents := []*FeatureModel{}

	for cursor.Next(ctx) {
		var document = &FeatureModel{}
		err := cursor.Decode(document)
		if err != nil {
			return nil, errors.NewFailedToDecodeError(err)
		}
		documents = append(documents, document)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.NewMongoCursorError(err)
	}

	return documents, nil
}
