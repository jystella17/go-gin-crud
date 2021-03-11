package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jystella17/go-gin-crud/models"
	"log"

	//"github.com/go-sql-driver/mysql"
	"net/http"
)

func GetHomePage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Go-Gin-CRUD",
	})
}

func GetAllUsers (c *gin.Context, db *DB){
	var (
		user models.User
		users []models.User
	)

	rows,err := db.Database.Query("SELECT auto_id, userid, password, name, student_num FROM user")
	if err != nil {
		log.Fatalf("SQL error : %s", err)
		panic(err)
	}

	for rows.Next(){
		err = rows.Scan(&user.Id, &user.UserId, &user.Password, &user.Name, &user.StudentNum)
		users = append(users,user)
		if err != nil{
			log.Fatalf("SQL Scan Error : %s",err)
			panic(err)
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK,users)
}

func UserLogin(c *gin.Context){ // Login with POST Method
	//var db *DB
}

func UserRegister(c *gin.Context){ // Register with POST Method
	//var db *DB
}

func GetUserById(c *gin.Context){ // Load User information with GET Method
	//var db *DB
	c.JSON(http.StatusOK, gin.H{"message" : "This is User Info Page"})
}

func GetPostById(c *gin.Context){
	//var db *DB
}
