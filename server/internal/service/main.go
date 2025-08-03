package service

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service struct {
	router      *gin.RouterGroup
	mongoClient *mongo.Client

	sensorService      *SensorService
	seederService      *SeederService
	measurementService *MeasurementService
	authService        *AuthService
}

func NewService(router *gin.RouterGroup, mongoClient *mongo.Client) *Service {
	s := &Service{
		router:      router,
		mongoClient: mongoClient,
	}

	// Initialize routes
	s.seederService = NewSeederService(router, mongoClient)
	s.sensorService = NewSensorService(router, mongoClient, s)
	s.measurementService = NewMeasurementService(router, mongoClient, s)
	s.authService = NewAuthService(router, mongoClient, s)

	return s
}

func (s *Service) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{Err: jwt.ErrSignatureInvalid}
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// get user ID from token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		userId := claims["user_id"]
		if userId == nil {
			c.JSON(401, gin.H{"error": "User ID not found in token"})
			c.Abort()
			return
		}

		// get user from database or context
		user, err := s.authService.userRepository.GetUserByID(userId.(string))
		if err != nil {
			c.JSON(401, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user)

		// continue to the next handler
		c.Next()
	}
}
