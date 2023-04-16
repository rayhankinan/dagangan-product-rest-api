package routes

import (
	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/controllers"
)

func NotFoundRoute(route *gin.Engine) {
	route.NoRoute(controllers.NotFoundHanlder())
}
