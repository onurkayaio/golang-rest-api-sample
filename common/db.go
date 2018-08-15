package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang-rest/models"
)

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:12541254@tcp(127.0.0.1:3306)/installer?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	db.AutoMigrate(&models.User{})

	return db, err
}

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
