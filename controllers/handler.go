package controllers

import (
	"crypto/rand"
	"fmt"
	"login-api/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
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
			// Generate a new session ID
			sessionID, err := generateSessionID()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to generate session ID",
				})
				return
			}

			// Set session data
			session.Set("username", user.Username)
			session.Set("sessionID", sessionID) // Store generated ID
			session.Save()

			// Return success response with session ID
			c.JSON(http.StatusOK, gin.H{
				"message":    "login successful",
				"session_id": sessionID,
			})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid username or password",
	})
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "logout successfully",
	})
}
