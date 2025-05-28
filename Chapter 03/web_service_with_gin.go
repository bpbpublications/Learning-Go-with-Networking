package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Golang Web Services with Gin!")
	})

	router.Run(":8080")
}
