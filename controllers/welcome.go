package controllers

import (
	"GoBazaar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	msg := models.WelcomeMsg{
		Msg: "Welcome to GoBazaar! Your own desi shop",
	}

	c.IndentedJSON(http.StatusOK, msg)
}
