package repository

import (
	"context"
	"fmt"
	"move/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	mongoClient    *mongo.Client
	collectionName string
	collection     *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(mongoClient *mongo.Client) *UserRepository {
	collection := mongoClient.Database("move").Collection("users")
	return &UserRepository{
		mongoClient:    mongoClient,
		collectionName: "users",
		collection:     collection,
		ctx:            context.Background(),
	}
}

func (r *UserRepository) GetUserByUsernameAndPassword(username, password string) (models.User, error) {
	var user models.User

	filter := bson.D{
		{Key: "username", Value: username},
	}

	err := r.collection.FindOne(r.ctx, filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	if !r.comparePassword(user.Password, password) {
		return models.User{}, fmt.Errorf("invalid username or password")
	}

	return user, nil
}

func (r *UserRepository) hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (r *UserRepository) comparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (r *UserRepository) CreateUser(user models.User) error {
	hashedPassword, err := r.hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user.Password = hashedPassword

	_, err = r.collection.InsertOne(r.ctx, user)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetUserByID(userID string) (models.User, error) {
	var user models.User

	id, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid user ID: %w", err)
	}

	filter := bson.D{{Key: "_id", Value: id}}

	err = r.collection.FindOne(r.ctx, filter).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("error finding user: %w", err)
	}

	return user, nil
}
