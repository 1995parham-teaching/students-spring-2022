package student

import (
	"context"
	"fmt"

	"githuh.com/1995parham-teaching/students/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoDBCollection = "students"

type Mongodb struct {
	database *mongo.Database
}

func NewMongodb(database *mongo.Database) *Mongodb {
	return &Mongodb{
		database: database,
	}
}

func (m *Mongodb) GetAll() ([]model.Student, error) {
	ctx := context.Background()
	result := make([]model.Student, 0)

	collection := m.database.Collection(MongoDBCollection)

	it, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("mongodb find query failed %w", err)
	}

	for it.Next(ctx) {
		var std model.Student

		if err := it.Decode(&std); err != nil {
			return nil, fmt.Errorf("mongodb student decode failed %w", err)
		}

		result = append(result, std)
	}

	return result, nil
}

func (m *Mongodb) Get(id int64) (model.Student, error) {
	var std model.Student

	ctx := context.Background()

	collection := m.database.Collection(MongoDBCollection)

	result := collection.FindOne(ctx, bson.M{"id": id})
	if result.Err() != nil {
		return std, fmt.Errorf("mongodb find query failed %w", result.Err())
	}

	if err := result.Decode(&std); err != nil {
		return std, fmt.Errorf("mongodb student decode failed %w", err)
	}

	return std, nil
}

func (m *Mongodb) Set(std model.Student) error {
	ctx := context.Background()

	collection := m.database.Collection(MongoDBCollection)

	if _, err := collection.InsertOne(ctx, std); err != nil {
		return fmt.Errorf("mongodb insertion failed %w", err)
	}

	return nil
}
