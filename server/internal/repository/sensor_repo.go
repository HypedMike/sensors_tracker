package repository

import (
	"context"
	"move/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SensorRepository struct {
	mongoClient    *mongo.Client
	collectionName string
	collection     *mongo.Collection
	ctx            context.Context
}

func NewSensorRepository(mongoClient *mongo.Client) *SensorRepository {
	return &SensorRepository{
		mongoClient:    mongoClient,
		collectionName: "sensors",
		collection:     mongoClient.Database(database).Collection("sensors"),
		ctx:            context.Background(),
	}
}

func (r *SensorRepository) GetAllSensors(options struct {
	Measurements bool
	Sort         *string
}) ([]models.Sensor, error) {
	var sensors []models.Sensor

	// if measurements are requested, we need to include them
	if options.Measurements {
		cursor, err := r.collection.Aggregate(r.ctx, bson.A{
			bson.D{
				{Key: "$lookup",
					Value: bson.D{
						{Key: "from", Value: "measurements"},
						{Key: "localField", Value: "_id"},
						{Key: "foreignField", Value: "sensor_id"},
						{Key: "as", Value: "measurements"},
					},
				},
			},
			bson.D{
				{Key: "$sort", Value: bson.D{{Key: *options.Sort, Value: 1}}},
			},
		})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(r.ctx)

		if err := cursor.All(r.ctx, &sensors); err != nil {
			return nil, err
		}
	} else {
		cursor, err := r.collection.Find(r.ctx, primitive.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(r.ctx)

		if err := cursor.All(r.ctx, &sensors); err != nil {
			return nil, err
		}
	}

	// reorder measurements by timestamp inside each sensor
	for i := range sensors {
		if options.Measurements {
			sensors[i].Measurements = sortMeasurementsByTimestamp(sensors[i].Measurements)
		}
	}

	return sensors, nil
}

func sortMeasurementsByTimestamp(measurements []models.Measurement) []models.Measurement {
	for i := 0; i < len(measurements)-1; i++ {
		for j := 0; j < len(measurements)-i-1; j++ {
			if measurements[j].Timestamp.After(measurements[j+1].Timestamp) {
				measurements[j], measurements[j+1] = measurements[j+1], measurements[j]
			}
		}
	}
	return measurements
}

func (r *SensorRepository) CreateSensor(sensor models.Sensor) (*models.Sensor, error) {
	sensor.CreatedAt = time.Now()
	sensor.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(r.ctx, sensor)
	if err != nil {
		return nil, err
	}

	sensor.ID = result.InsertedID.(bson.ObjectID)
	return &sensor, nil
}

func (r *SensorRepository) UpdateSensor(sensor models.Sensor) error {
	sensor.UpdatedAt = time.Now()

	filter := bson.M{"_id": sensor.ID}
	update := bson.M{
		"$set": bson.M{
			"name":       sensor.Name,
			"location":   sensor.Location,
			"threshold":  sensor.Threshold,
			"updated_at": sensor.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(r.ctx, filter, update)
	return err
}

func (r *SensorRepository) DeleteSensor(id bson.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.collection.DeleteOne(r.ctx, filter)
	return err
}

func (r *SensorRepository) GetSensorByID(id bson.ObjectID, options struct {
	Measurements bool
}) (*models.Sensor, error) {
	if options.Measurements {
		pipeline := mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}},
			bson.D{{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "measurements"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "sensor_id"},
				{Key: "as", Value: "measurements"},
			}}},
		}
		cursor, err := r.collection.Aggregate(r.ctx, pipeline)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(r.ctx)

		var sensors []models.Sensor
		if err := cursor.All(r.ctx, &sensors); err != nil {
			return nil, err
		}
		if len(sensors) == 0 {
			return nil, mongo.ErrNoDocuments
		}

		// reorder measurements by timestamp
		sensors[0].Measurements = sortMeasurementsByTimestamp(sensors[0].Measurements)

		// return the first sensor (there should only be one)
		return &sensors[0], nil
	} else {
		var sensor models.Sensor
		filter := bson.M{"_id": id}
		err := r.collection.FindOne(r.ctx, filter).Decode(&sensor)
		if err != nil {
			return nil, err
		}
		return &sensor, nil
	}
}
