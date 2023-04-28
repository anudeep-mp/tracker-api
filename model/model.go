package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"userId,omitempty" bson:"userId"`
	Sessions  []SessionStamp     `json:"sessions"`
	CreatedAt string             `json:"createdAt"`
}

type APIResponse struct {
	Success bool        `json:"isSuccess"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type ResponseWatchStamp struct {
	UserCount int                 `json:"userCount"`
	Users     []ResponseUserStamp `json:"users"`
}

type ResponseUserStamp struct {
	ID             primitive.ObjectID     `json:"id" bson:"_id,omitempty"`
	UserID         string                 `json:"userId" bson:"userId"`
	CreatedAt      string                 `json:"createdAt"`
	LastSeenAt     string                 `json:"lastSeenAt"`
	TotalTimeSpent int                    `json:"totalTimeSpent"`
	SessionCount   int                    `json:"sessionCount"`
	Sessions       []ResponseSessionStamp `json:"sessions"`
}

type ResponseSessionStamp struct {
	UserID           string   `json:"userId" bson:"userId"`
	SessionID        string   `json:"sessionId" bson:"sessionId"`
	Timestamps       []string `json:"timeStamps"`
	TimestampsCount  int      `json:"timeStampsCount"`
	SessionStartedAt string   `json:"sessionStartedAt"`
	SessionEndedAt   string   `json:"sessionEndedAt"`
	SessionDuration  int      `json:"sessionDuration"`
}
