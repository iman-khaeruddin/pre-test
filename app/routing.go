package app

import (
	"github.com/iman-khaeruddin/test-warpin/controller"
)

// mapUrl listed all mapping URLs.
func mapUrl() {

	router.GET("/all-message", controller.GetMessageWS)
	router.POST("/post/:message", controller.PostMassage)
	router.GET("/get/all-message", controller.GetMessage)

}
