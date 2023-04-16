package routes

import (
	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/controllers"
	"dagangan-product-rest-api/middlewares"
	"dagangan-product-rest-api/types"
	"dagangan-product-rest-api/utils/cache"
)

func ProductRoute(route *gin.Engine) {
	productGroup := route.Group("/product")

	productGroup.GET("/:id", middlewares.AuthMiddleware(types.Admin, types.Viewer), cache.Store.GetHandlerFunc(controllers.GetProductHandler()))
	productGroup.GET("/", middlewares.AuthMiddleware(types.Admin, types.Viewer), cache.Store.GetHandlerFunc(controllers.SearchProductHandler()))
	productGroup.POST("/", middlewares.AuthMiddleware(types.Admin), controllers.AddProductHandler())
	productGroup.PUT("/:id", middlewares.AuthMiddleware(types.Admin), controllers.EditProductHandler())
	productGroup.DELETE("/:id", middlewares.AuthMiddleware(types.Admin), controllers.RemoveProductHandler())
}
