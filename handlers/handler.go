package handlers

import (
	"GoBazaar/controllers"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	message := controllers.Welcome(c)
	c.IndentedJSON(200, message)
}
