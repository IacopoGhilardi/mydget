package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/mydget-backend/internals/db"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Unauthorized", "Token is required"))
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.VerifyJWT(token)
		if err != nil {
			log.Printf("Error verifying JWT: %v", err)
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Unauthorized", "Invalid token"))
			c.Abort()
			return
		}

		_, err = repositories.NewUserRepository(db.GetDB()).FindByUUID(claims.UUID)
		if err != nil {
			log.Printf("Error finding user by UUID: %v", err)
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("Unauthorized", "Invalid token"))
			c.Abort()
			return
		}

		c.Set("user_uuid", claims.UUID)
		c.Next()
	}
}
