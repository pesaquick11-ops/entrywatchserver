package repository

import (
	"context"
	"entrywatchserver/internal/models"
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

func (r *AttendanceRepository) FindAll(ctx context.Context) ([]models.Attendance, error) {

	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var records []models.Attendance

	err = cur.All(ctx, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
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
