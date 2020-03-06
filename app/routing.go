package app

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/test-warpin/controller"
)

func mapUrl() {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})*/

	router.POST("/post/:message", controller.PostMassage)

}
