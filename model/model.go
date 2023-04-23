package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type WatchStamp struct {
	UserID    string `json:"userId,omitempty" bson:"userId"`
	SessionID string `json:"sessionId,omitempty" bson:"sessionId"`
	TimeStamp string `json:"timeStamp,omitempty"`
}
type SessionStamp struct {
	UserID     string   `json:"userId,omitempty" bson:"userId"`
	SessionID  string   `json:"sessionId,omitempty" bson:"sessionId"`
	TimeStamps []string `json:"timeStamps,omitempty"`
}
type UserStamp struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID   string             `json:"userId,omitempty" bson:"userId"`
	Sessions []SessionStamp     `json:"sessions"`
}
