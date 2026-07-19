package repository

import (
	"context"
	"entrywatchserver/internal/models"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{col: db.Collection("users")}
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (bson.M, error) {
	var result bson.M
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	return result, err
}

func (r *UserRepository) FindAll(ctx context.Context) ([]bson.M, error) {
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []bson.M
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *UserRepository) AddUser(ctx context.Context, user models.User) (bson.ObjectID, error) {
	result, err := r.col.InsertOne(ctx, user)
	if err != nil {
		return bson.ObjectID{}, err
	}
	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return bson.ObjectID{}, fmt.Errorf("Unexpected isert Id type")
	}
	return id, nil
}
