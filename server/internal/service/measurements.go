package service

import (
	"move/internal/models"
	"move/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MeasurementService struct {
	service               *Service
	router                *gin.RouterGroup
	path                  string
	mongoClient           *mongo.Client
	measurementRepository *repository.MeasurementRepository
}

func NewMeasurementService(group *gin.RouterGroup, mongoClient *mongo.Client, service *Service) *MeasurementService {
	s := MeasurementService{
		router:                group,
		mongoClient:           mongoClient,
		path:                  "/measurements",
		measurementRepository: repository.NewMeasurementRepository(mongoClient),
		service:               service,
	}
	// Initialize routes or other service-specific setup here
	s.start()

	return &s
}

func (s *MeasurementService) start() {
	g := s.router.Group(s.path)

	// add auth middleware
	g.Use(s.service.AuthMiddleware())

	// Get measurements by sensor ID
	g.GET("", s.getMeasurementsBySensorID)

	// Add measurements to a specific sensor
	g.POST("/:sensor_id", s.addMeasurements)

	// Add measurements to multiple sensors
	g.POST("", s.addMeasurementsToMultipleSensors)

	// Delete measurements by sensor ID
	g.DELETE("/:sensor_id", s.deleteMeasurementsBySensorID)

}

func (s *MeasurementService) getMeasurementsBySensorID(c *gin.Context) {
	sensorID := c.Query("deviceId")

	// get filters from query parameters
	from := c.Query("from")
	to := c.Query("to")
	var filter struct {
		From time.Time
		To   time.Time
	}
	if from != "" {
		filter.From, _ = time.Parse(time.RFC3339, from)
	}
	if to != "" {
		filter.To, _ = time.Parse(time.RFC3339, to)
	}

	// get measurements from the repository
	measurements, err := s.measurementRepository.GetMeasurementsBySensorID(sensorID, filter)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get measurements"})
		return
	}

	c.JSON(200, measurements)
}

// Add measurements to a specific sensor
func (s *MeasurementService) addMeasurements(c *gin.Context) {
	var values []float64
	if err := c.BindJSON(&values); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	sensorID := c.Param("sensor_id")
	if err := s.measurementRepository.AddValues(sensorID, values); err != nil {
		c.JSON(500, gin.H{"error": "Failed to add measurements"})
		return
	}
	c.JSON(200, gin.H{"message": "Measurements added successfully"})
}

// Add measurements to multiple sensors
func (s *MeasurementService) addMeasurementsToMultipleSensors(c *gin.Context) {
	var measurementsRaw []struct {
		SensorID string  `json:"sensor_id"`
		Value    float64 `json:"value"`
	}
	if err := c.BindJSON(&measurementsRaw); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// convert raw measurements to models.Measurement
	var measurements []models.Measurement
	for _, raw := range measurementsRaw {
		id, err := bson.ObjectIDFromHex(raw.SensorID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid sensor ID"})
			return
		}
		measurements = append(measurements, models.Measurement{
			SensorID: id,
			Value:    raw.Value,
		})
	}

	// add measurements to the repository
	if err := s.measurementRepository.AddMeasurements(measurements); err != nil {
		c.JSON(500, gin.H{"error": "Failed to add measurements"})
		return
	}
	c.JSON(200, gin.H{"message": "Measurements added successfully"})
}

// Delete measurements by sensor ID
func (s *MeasurementService) deleteMeasurementsBySensorID(c *gin.Context) {
	sensorID := c.Param("sensor_id")

	// delete measurements from the repository
	if err := s.measurementRepository.DeleteMeasurementsBySensorID(sensorID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete measurements"})
		return
	}

	c.JSON(200, gin.H{"message": "Measurements deleted successfully"})
}
