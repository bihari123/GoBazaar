package main

import (
	"GoBazaar/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", handlers.HomeHandler)
	router.Run("localhost:8080")
}
