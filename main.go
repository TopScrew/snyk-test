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

	// 使用 grafana-plugin-sdk-go 库示例
	response := backend.DataResponse{
		Error: fmt.Errorf("example error"),
	}
	fmt.Println(response)
	fmt.Println(nosurf.CookieName)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
