package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SetupRoutes(r *gin.Engine) {
	// listen and serve on 0.0.0.0:8080
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{Message: "Hello y!"})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, Response{Message: "Hello " + name + "!"})
	})

	r.POST("/user", func(c *gin.Context) {
		var body User
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if body.Name == "" || body.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
			return
		}
		c.JSON(http.StatusOK, Response{Message: "Hello " + body.Name + body.Email + "!"})
	})

}
