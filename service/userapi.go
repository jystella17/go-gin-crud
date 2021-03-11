package service

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-sql-driver/mysql"
	"net/http"
)

func (db *DB) GetHomePage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Go-Gin-CRUD",
	})
}

func (db *DB) UserLogin(c *gin.Context){ // Login with POST Method
}

func (db *DB) UserRegister(c *gin.Context){ // Register with POST Method
	return
}

func (db *DB) GetUserById(c *gin.Context){ // Load User information with GET Method
	c.JSON(http.StatusOK, gin.H{"message" : "This is User Info Page"})
}

func (db *DB) GetPostById(c *gin.Context){
}