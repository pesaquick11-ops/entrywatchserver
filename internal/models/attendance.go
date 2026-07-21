package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Attendance struct {
	ID                  bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username            string        `bson:"username" json:"username"`
	Date                string        `bson:"date" json:"date"` // "2026-07-17"
	FirstSeen           time.Time     `bson:"first_seen" json:"firstSeen"`
	FirstSeenConfidence string        `bson:"first_seen_confidence" json:"firstSeenConfidence"`
	LastSeen            time.Time     `bson:"last_seen" json:"lastSeen"`
	LastSeenConfidence  string        `bson:"last_seen_confidence" json:"lastSeenConfidence"`
}
