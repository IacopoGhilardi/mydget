package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/bootstrap"
	v1 "github.com/iacopoghilardi/mydget-backend/internals/routes/v1"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

func SetupRoutes(r *gin.Engine, handlers *bootstrap.Handlers) {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.BuildSuccessResponse("OK"))
	})

	v1Group := r.Group("/api/v1")

	v1.SetupUserRoutes(v1Group.Group("/users"), handlers.UserHandler)
	RegisterAuthRoutes(v1Group.Group("/auth"), handlers.AuthHandler)
}
