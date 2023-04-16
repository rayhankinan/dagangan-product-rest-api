package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	authConfig "dagangan-product-rest-api/config/authentication"
	"dagangan-product-rest-api/middlewares"
	"dagangan-product-rest-api/models"
	"dagangan-product-rest-api/repositories"
	databaseService "dagangan-product-rest-api/services/database"
	"dagangan-product-rest-api/types"
)

func SignInUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := databaseService.DB.GetConnection()
		config := authConfig.Config.GetMetadata()
		response := repositories.Response[string]{}

		request := repositories.SignInUserRequest{}
		if err := c.ShouldBindJSON(&request); err != nil {
			response.Message = "ERROR: BAD REQUEST"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		condition := models.User{Username: request.Username}
		admin := models.User{}
		if err := db.Where(&condition).Find(&admin).Error; err != nil {
			response.Message = "ERROR: INVALID USERNAME"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if err := bcrypt.CompareHashAndPassword(admin.HashedPassword, []byte(request.Password)); err != nil {
			response.Message = "ERROR: INVALID PASSWORD"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		adminClaims := middlewares.AuthClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    config.ApplicationName,
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.LoginExpirationDuration)),
			},
			ID:   admin.ID,
			Role: types.Admin,
		}

		unsignedAuthToken := jwt.NewWithClaims(config.JWTSigningMethod, adminClaims)
		signedAuthToken, err := unsignedAuthToken.SignedString(config.JWTSignatureKey)
		if err != nil {
			response.Message = "ERROR: JWT SIGNING ERROR"
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		response.Message = "SUCCESS"
		response.Data = signedAuthToken
		c.JSON(http.StatusCreated, response)
	}
}
