package seeders

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Seeder struct {
	mongoClient *mongo.Client
	ctx         context.Context
}

func NewSeeder(client *mongo.Client) *Seeder {
	s := &Seeder{
		mongoClient: client,
		ctx:         context.Background(),
	}

	// Run the sensors seeder
	if err := s.sensorsSeeder(); err != nil {
		panic(err)
	}

	// Run the measurements seeder
	if err := s.measurementsSeeder(); err != nil {
		panic(err)
	}

	return s
}
