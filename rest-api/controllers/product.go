package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"dagangan-product-rest-api/models"
	"dagangan-product-rest-api/repositories"
	databaseService "dagangan-product-rest-api/services/database"
)

func GetProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databaseService.DB.GetConnection()
		response := repositories.Response[models.Product]{}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		condition := models.Product{Model: gorm.Model{ID: uint(id)}}
		product := models.Product{}
		if err := db.Where(&condition).Find(&product).Error; err != nil {
			response.Message = "ERROR: NOT FOUND"
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		response.Message = "SUCCESS"
		response.Data = product
		c.JSON(http.StatusOK, response)
		return
	}
}

func SearchProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databaseService.DB.GetConnection()
		query := repositories.SearchProductsQuery{}
		response := repositories.Response[[]models.Product]{}

		if err := c.ShouldBindQuery(&query); err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		products := []models.Product{}
		if err := db.Where("name LIKE ?", "%"+query.Search+"%").Offset((query.Page - 1) * query.Size).Limit(query.Size).Find(&products).Error; err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response.Message = "SUCCESS"
		response.Data = products
		c.JSON(http.StatusOK, response)
		return
	}
}

func AddProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databaseService.DB.GetConnection()
		request := repositories.AddProductRequest{}
		response := repositories.Response[models.Product]{}

		value, exists := c.Get("user_id")
		if !exists {
			response.Message = "UNAUTHORIZED"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := value.(uint)

		if err := c.ShouldBindJSON(&request); err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		product := models.Product{Name: request.Name, Image: request.Image, Description: request.Description, Price: request.Price, Stock: request.Stock, UserID: userID}
		if err := db.Create(&product).Error; err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response.Message = "CREATED"
		response.Data = product
		c.JSON(http.StatusCreated, response)
		return
	}
}

func EditProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databaseService.DB.GetConnection()
		request := repositories.EditProductRequest{}
		response := repositories.Response[models.Product]{}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		value, exists := c.Get("user_id")
		if !exists {
			response.Message = "UNAUTHORIZED"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := value.(uint)

		if err := c.ShouldBindJSON(&request); err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		condition := models.Product{Model: gorm.Model{ID: uint(id)}, UserID: userID}
		product := models.Product{Name: request.Name, Image: request.Image, Description: request.Description, Price: request.Price, Stock: request.Stock}
		if err := db.Where(&condition).Updates(&product).Error; err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response.Message = "SUCCESS"
		response.Data = product
		c.JSON(http.StatusOK, response)
		return
	}
}

func RemoveProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databaseService.DB.GetConnection()
		response := repositories.Response[models.Product]{}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		value, exists := c.Get("user_id")
		if !exists {
			response.Message = "UNAUTHORIZED"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := value.(uint)

		condition := models.Product{Model: gorm.Model{ID: uint(id)}, UserID: userID}
		product := models.Product{}
		if err := db.Where(&condition).Delete(&product).Error; err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response.Message = "SUCCESS"
		response.Data = product
		c.JSON(http.StatusOK, response)
		return
	}
}
