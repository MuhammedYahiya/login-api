package main

import (
	"login-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.InitializeRoutes(r)
	r.Run(":8080")
}
