package seeders

import (
	"fmt"
	"math/rand"
	"move/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Seeder) measurementsSeeder() error {
	// drop the measurements collection if it exists
	measurementCollection := s.mongoClient.Database("move").Collection("measurements")
	if err := measurementCollection.Drop(s.ctx); err != nil {
		return fmt.Errorf("failed to drop measurements collection: %w", err)
	}

	const numMeasurements = 100

	// get all the sensors
	var sensors []models.Sensor
	collection := s.mongoClient.Database("move").Collection("sensors")
	cursor, err := collection.Find(s.ctx, primitive.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(s.ctx)

	if err := cursor.All(s.ctx, &sensors); err != nil {
		return fmt.Errorf("failed to retrieve sensors: %w", err)
	}

	// create measurements for each sensor
	var measurements []models.Measurement
	for _, sensor := range sensors {
		for i := 0; i < numMeasurements; i++ {
			// Pick a random time between 30 days ago and now
			start := time.Now().Add(-30 * 24 * time.Hour)
			end := time.Now()
			delta := end.Unix() - start.Unix()
			randomSec := rand.Int63n(delta + 1) // +1 to include 'end' as a possible value
			randomTimestamp := start.Add(time.Duration(randomSec) * time.Second)

			// Create a measurement
			measurement := models.Measurement{
				SensorID:  sensor.ID,
				Value:     rand.Float64() * sensor.Threshold,
				Timestamp: randomTimestamp,
				Metric:    "dispmm", // assuming dispmm is the metric for all measurements
			}
			measurements = append(measurements, measurement)
		}
	}

	// Insert measurements into the database
	measurementCollection = s.mongoClient.Database("move").Collection("measurements")
	_, err = measurementCollection.InsertMany(s.ctx, measurements)
	if err != nil {
		return err
	}

	return nil
}
