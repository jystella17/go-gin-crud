package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jystella17/go-gin-crud/database"
	"github.com/jystella17/go-gin-crud/service"
	"log"
	"time"
)

func main(){
	r := setupRouter()

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start Gin Engine")
	}

	r.Run()
}

func setupRouter() *gin.Engine{
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET","POST","PUT","DELETE","PATCH"},
		AllowHeaders: []string{"Authorization","Access-Control-Allow-Origin","Origin","Content-Length","Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 1 * time.Hour,
	})) // CORS setting

	db,err := database.Connection()
	if err != nil{
		log.Fatalf("Failed to connect DB")
		panic(err)
	}

	DBCtx := service.DB{
		Database: db,
	}

	router.GET("/home")
	router.GET("/user/:user_id", service.GetUserById)
	router.GET("/board/list")
	router.GET("/board/list/:_id", service.GetPost)
	router.GET("/user/:user_id/posts", service.GetPostById)
	router.GET("/admin/user-info")
	router.POST("/board/create-post", service.CreatePost)
	router.POST("/login", service.UserLogin)
	router.POST("/register", service.UserRegister)
	router.PUT("/board/update-post", service.UpdatePost)
	router.DELETE("/board/delete-post", service.DeletePost)

	return router
}