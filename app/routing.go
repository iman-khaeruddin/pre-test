package app

import (
	"github.com/iman-khaeruddin/test-warpin/controller"
)

func mapUrl() {

	router.GET("/ws", controller.GetMessageWS)

	router.POST("/post/:message", controller.PostMassage)
	router.GET("/get/all-message", controller.GetMessage)

}
