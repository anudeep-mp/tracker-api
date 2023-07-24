package database

import (
	"context"
	"fmt"
	"os"

	"github.com/anudeep-mp/tracker/constants"
	"github.com/anudeep-mp/tracker/utilities"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TrackingCollection *mongo.Collection
var clientPointer *mongo.Client

func init() {
	//load env variables
	err := godotenv.Load()
	utilities.CheckError(err)

	//mongo db connection
	clientOption := options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECTION_STRING"))
	client, err := mongo.Connect(context.TODO(), clientOption)
	clientPointer = client

	utilities.CheckError(err)

	fmt.Println("MongoDB Connection Succesful")

	TrackingCollection = (*mongo.Collection)(client.Database(constants.DB_NAME_PROD).Collection(constants.TRACKING_COLLECTION_NAME))
}

func UpdateCollection(environment string) {
	if environment == "uat-eol" {
		TrackingCollection = (*mongo.Collection)(clientPointer.Database(constants.DB_NAME_UAT).Collection(constants.TRACKING_COLLECTION_NAME))
	} else if environment == "uat" {
		TrackingCollection = (*mongo.Collection)(clientPointer.Database(constants.DB_NAME_PROD).Collection(constants.TRACKING_COLLECTION_NAME))
	} else if environment == "prod" {
		TrackingCollection = (*mongo.Collection)(clientPointer.Database(constants.DB_NAME_OFFICIAL).Collection(constants.TRACKING_COLLECTION_NAME))
	}
}
