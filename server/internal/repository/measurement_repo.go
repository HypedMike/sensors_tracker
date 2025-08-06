package repository

import (
	"context"
	"move/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MeasurementRepository struct {
	mongoClient    *mongo.Client
	collectionName string
	collection     *mongo.Collection
	ctx            context.Context
}

func NewMeasurementRepository(mongoClient *mongo.Client) *MeasurementRepository {
	return &MeasurementRepository{
		mongoClient:    mongoClient,
		collectionName: "measurements",
		collection:     mongoClient.Database(database).Collection("measurements"),
		ctx:            context.Background(),
	}
}

func (r *MeasurementRepository) GetMeasurementsBySensorID(sensorID string, filter struct {
	From time.Time
	To   time.Time
}) ([]models.Measurement, error) {
	var measurements []models.Measurement

	// convert sensorID to ObjectID
	id, err := bson.ObjectIDFromHex(sensorID)
	if err != nil {
		return nil, err
	}

	// create a filter for the query and sort by timestamp
	query := bson.M{"sensor_id": id}
	if !filter.From.IsZero() || !filter.To.IsZero() {
		if !filter.From.IsZero() && !filter.To.IsZero() {
			query["timestamp"] = bson.M{"$gte": filter.From, "$lte": filter.To}
		} else {
			if !filter.From.IsZero() {
				query["timestamp"] = bson.M{"$gte": filter.From}
			} else if !filter.To.IsZero() {
				query["timestamp"] = bson.M{"$lte": filter.To}
			}
		}
	}

	cursor, err := r.collection.Aggregate(r.ctx, bson.A{
		bson.D{{Key: "$match", Value: query}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "timestamp", Value: 1}}}},
	})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(r.ctx)

	if err := cursor.All(r.ctx, &measurements); err != nil {
		return nil, err
	}

	return measurements, nil
}

// Add measurements
func (r *MeasurementRepository) AddMeasurements(measurement []models.Measurement) error {
	for i, m := range measurement {
		// set metric to dispmm
		m.Metric = "dispmm"
		measurement[i] = m
	}

	// insert measurements into the collection
	_, err := r.collection.InsertMany(r.ctx, measurement)
	return err
}

// Add values as measurements
func (r *MeasurementRepository) AddValues(sensorID string, values []float64) error {
	var measurements []models.Measurement

	// convert sensorID to ObjectID
	id, err := bson.ObjectIDFromHex(sensorID)
	if err != nil {
		return err
	}

	for _, value := range values {
		measurements = append(measurements, models.Measurement{
			SensorID:  id,
			Value:     value,
			Metric:    "dispmm",
			Timestamp: time.Now(),
		})
	}

	return r.AddMeasurements(measurements)
}

// Delete measurements by sensor ID
func (r *MeasurementRepository) DeleteMeasurementsBySensorID(sensorID string) error {
	// convert sensorID to ObjectID
	id, err := bson.ObjectIDFromHex(sensorID)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteMany(r.ctx, bson.M{"sensor_id": id})
	return err
}

func (r *MeasurementRepository) StatsBySensorID(sensorID string) (models.Sensor, error) {
	// convert sensorID to ObjectID
	id, err := bson.ObjectIDFromHex(sensorID)
	if err != nil {
		return models.Sensor{}, err
	}

	// aggregate to get stats
	now := time.Now()
	threeDaysAgo := now.AddDate(0, 0, -3)
	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.M{
			"sensor_id": id,
			"timestamp": bson.M{
				"$gte": threeDaysAgo,
				"$lte": now,
			},
		}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":     nil,
			"average": bson.M{"$avg": "$value"},
			"min":     bson.M{"$min": "$value"},
			"max":     bson.M{"$max": "$value"},
		}}},
	}

	cursor, err := r.collection.Aggregate(r.ctx, pipeline)
	if err != nil {
		return models.Sensor{}, err
	}
	defer cursor.Close(r.ctx)

	var result []struct {
		Average float64 `bson:"average"`
		Min     float64 `bson:"min"`
		Max     float64 `bson:"max"`
	}

	if err := cursor.All(r.ctx, &result); err != nil {
		return models.Sensor{}, err
	}

	if len(result) == 0 {
		return models.Sensor{}, nil // no measurements found
	}

	stats := result[0]

	return models.Sensor{
		ID:      id,
		Average: stats.Average,
		Min:     stats.Min,
		Max:     stats.Max,
	}, nil
}
