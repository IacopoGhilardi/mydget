package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/bootstrap"
)

func SetupRoutes(r *gin.Engine, handlers *bootstrap.Handlers) {
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})

	SetupUserRoutes(r, handlers.UserHandler)
	RegisterAuthRoutes(r.Group("/auth"), handlers.AuthHandler)
}
