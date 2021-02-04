package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	driver *mongo.Client
}

func newMongoClient() *mongo.Client {

	opts := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	c, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = c.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return c
}

func (c *mongoClient) TweetsPerID(userID string) ([]itemTweet, error) {
	collection := c.driver.Database("zwitscher_store").Collection("tweets")

	var results []itemTweet
	cur, err := collection.Find(context.TODO(), bson.D{primitive.E{Key: "user_id", Value: userID}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var tweet itemTweet
		err := cur.Decode(&tweet)
		if err != nil {
			log.Printf("Query monog: %v\n", err)
			continue
		}
		results = append(results, tweet)

	}

	return results, nil
}
