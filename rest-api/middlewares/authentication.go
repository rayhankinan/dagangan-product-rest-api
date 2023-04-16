package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	authConfig "dagangan-product-rest-api/config/authentication"
	"dagangan-product-rest-api/repositories"
	"dagangan-product-rest-api/types"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	ID   uint           `json:"id"`
	Role types.AuthRole `json:"role"`
}

// Private
func containRole(s []types.AuthRole, str types.AuthRole) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Public
func AuthMiddleware(roles ...types.AuthRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		config := authConfig.Config.GetMetadata()
		response := repositories.Response[string]{}

		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response.Message = "ERROR: NO TOKEN PROVIDED"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		authString := strings.Replace(authHeader, "Bearer ", "", -1)
		authClaim := AuthClaims{}
		authToken, err := jwt.ParseWithClaims(authString, &authClaim, func(authToken *jwt.Token) (interface{}, error) {
			if method, ok := authToken.Method.(*jwt.SigningMethodHMAC); !ok || method != config.JWTSigningMethod {
				return nil, fmt.Errorf("ERROR: SIGNING METHOD INVALID")
			}
			return config.JWTSignatureKey, nil
		})
		if err != nil {
			response.Message = "ERROR: TOKEN CANNOT BE PARSED"
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		if !authToken.Valid {
			response.Message = "ERROR: CLAIMS INVALID"
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if len(roles) > 0 && !containRole(roles, authClaim.Role) {
			response.Message = "ERROR: UNAUTHORIZED"
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("user_id", authClaim.ID)
		c.Set("role", authClaim.Role)
		c.Next()
	}
}
