package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jystella17/go-gin-crud/models"
	"log"

	//"github.com/go-sql-driver/mysql"
	"net/http"
)

func (db *DB) GetHomePage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "Go-Gin-CRUD",
	})
}

func (db *DB) GetAllUsers (c *gin.Context){
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