package service

import (
	//"github.com/jystella17/SM-Literature-Map/model"
	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func GetHomePage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "SM-Literature-Map",
	})
}

func UserLogin(c *gin.Context){ // Login with POST Method
}

func UserRegister(c *gin.Context){ // Register with POST Method
	return
}

func GetUserById(c *gin.Context){ // Load User information with GET Method
	c.JSON(http.StatusOK, gin.H{"message" : "This is User Info Page"})
}