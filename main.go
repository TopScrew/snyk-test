package main

import (
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/justinas/nosurf"
)

func main() {
	fmt.Println("a12")

	fmt.Println(nosurf.CookieName)
	response := backend.DataResponse{
		Error: fmt.Errorf("example error"),
	}
	fmt.Println("Example response:", response)
	fmt.Println("fffffff")

}
