package service

import (
	"fmt"
	"move/internal/models"
	"move/internal/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SensorService struct {
	router                *gin.RouterGroup
	path                  string
	mongoClient           *mongo.Client
	sensorRepository      *repository.SensorRepository
	measurementRepository *repository.MeasurementRepository
	service               *Service
}

func NewSensorService(group *gin.RouterGroup, mongoClient *mongo.Client, service *Service) *SensorService {
	s := SensorService{
		router:                group,
		mongoClient:           mongoClient,
		path:                  "/devices",
		sensorRepository:      repository.NewSensorRepository(mongoClient),
		measurementRepository: repository.NewMeasurementRepository(mongoClient),
		service:               service,
	}
	// Initialize routes or other service-specific setup here
	s.start()

	return &s
}

func (s *SensorService) start() {
	group := s.router.Group(s.path)

	// auth middleware
	group.Use(s.service.AuthMiddleware())

	// Get all sensors
	group.GET("", s.getAllSensors)

	// Get sensor by ID
	group.GET("/:sensor_id", s.getSensorByID)

	// Create a new sensor
	group.POST("", s.createSensor)

	// Delete a sensor by ID
	group.DELETE("/:sensor_id", s.deleteSensor)
}

func (s *SensorService) getAllSensors(c *gin.Context) {
	// check if there is a query parameter for measurements
	measurements := c.Query("m") == "true"
	sort := c.Query("sort")
	if sort == "" {
		sort = "created_at" // default sort field
	}

	// get all sensors from the repository
	sensor, err := s.sensorRepository.GetAllSensors(struct {
		Measurements bool
		Sort         *string
	}{
		Measurements: measurements,
		Sort:         &sort,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get sensors"})
		return
	}

	// if measurements are requested, calculate the stats
	if measurements {
		for i := range sensor {
			stats, err := s.measurementRepository.StatsBySensorID(sensor[i].ID.Hex())
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to get sensor stats"})
				return
			}
			sensor[i].Average = stats.Average
			sensor[i].Min = stats.Min
			sensor[i].Max = stats.Max
		}
	}

	// return the sensors
	c.JSON(200, sensor)
}

func (s *SensorService) getSensorByID(c *gin.Context) {
	sensorID := c.Param("sensor_id")

	id, err := bson.ObjectIDFromHex(sensorID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid sensor ID"})
		return
	}

	// get sensor from the repository
	sensor, err := s.sensorRepository.GetSensorByID(id, struct {
		Measurements bool
	}{
		Measurements: true, // assuming we want measurements by default
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get sensor"})
		return
	}

	// return the sensor
	c.JSON(200, sensor)
}

func (s *SensorService) createSensor(c *gin.Context) {
	// create a new sensor
	sensor := models.Sensor{}
	// print the request body for debugging
	fmt.Printf("Request body: %v\n", c.Request.Body)
	if err := c.BindJSON(&sensor); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}
	result, err := s.sensorRepository.CreateSensor(sensor)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create sensor"})
		return
	}
	c.JSON(201, gin.H{"id": result.ID, "message": "Sensor created successfully"})
}

func (s *SensorService) deleteSensor(c *gin.Context) {
	sensorID := c.Param("sensor_id")

	// validate sensor ID
	if sensorID == "" {
		c.JSON(400, gin.H{"error": "Sensor ID is required"})
		return
	}

	// delete measurements associated with the sensor
	if err := s.measurementRepository.DeleteMeasurementsBySensorID(sensorID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete measurements"})
		return
	}

	// delete sensor from the repository
	if err := s.sensorRepository.DeleteSensorByID(sensorID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete sensor"})
		return
	}

	c.JSON(200, gin.H{"message": "Sensor deleted successfully"})
}
