package handlers

import (
	"GoBazaar/controllers"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	controllers.Welcome(c)

}
