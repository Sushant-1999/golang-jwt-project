package routes

import (
	"golang-jwt-project/controllers"
	"golang-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

// * Before using users APIS we need to authenticate the User first
// * These user Routes are protected and private and Before calling them we need to authenticate the User first. Hence we are using the Middleware and make request authenticated.
// * SignUp and Login are public routes. 
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
