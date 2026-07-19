package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AttendanceRepository struct {
	col *mongo.Collection
}

func NewAttendanceRepository(db *mongo.Database) *AttendanceRepository {
	return &AttendanceRepository{col: db.Collection("attendance")}
}

func (r *AttendanceRepository) RecordScan(ctx context.Context, username string, ts time.Time) error {
	date := ts.Format("2006-01-02")

	filter := bson.M{"username": username, "date": date}
	update := bson.M{
		"$min": bson.M{"first_seen": ts},
		"$max": bson.M{"last_seen": ts},
	}

	opts := options.UpdateOne().SetUpsert(true)

	_, err := r.col.UpdateOne(ctx, filter, update, opts)
	return err
}
