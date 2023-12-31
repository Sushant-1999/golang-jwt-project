package main

import (
	routes "golang-jwt-project/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//* create the routes
	router := gin.New()
	//* use the gin logging system
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//* creating sample APIS

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Success": "Access granted for api-2"})
	})

	router.Run(":" + port)

}
