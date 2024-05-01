package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sw-api-go/controller"
	"sw-api-go/database"
	"sw-api-go/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var Id int

func setupTestsRoutes() *gin.Engine {
	gin.SetMode((gin.ReleaseMode))
	routes := gin.Default()
	return routes
}

func CreateMockCharacter() {
	character := model.Character{Name: "Yoda", Description: "Jedy Mester"}
	database.DB.Create(&character)
	Id = int(character.ID)
}

func DeleteMockCharacter() {
	var character model.Character
	database.DB.Delete(&character, Id)
}

func TestValidateStatusCodeHelloThere(test *testing.T) {
	request := setupTestsRoutes()
	request.GET("/:hello", controller.Hello)
	req, _ := http.NewRequest("GET", "/test", nil)
	result := httptest.NewRecorder()
	request.ServeHTTP(result, req)
	assert.Equal(test, http.StatusOK, result.Code)
	mockResponse := `{"Hello-There":"test!"}`
	response, _ := ioutil.ReadAll(result.Body)
	assert.Equal(test, mockResponse, string(response))
}

func TestGetAllCharacters(test *testing.T) {
	database.ConectDB()
	CreateMockCharacter()
	defer DeleteMockCharacter()
	request := setupTestsRoutes()
	request.GET("/character", controller.GetCharacter)
	req, _ := http.NewRequest("GET", "/character", nil)
	result := httptest.NewRecorder()
	request.ServeHTTP(result, req)
	assert.Equal(test, http.StatusOK, result.Code)
}

func TestSearchCharacters(test *testing.T) {
	database.ConectDB()
	CreateMockCharacter()
	defer DeleteMockCharacter()
	request := setupTestsRoutes()
	request.GET("/character/search/:search", controller.GetCharacter)
	req, _ := http.NewRequest("GET", "/character/search/yod", nil)
	result := httptest.NewRecorder()
	request.ServeHTTP(result, req)
	assert.Equal(test, http.StatusOK, result.Code)
}

func TestGetCharactersById(test *testing.T) {
	database.ConectDB()
	CreateMockCharacter()
	defer DeleteMockCharacter()
	request := setupTestsRoutes()
	request.GET("/character/:id", controller.GetCharacterId)
	req, _ := http.NewRequest("GET", "/character/"+strconv.Itoa(Id), nil)
	result := httptest.NewRecorder()
	request.ServeHTTP(result, req)
	assert.Equal(test, http.StatusOK, result.Code)
	var characterMock model.Character
	json.Unmarshal(result.Body.Bytes(), &characterMock)
	assert.Equal(test, "Yoda", characterMock.Name)
}

func TestDeleteCharacter(test *testing.T) {
	database.ConectDB()
	CreateMockCharacter()
	request := setupTestsRoutes()
	request.DELETE("/character/:id", controller.DeleteCharacter)
	req, _ := http.NewRequest("DELETE", "/character/"+strconv.Itoa(Id), nil)
	result := httptest.NewRecorder()
	request.ServeHTTP(result, req)
	assert.Equal(test, http.StatusOK, result.Code)
}
