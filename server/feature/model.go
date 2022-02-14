package feature

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FeatureModel represents a feature object. Both in the database and in the API.
type FeatureModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Data      string             `json:"data,omitempty" binding:"required" validate:"custom-validation" bson:"data,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// setTimestamps sets the createdAt and updatedAt fields to the current time.
func (feature *FeatureModel) setTimestamps() {
	feature.CreatedAt = time.Now()
	feature.UpdatedAt = feature.CreatedAt
}

// setUpdatedAt sets the updatedAt field to the current time.
func (feature *FeatureModel) setUpdatedAt() {
	feature.UpdatedAt = time.Now()
}
