package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/test-warpin/database"
	"golang.org/x/net/websocket"
	"io"
)

func GetMessageWS(c *gin.Context) {
	handler := websocket.Handler(func(conn *websocket.Conn) {

		for _, row := range database.LocalDB {
			conn.Write([]byte(row))
		}

		io.Copy(conn, conn)

	})

	handler.ServeHTTP(c.Writer, c.Request)
}