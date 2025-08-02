package routes

import (
	"encoding/json"
	"io"
	"log"
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

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author,omitempty"`
	ISBN   string `json:"isbn,omitempty"`
	// Puedes agregar más campos según la respuesta real
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

	r.GET("/potterapibooks", func(c *gin.Context) {
		resp, err := http.Get("https://potterapi-fedeperin.vercel.app/en/books")
		if err != nil {
			log.Printf("Error al hacer request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hacer la solicitud"})
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			bodyBytes, _ := io.ReadAll(resp.Body)
			c.JSON(resp.StatusCode, gin.H{
				"error":  "Respuesta no exitosa",
				"status": resp.StatusCode,
				"body":   string(bodyBytes),
			})
			return
		}

		var books []Book
		if err := json.NewDecoder(resp.Body).Decode(&books); err != nil {
			log.Printf("Error al parsear JSON: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear JSON"})
			return
		}

		c.JSON(http.StatusOK, books)
	})

}
