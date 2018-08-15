package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-rest/models"
	"golang-rest/utils"
	"net/http"
)

func GetAccessToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// create login user struct to restrict data.
	var login models.Login
	err := c.BindJSON(&login)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	// create user struct and get user by username.
	var user models.User
	if err := db.Where("username = ?", login.Username).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	// check user password is equal with login user password, if equal generate token and return it.
	if utils.CheckPassword(user, login.Password) {
		token, err := utils.GenerateJWT(user)
		if err != nil {
			c.AbortWithError(http.StatusUnprocessableEntity, err)
			c.Abort()
		}

		c.JSON(http.StatusCreated, gin.H{"token": token})
	} else {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
}
