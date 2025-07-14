package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWT middleware for authentication
func jwtMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		splittedHeader := strings.Split(authHeader, " ")

		if len(splittedHeader) != 2 || splittedHeader[0] != "Bearer" {
			c.JSON(401, gin.H{
				"error": "Unable to parse 'Authorization' header",
			})
			c.Abort()
			return
		}

		token := splittedHeader[1]
		if token == "" {
			c.JSON(401, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		} else if token != secret {
			c.JSON(403, gin.H{"error": "Invalid token provided"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, 1)
	})

	// Secret endpoint with JWT authentication
	r.GET("/secret", jwtMiddleware("expected-token"), func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
