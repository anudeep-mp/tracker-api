package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/anudeep-mp/tracker/database"
	"github.com/anudeep-mp/tracker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostWatchStamp(watchStamp model.WatchStamp) (model.UserStamp, error) {

	filter := bson.M{"userId": watchStamp.UserID}

	var user model.UserStamp

	databaseErr := database.TrackingCollection.FindOne(context.Background(), filter).Decode(&user)

	if databaseErr == mongo.ErrNoDocuments {
		//if documnet if not available, create a new one
		user = model.UserStamp{
			UserID:    watchStamp.UserID,
			Sessions:  make([]model.SessionStamp, 0),
			CreatedAt: watchStamp.TimeStamp,
		}

		user = updateSession(watchStamp, user)

		insertedRes, insertionErr := database.TrackingCollection.InsertOne(context.Background(), user)

		if insertionErr != nil {
			return user, fmt.Errorf("failed to insert new user %w", insertionErr)
		}

		user.ID = insertedRes.InsertedID.(primitive.ObjectID)

	} else {

		user = updateSession(watchStamp, user)

		update := bson.M{"$set": bson.M{"sessions": user.Sessions}}

		_, updatedErr := database.TrackingCollection.UpdateOne(context.Background(), filter, update)

		if updatedErr != nil {
			return user, fmt.Errorf("failed to update user %w", updatedErr)
		}
	}

	return user, nil
}

func updateSession(watchStamp model.WatchStamp, user model.UserStamp) model.UserStamp {

	var isSessionAvailable = false

	for i := range user.Sessions {
		tempSession := user.Sessions[i]

		if tempSession.SessionID == watchStamp.SessionID {
			tempSession.TimeStamps = append(tempSession.TimeStamps, watchStamp.TimeStamp)
			user.Sessions[i] = tempSession
			isSessionAvailable = true
			break
		}
	}

	if !isSessionAvailable {
		var newSession = model.SessionStamp{
			UserID:    watchStamp.UserID,
			SessionID: watchStamp.SessionID,
			TimeStamps: []string{
				watchStamp.TimeStamp,
			},
		}
		user.Sessions = append(user.Sessions, newSession)
	}

	return user
}

func GetWatchStamps() ([]model.ResponseUserStamp, error) {
	cursor, err := database.TrackingCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	defer cursor.Close(context.Background())

	var users = make([]model.UserStamp, 0)

	for cursor.Next(context.Background()) {
		var user model.UserStamp

		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("error decoding user details: %w", err)
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error during cursor iteration: %w", err)
	}

	var usersResponse []model.ResponseUserStamp

	for _, user := range users {
		usersResponse = append(usersResponse, updateUserResponse(user))
	}

	return usersResponse, nil
}

func updateUserResponse(user model.UserStamp) model.ResponseUserStamp {

	var sessionsResponse []model.ResponseSessionStamp
	var totalTimeSpent int = 0

	for _, session := range user.Sessions {
		var sessionResponse model.ResponseSessionStamp

		sessionResponse.UserID = session.UserID
		sessionResponse.SessionID = session.SessionID
		sessionResponse.SessionStartedAt = session.TimeStamps[0]
		sessionResponse.SessionEndedAt = session.TimeStamps[len(session.TimeStamps)-1]

		firstTime, _ := time.Parse(time.RFC3339, session.TimeStamps[0])
		lastTime, _ := time.Parse(time.RFC3339, session.TimeStamps[len(session.TimeStamps)-1])
		duration := int(lastTime.Sub(firstTime).Seconds())
		totalTimeSpent += duration
		sessionResponse.SessionDuration = duration
		sessionResponse.TimestampsCount = len(session.TimeStamps)
		sessionResponse.Timestamps = session.TimeStamps
		sessionsResponse = append(sessionsResponse, sessionResponse)
	}

	var userResponse model.ResponseUserStamp
	userResponse.ID = user.ID
	userResponse.UserID = user.UserID
	userResponse.CreatedAt = user.CreatedAt
	userResponse.Sessions = sessionsResponse
	userResponse.LastSeenAt = sessionsResponse[len(sessionsResponse)-1].SessionEndedAt
	userResponse.TotalTimeSpent = totalTimeSpent
	userResponse.SessionCount = len(user.Sessions)

	return userResponse
}

func DeleteAllWatchStamps() error {
	_, err := database.TrackingCollection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		return fmt.Errorf("failed to delete userstamps: %w", err)
	}
	return nil
}

func DeleteWatchStamp(userId string) error {
	_, err := database.TrackingCollection.DeleteOne(context.Background(), bson.M{"userId": userId})

	if err != nil {
		return fmt.Errorf("failed to delete userstamp: %w", err)
	}
	return nil
}
