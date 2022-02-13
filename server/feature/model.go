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

// setTimestamps sets the createdAt and updatedAt fields to the current time.
func (feature *featureModel) setTimestamps() {
	feature.CreatedAt = time.Now()
	feature.UpdatedAt = feature.CreatedAt
}

// setUpdatedAt sets the updatedAt field to the current time.
func (feature *featureModel) setUpdatedAt() {
	feature.UpdatedAt = time.Now()
}
