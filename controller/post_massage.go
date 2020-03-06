package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/test-warpin/database"
	"net/http"
)

// PostMassage used for save massage into array.
func PostMassage(c *gin.Context) {
	message := c.Param("message")

	database.LocalDB = append(database.LocalDB, message)

	c.String(http.StatusOK, message+" successfully receive")

	return
}
