package service

import (
	"move/internal/repository"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthService struct {
	router         *gin.RouterGroup
	path           string
	mongoClient    *mongo.Client
	userRepository *repository.UserRepository
	service        *Service
}

func NewAuthService(router *gin.RouterGroup, mongoClient *mongo.Client, service *Service) *AuthService {
	a := &AuthService{
		router:         router,
		path:           "/auth",
		mongoClient:    mongoClient,
		userRepository: repository.NewUserRepository(mongoClient),
		service:        service,
	}
	a.Start()
	return a
}

func (a *AuthService) Start() {
	authGroup := a.router.Group(a.path)

	authGroup.POST("/login", a.login)
	authGroup.GET("/user", a.getUserByJWT)
}

func (a *AuthService) login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	user, err := a.userRepository.GetUserByUsernameAndPassword(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	// create a jwt token
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      jwt.TimeFunc().Add(24 * time.Hour).Unix(), // Token valid for 24 hours
	}).SignedString([]byte(os.Getenv("JWT_SECRET"))) // Use a secure secret in production
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful", "user": user, "token": jwtToken})
}

func (a *AuthService) getUserByJWT(c *gin.Context) {
	// get auth header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	// parse token
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
	user, err := a.userRepository.GetUserByID(userId.(string))
	if err != nil {
		c.JSON(401, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	c.JSON(200, user)
}
