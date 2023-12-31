package routes

import (
	"golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
