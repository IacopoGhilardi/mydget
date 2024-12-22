package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/bootstrap"
	"github.com/iacopoghilardi/mydget-backend/internals/config"
	"github.com/iacopoghilardi/mydget-backend/internals/db"
	"github.com/iacopoghilardi/mydget-backend/internals/routes"
)

func main() {
	app, err := bootstrap.NewApplication()
	if err != nil {
		log.Fatal("Failed to initialize application: ", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, app.Handlers)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")

		if err := db.Close(); err != nil {
			log.Fatal("Failed to close database connection: ", err)
		}

		os.Exit(0)
	}()

	log.Println("Server is running on port: ", config.GetConfig().BackendPort)
	r.Run(":" + config.GetConfig().BackendPort)
}
