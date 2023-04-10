package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPlayerURLs(pid uint32) []string {
	var result bson.M

	err := sessionsCollection.FindOne(context.TODO(), bson.D{{"pid", pid}}, options.FindOne()).Decode(&result)
	if err != nil {
		panic(err)
	}

	oldurlArray := result["urls"].(bson.A)
	newurlArray := make([]string, len(oldurlArray))
	for i := 0; i < len(oldurlArray); i++ {
		newurlArray[i] = oldurlArray[i].(string)
	}

	return newurlArray
}
