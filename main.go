package main

import (
	"rest-api-gin-gonic-mongo/controllers"
	"rest-api-gin-gonic-mongo/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		movie := new(controllers.MovieController)
		userAuth := middlewares.Authentication()
		v1.POST("/movies", movie.Create)
		v1.GET("/movies", movie.Find)
		v1.GET("/movies/:id", movie.Get)
		v1.PUT("/movies/:id", movie.Update)
		v1.DELETE("/movies/:id", movie.Delete)
		user := new(controllers.UserAuthController)
		v1.POST("/register", user.Register)
		v1.POST("/login", user.Login)
		v1.GET("/profile", userAuth, user.Profile)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
