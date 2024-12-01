package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/bootstrap"
	"github.com/iacopoghilardi/mydget-backend/internals/routes"
)

func main() {
	app, err := bootstrap.NewApplication()
	if err != nil {
		log.Fatal("Failed to initialize application: ", err)
	}

	r := gin.Default()

	routes.SetupRoutes(r, app.Handlers)
	r.Run(":8080")
}
