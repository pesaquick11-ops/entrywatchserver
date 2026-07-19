package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Attendance struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string        `bson:"username" json:"username"`
	Date      string        `bson:"date" json:"date"` // "2026-07-17"
	FirstSeen time.Time     `bson:"first_seen" json:"firstSeen"`
	LastSeen  time.Time     `bson:"last_seen" json:"lastSeen"`
}
