package main

import (
	"log"

	"github.com/helloDevAman/movie-base/apis/routes"
)

func main() {
	routes.InitRoutes()
	log.Println("Server started on port 8080")
}
