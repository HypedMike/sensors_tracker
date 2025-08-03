package seeders

import (
	"fmt"
	"move/internal/models"
	"move/internal/repository"
)

func (s *Seeder) usersSeeder() error {
	// Initialize the user repository
	userRepo := repository.NewUserRepository(s.mongoClient)

	// Drop the users collection if it exists
	userCollection := s.mongoClient.Database("move").Collection("users")
	if err := userCollection.Drop(s.ctx); err != nil {
		return fmt.Errorf("failed to drop users collection: %w", err)
	}

	// Create a sample user
	user := models.User{
		Username: "admin",
		Password: "admin",
	}

	// Insert the sample user into the database
	if err := userRepo.CreateUser(user); err != nil {
		return fmt.Errorf("error seeding user: %w", err)
	}

	return nil
}
