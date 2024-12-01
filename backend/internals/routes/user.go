package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/handlers"
)

func SetupUserRoutes(r *gin.Engine, handler *handlers.UserHandler) {
	r.GET("/users", handler.GetAll)
	r.POST("/users", handler.Create)
	r.GET("/users/:id", handler.GetById)
	r.PUT("/users/:id", handler.Update)
	r.DELETE("/users/:id", handler.Delete)
}
