package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/test-warpin/database"
	"net/http"
)

func PostMassage(c *gin.Context) {
	message := c.Param("message")

	//save message to array
	database.LocalDB = append(database.LocalDB, message)

	c.String(http.StatusOK, message+" successfully receive")

	return
}
