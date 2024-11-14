package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
	"net/http"
)

func main() {
	fmt.Println(nosurf.CookieName)
	fmt.Println("aaaaaa1")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
