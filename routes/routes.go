package routes

import (
	"sw-api-go/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	request := gin.Default()
	request.LoadHTMLGlob("templates/*")
	request.Static("/assets", "./assets")
	request.GET("/:hello", controller.Hello)
	request.GET("/character", controller.GetCharacter)
	request.GET("/character/:id", controller.GetCharacterId)
	request.POST("/character", controller.CreateCharacter)
	request.DELETE("/character/:id", controller.DeleteCharacter)
	request.GET("/character/search/:search", controller.SearchCharacter)
	request.PATCH("/character/:id", controller.UpdateCharacter)
	request.GET("/home", controller.ViewIndexPage)
	request.NoRoute(controller.NotFoundPage)
	request.Run()
}
