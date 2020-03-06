package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostMassage(c *gin.Context) {
	message := c.Param("message")

	c.String(http.StatusOK, message)

	return
}
