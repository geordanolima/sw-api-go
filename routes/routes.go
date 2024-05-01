package routes

import (
	"sw-api-go/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	request := gin.Default()
	request.GET("/character", controller.GetCharacter)
	request.GET("/character/:id", controller.GetCharacterId)
	request.GET("/:hello", controller.Hello)
	request.POST("/character", controller.CreateCharacter)
	request.DELETE("/character/:id", controller.DeleteCharacter)
	request.GET("/character/search/:search", controller.SearchCharacter)
	request.Run()
}
