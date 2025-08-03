package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Sensor struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Location  []float64     `json:"location" bson:"location"`
	Threshold float64       `json:"threshold" bson:"threshold"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

	Measurements []Measurement `json:"measurements,omitempty" bson:"measurements,omitempty"`

	// Stats
	Average float64 `json:"average,omitempty" bson:"average,omitempty"`
	Min     float64 `json:"min,omitempty" bson:"min,omitempty"`
	Max     float64 `json:"max,omitempty" bson:"max,omitempty"`
}
