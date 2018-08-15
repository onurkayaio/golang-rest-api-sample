package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"golang-rest/common"
	"golang-rest/middlewares"
	"golang-rest/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter() *gin.Engine {
	db, _ := common.Initialize()

	router := gin.Default()
	router.Use(common.Inject(db))
	router.Use(gin.Recovery())

	auth := router.Group("/api/v1/")
	auth.Use(middlewares.CORSMiddlewareHandler())

	// public endpoints.
	auth.POST("/access-token", GetAccessToken)

	return router
}

func TestGetAccessToken(t *testing.T) {
	testRouter := SetupRouter()

	login := models.Login{}
	login.Username = "onurkaya"
	login.Password = "onurkaya"

	body, _ := json.Marshal(login)

	buf := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/api/v1/access-token", buf)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 201 {
		t.Fatal("Invalid Request")
	}
}
