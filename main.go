package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/justinas/nosurf"
	"net/http"
)

func main() {
	fmt.Println("a12")

	r := gin.Default()
	response := backend.DataResponse{
		Error: fmt.Errorf("example error"),
	}
	fmt.Println("Example response:", response)
	fmt.Println(nosurf.CookieName)
	fmt.Println("fffffff")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
