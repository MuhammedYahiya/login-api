package controllers

import (
	"login-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	for _, u := range models.Users {
		if u.Username == user.Username && u.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{
				"message": "login successful",
			})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid username or password",
	})
}
