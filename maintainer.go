package maintainer

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/go-redis/redis"
	"log"
)

type Maintainer struct {
	mongodb *mongo.Client
	redis *redis.Client
}

func New(mongoAddress string, redisAddress string) *Maintainer {

	maint := &Maintainer{}

	// Connect to the mongodb database
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoAddress)
	mdb, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil
	}

	// Check the connection
	err = mdb.Ping(context.TODO(), nil)
	if err != nil {
		return nil
	}

	maint.mongodb = mdb

	// Connect to Redis!
	rclient := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	pong, err := rclient.Ping().Result()
	if err != nil || pong != "PONG" {
		log.Fatalf("Failed to connect to redis: %s - %s\n", pong, err)
	}
	maint.redis = rclient

	return maint
}