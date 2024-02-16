package main

import (
	"login-api/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))
	routes.InitializeRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
