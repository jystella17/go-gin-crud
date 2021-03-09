package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client *mongo.Client
	db *mongo.Database
)

func Connection()(*mongo.Client,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client,err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalf("Connection error : %s",err)
		panic(err)
	}

	if err := client.Ping(ctx,nil); err != nil{
		log.Fatalf("Ping error : %s", err)
		panic(err)
	}

	return client,nil
}