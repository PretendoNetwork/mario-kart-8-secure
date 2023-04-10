package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func DeletePlayerSession(pid uint32) {
	_, err := sessionsCollection.DeleteOne(context.TODO(), bson.D{{"pid", pid}})
	if err != nil {
		panic(err)
	}
}
