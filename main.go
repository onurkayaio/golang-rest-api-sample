package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang-rest/common"
	"golang-rest/handlers"
	"golang-rest/middlewares"
)

var db *gorm.DB
var err error

func main() {
	db, _ := common.Initialize()

	router := gin.Default()
	router.Use(common.Inject(db))
	router.Use(gin.Recovery())

	auth := router.Group("/api/v1/")
	auth.Use(middlewares.CORSMiddlewareHandler())

	// public endpoints.
	auth.POST("/access-token", handlers.GetAccessToken)

	auth.Use(middlewares.Authorization())
	// restricted endpoints.
	{
		auth.GET("/", handlers.Index)
		auth.GET("/users", handlers.GetUsers)
		auth.GET("/users/:id", handlers.GetUser)
		auth.POST("/users", handlers.InsertUser)
		auth.PUT("/users/:id", handlers.UpdateUser)
		auth.DELETE("/users/:id", handlers.DeleteUser)
	}

	router.Run(":8081")
}
