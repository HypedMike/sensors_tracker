package service

import (
	"move/internal/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StatsService struct {
	router                *gin.RouterGroup
	path                  string
	mongoClient           *mongo.Client
	sensorRepository      *repository.SensorRepository
	measurementRepository *repository.MeasurementRepository
	service               *Service
}

func NewStatsService(group *gin.RouterGroup, mongoClient *mongo.Client, service *Service) *StatsService {
	s := StatsService{
		router:                group,
		mongoClient:           mongoClient,
		path:                  "",
		sensorRepository:      repository.NewSensorRepository(mongoClient),
		measurementRepository: repository.NewMeasurementRepository(mongoClient),
		service:               service,
	}
	// Initialize routes or other service-specific setup here
	s.start()

	return &s
}

func (s *StatsService) start() {
	group := s.router.Group(s.path)

	// auth middleware
	group.Use(s.service.AuthMiddleware())

	// Get sensor stats
	group.GET("/stats", s.getSensorStats)
}

func (s *StatsService) getSensorStats(c *gin.Context) {
	// get sensor id from query parameters
	sensorID := c.Query("deviceId")
	if sensorID == "" {
		c.JSON(400, gin.H{"error": "deviceId is required"})
		return
	}

	metric := c.Query("metric")
	if metric == "" {
		metric = "dispmm"
	}

	// sensor id to bson.ObjectId
	_, err := bson.ObjectIDFromHex(sensorID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid deviceId format"})
		return
	}

	// get stats for the sensor
	stats, err := s.measurementRepository.StatsBySensorID(sensorID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get stats"})
		return
	}

	c.JSON(200, stats)
}
