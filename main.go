package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router.GET("/home", service.GetHomePage)
	router.GET("/user/:user_id", service.GetUserById)
	router.GET("/board/list", service.GetAllPosts)
	router.GET("/board/list/:user_id", service.GetPost)
	router.GET("/board/create-post", service.CreatePost)
	router.POST("/login", service.UserLogin)
	router.POST("/register", service.UserRegister)
	router.PUT("/board/update-post", service.UpdatePost)
	router.DELETE("/board/delete-post", service.DeletePost)

	return router

	return router
}