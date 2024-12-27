package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/handlers"
	"github.com/iacopoghilardi/mydget-backend/internals/routes/middlewares"
)

func SetupProfileRoutes(r *gin.RouterGroup, handler *handlers.ProfileHandler) {
	r.Use(middlewares.AuthMiddleware())

	r.GET("/:id", handler.GetProfile)
	r.POST("/", handler.CreateProfile)
	r.PUT("/:id", handler.UpdateProfile)
}
