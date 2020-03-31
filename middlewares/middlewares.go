package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"rest-api-gin-gonic-mongo/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)


//Authentication is for auth middleware
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			c.JSON(402, gin.H{
				"error": "Unauthorized Access !",
			})
			return
		}
		
		tokenString := strings.TrimSpace(authHeader)
		fmt.Println("tokenString is ", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			
			secretKey := "qwertyy"
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(402, gin.H{
				"error": "Unauthorized Access !",
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			fmt.Println("email is ", email)

			user,err := userModel.GetEmail(email)
	
			if err != nil {
				c.JSON(402, gin.H{
				"error": "Unauthorized Access !",
				})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(402, gin.H{
				"error": "Unauthorized Access !",
			})
		}
	}
}

//ErrorHandler is for global error
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}
