package seeders

import (
	"fmt"
	"math/rand"
	"move/internal/models"
	"time"
)

func (s *Seeder) sensorsSeeder() error {
	// drop the sensors collection if it exists
	collection := s.mongoClient.Database("move").Collection("sensors")
	if err := collection.Drop(s.ctx); err != nil {
		return fmt.Errorf("failed to drop sensors collection: %w", err)
	}

	var sensors []models.Sensor

	for i := 0; i <= 10; i++ {
		// Create random name of kind "Sensor_{random_number}"
		name := fmt.Sprintf("Sensor_%d", rand.Intn(1000))
		randomCoordinates := []float64{
			rand.Float64()*(18.0-6.0) + 6.0,   // Longitude: 6.0 to 18.0
			rand.Float64()*(47.0-36.0) + 36.0, // Latitude: 36.0 to 47.0
		}
		threshold := rand.Float64() * 30.0 // Random threshold between 0 and 100

		sensor := models.Sensor{
			Name:      name,
			Location:  randomCoordinates,
			Threshold: threshold,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		sensors = append(sensors, sensor)
	}

	// Insert sensors into the database
	collection = s.mongoClient.Database("move").Collection("sensors")
	_, err := collection.InsertMany(s.ctx, sensors)

	return err
}
