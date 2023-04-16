package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/repositories"
)

func NotFoundHanlder() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := repositories.Response[string]{}
		response.Message = "ERROR: PATH NOT FOUND"
		c.AbortWithStatusJSON(http.StatusNotFound, response)
	}
}
