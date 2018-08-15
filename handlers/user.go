package handlers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang-rest/models"
	"golang-rest/utils"
	"net/http"
)

var err error

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"welcome": "to the api"})
}

func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id := c.Params.ByName("id")
	var user models.User

	if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user []models.User

	if err := db.Find(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func InsertUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	c.BindJSON(&user)

	// check validations. validate settings added to struct body.
	_, validErrs := govalidator.ValidateStruct(user)
	if validErrs != nil {
		err := map[string]interface{}{"errors": govalidator.ErrorsByField(validErrs)}
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// hash user password after validate it.
	user.Password = utils.Hash(user.Password)

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var user models.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.BindJSON(&user)

		// check validations. validate settings added to struct body.
		_, validErrs := govalidator.ValidateStruct(user)
		if validErrs != nil {
			err := map[string]interface{}{"errors": govalidator.ErrorsByField(validErrs)}
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		// hash user password after validate it.
		user.Password = utils.Hash(user.Password)

		db.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	} else {
		db.Delete(&user)
	}
}
