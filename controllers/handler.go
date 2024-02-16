package controllers

import (
	"login-api/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	for _, u := range models.Users {
		if u.Username == user.Username && u.Password == user.Password {
			session.Set("username", user.Username)
			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"message":    "login successful",
				"session_id": session.Get("username"),
			})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid username or password",
	})
}
