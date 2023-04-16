package routes

import (
	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/controllers"
)

func UserRoute(route *gin.Engine) {
	userGroup := route.Group("/user")

	userGroup.POST("/sign-in", controllers.SignInUserHandler())
}
