package feature

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// featureModel represents a feature object. Both in the database and in the API.
type featureModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Data      string             `json:"data,omitempty" binding:"required" validate:"custom-validation" bson:"data,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// newFeature creates a new feature object with the given data and current timestamps.
func newFeature(data string) *featureModel {
	feature := featureModel{Data: data}
	feature.CreatedAt = time.Now()
	feature.UpdatedAt = feature.CreatedAt
	return &feature
}

// newFeatureFromPartial creates a new feature object from the given partial feature object.
func newFeatureFromPartial(feature featureModel) *featureModel {
	feature.CreatedAt = time.Now()
	feature.UpdatedAt = feature.CreatedAt
	return &feature
}

// newFeatureUpdate creates a new feature update object with the given data and current updatedAt timestamp.
func newFeatureUpdate(data string) *featureModel {
	feature := featureModel{
		Data:      data,
		UpdatedAt: time.Now(),
	}
	return &feature
}

//newFeatureUpdateFromPartial creates a new feature update object from the given partial feature object with the current updatedAt timestamp.
func newFeatureUpdateFromPartial(feature featureModel) *featureModel {
	feature.UpdatedAt = time.Now()
	return &feature
}
