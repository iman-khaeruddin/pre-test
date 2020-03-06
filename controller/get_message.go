package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/test-warpin/database"
	"net/http"
)

func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, database.LocalDB)
	return
}
