# Golang GIN Framework REST Api sample with JWT, Gorm, Testify.

## Overview [![GoDoc](https://godoc.org/github.com/onurkayaio/golang-rest-api-sample?status.svg)](https://godoc.org/github.com/onurkayaio/golang-rest-api-sample)

An example of Golang REST Api with JWT, Gorm, Testify.

## Install

```
go get github.com/onurkayaio/golang-rest-api-sample
```
Before build the project you should change database settings.

```
db, err := gorm.Open("mysql", "YOUR_USERNAME:YOUR_PASSWORD@tcp(YOUR_HOST)/DB_NAME?charset=utf8&parseTime=True&loc=Local")
```

## Used Packages

```
https://github.com/gin-gonic/gin
https://github.com/stretchr/testify
https://github.com/jinzhu/gorm
```

## Endpoints

```
  auth.POST("/access-token", handlers.GetAccessToken)
  auth.GET("/", handlers.Index)
  auth.GET("/users", handlers.GetUsers)
  auth.GET("/users/:id", handlers.GetUser)
  auth.POST("/users", handlers.InsertUser)
  auth.PUT("/users/:id", handlers.UpdateUser)
  auth.DELETE("/users/:id", handlers.DeleteUser)
```  
  
## Run Tests

For example to run auth_test.go, your current path should be:

```
/go/src/golang-rest/handlers
```
