package main

import (
	"github.com/jinzhu/gorm"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	_, err := gorm.Open("mysql", "root:12541254@tcp(127.0.0.1:3306)/installer?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		t.Fail()
	}
}
