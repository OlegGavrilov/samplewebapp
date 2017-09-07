package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	go handleMessages()

	router := gin.Default()

	router.LoadHTMLGlob("assets/html/*")

	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	router.Run(":8080")
}
