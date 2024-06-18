package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/brayanzuritadev/dailyapi/cmd/api/handlers/user"
)

func SetRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("", user.CreateUser)
	}
}