package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/bootstrap"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

func SetupRoutes(r *gin.Engine, handlers *bootstrap.Handlers) {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.BuildSuccessResponse("OK"))
	})

	v1 := r.Group("/api/v1")

	SetupUserRoutes(v1, handlers.UserHandler)
	RegisterAuthRoutes(v1.Group("/auth"), handlers.AuthHandler)
}
