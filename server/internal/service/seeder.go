package service

import (
	"move/internal/seeders"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SeederService struct {
	router      *gin.RouterGroup
	mongoClient *mongo.Client
}

func NewSeederService(router *gin.RouterGroup, mongoClient *mongo.Client) *SeederService {
	// Get the MongoDB client from the router context

	s := SeederService{
		router:      router,
		mongoClient: mongoClient,
	}

	// Initialize routes or other service-specific setup here
	s.start()

	return &s
}

func (s *SeederService) start() {
	group := s.router.Group("/seeders")

	group.GET("", func(c *gin.Context) {
		seeders.NewSeeder(s.mongoClient)

		c.JSON(200, gin.H{
			"message": "Database seeded successfully",
		})
	})
}
