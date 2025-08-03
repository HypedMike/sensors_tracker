package main

import (
	"move/internal/service"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// Initialize godotenv
	godotenv.Load()

	// initialize MongoDB connection
	client, _ := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	router := gin.Default()

	// Set up CORS
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router.Use(corsConfig)

	// Put client connection in context and pass it to services
	router.Use(func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	})

	// Create a router group for the API
	api := router.Group("/api")

	// Initialize services
	_ = service.NewService(api, client)

	router.Run() // listen and serve on 0.0.0.0:8080
}
