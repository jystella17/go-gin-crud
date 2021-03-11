package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jystella17/go-gin-crud/models"
	//"github.com/jystella17/go-gin-crud/database"
	"log"
	"net/http"
	//"github.com/go-sql-driver/mysql"
	//"time"
)

type DB struct {
	Database *sql.DB
	//Status sql.DBStats
}

func CreatePost(c *gin.Context){
}

func GetAllPosts(c *gin.Context, db *DB) {
	var (
		post models.Board
		posts []models.Board
	)

	rows,err := db.Database.Query("SELECT auto_id, title, content, author, author_num, date, page_view FROM board")
	if err != nil {
		log.Fatalf("SQL error : %s", err)
		panic(err)
	}

	for rows.Next(){
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.AuthorNum, &post.Date, &post.PageView)
		posts = append(posts,post)
		if err != nil{
			log.Fatalf("SQL Scan Error : %s",err)
			panic(err)
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK,posts)
}

func GetPost(c *gin.Context)  { // Read Post by Userid
	c.JSON(http.StatusOK, gin.H{"message" : "This is Post View Page"})
}

func UpdatePost(c *gin.Context){

}

func DeletePost(C *gin.Context){

}
