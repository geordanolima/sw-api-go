package controller

import (
	"net/http"
	"sw-api-go/database"
	"sw-api-go/model"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	name := context.Params.ByName("hello")
	context.JSON(http.StatusOK, gin.H{
		"Hello-There": name + "!",
	})
}

func GetCharacter(context *gin.Context) {
	var characters []model.Character
	database.DB.Find(&characters)
	context.JSON(http.StatusOK, characters)
}

func GetCharacterId(context *gin.Context) {
	var character model.Character
	id := context.Params.ByName("id")
	database.DB.Find(&character, id)
	if character.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"fail": "Character not found",
		})
		return
	}
	context.JSON(http.StatusOK, character)
}

func SearchCharacter(context *gin.Context) {
	var characters []model.Character
	search := context.Param("search")
	database.DB.Where("LOWER(Name) LIKE LOWER(?)", "%%"+search+"%%").Find(&characters)
	context.JSON(http.StatusOK, characters)
}

func CreateCharacter(context *gin.Context) {
	var character model.Character
	if err := context.ShouldBindJSON(&character); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&character)
	context.JSON(http.StatusOK, character)
}

func UpdateCharacter(context *gin.Context) {
	var character model.Character
	id := context.Params.ByName("id")
	database.DB.Find(&character, id)
	if character.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"fail": "Character not found",
		})
		return
	}
	if err := context.ShouldBindJSON(&character); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&character).UpdateColumns(character)
	context.JSON(http.StatusOK, character)
}

func DeleteCharacter(context *gin.Context) {
	var character model.Character
	id := context.Params.ByName("id")
	database.DB.Delete(&character, id)
	if character.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"fail": "Character not found",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"removed": "character removed",
	})

}
