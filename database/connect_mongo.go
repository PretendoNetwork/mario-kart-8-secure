package database

import (
	"context"
	"time"

	"github.com/PretendoNetwork/mario-kart-8-secure/globals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var mongoContext context.Context
var accountDatabase *mongo.Database
var mk8Database *mongo.Database
var pnidCollection *mongo.Collection
var nexAccountsCollection *mongo.Collection
var regionsCollection *mongo.Collection
var usersCollection *mongo.Collection
var sessionsCollection *mongo.Collection
var roomsCollection *mongo.Collection
var tourneysCollection *mongo.Collection

func connectMongo() {
	if globals.Config.DatabaseUseAuth {
		mongoClient, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://" + globals.Config.DatabaseUsername + ":" + globals.Config.DatabasePassword + "@" + globals.Config.DatabaseIP + ":" + globals.Config.DatabasePort + "/"))
	} else {
		mongoClient, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://" + globals.Config.DatabaseIP + ":" + globals.Config.DatabasePort + "/"))
	}
	mongoContext, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_ = mongoClient.Connect(mongoContext)

	accountDatabase = mongoClient.Database(globals.Config.AccountDatabase)
	pnidCollection = accountDatabase.Collection(globals.Config.PNIDCollection)
	nexAccountsCollection = accountDatabase.Collection(globals.Config.NexAccountsCollection)

	mk8Database = mongoClient.Database(globals.Config.MK8Database)
	regionsCollection = mk8Database.Collection(globals.Config.RegionsCollection)
	usersCollection = mk8Database.Collection(globals.Config.UsersCollection)
	sessionsCollection = mk8Database.Collection(globals.Config.SessionsCollection)
	roomsCollection = mk8Database.Collection(globals.Config.RoomsCollection)
	tourneysCollection = mk8Database.Collection(globals.Config.TournamentsCollection)

	sessionsCollection.DeleteMany(context.TODO(), bson.D{})
	roomsCollection.DeleteMany(context.TODO(), bson.D{})
}
