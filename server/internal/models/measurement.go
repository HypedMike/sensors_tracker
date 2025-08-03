package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Measurement struct {
	ID        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	SensorID  bson.ObjectID `json:"sensor_id" bson:"sensor_id"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	Value     float64       `json:"value" bson:"value"`
	Metric    string        `json:"metric" bson:"metric"`
}
