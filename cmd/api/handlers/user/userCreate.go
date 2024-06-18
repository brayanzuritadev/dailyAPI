package user

import (
	"github.com/brayanzuritadev/dailyapi/internal/domain"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var user domain.UserRequest

	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, user)
}