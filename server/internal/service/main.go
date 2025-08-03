package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service struct {
	router      *gin.RouterGroup
	mongoClient *mongo.Client

	sensorService      *SensorService
	seederService      *SeederService
	measurementService *MeasurementService
}

func NewService(router *gin.RouterGroup, mongoClient *mongo.Client) *Service {
	s := &Service{
		router:             router,
		mongoClient:        mongoClient,
		sensorService:      NewSensorService(router, mongoClient),
		seederService:      NewSeederService(router, mongoClient),
		measurementService: NewMeasurementService(router, mongoClient),
	}

	return s
}
