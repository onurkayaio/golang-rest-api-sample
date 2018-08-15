package fixtures

import (
	"github.com/jinzhu/gorm"
	"golang-rest/models"
	"golang-rest/utils"
	"strconv"
	"testing"
)

func TestInsertUser(t *testing.T) {
	db, err := gorm.Open("mysql", "root:12541254@tcp(127.0.0.1:3306)/installer?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		t.Fail()
	}

	for i := 1; i < 5; i++ {
		user := models.User{
			Username: "testuser" + strconv.Itoa(i),
			Password: utils.Hash("testuser" + strconv.Itoa(i)),
			Email:    "test" + strconv.Itoa(i) + "@test.com",
			Role:     "admin",
			Status:   "active",
		}

		if err := db.Create(&user).Error; err != nil {
			t.Fail()
		}
	}
}
