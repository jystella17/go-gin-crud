package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context){
}

func GetAllPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Connection error : %s", err)
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Ping error : %s", err)
		panic(err)
	}
}

func GetPost(c *gin.Context)  { // Read Post by Userid
	c.JSON(http.StatusOK, gin.H{"message" : "This is Post View Page"})
}

func UpdatePost(c *gin.Context){

}

func DeletePost(C *gin.Context){

}
