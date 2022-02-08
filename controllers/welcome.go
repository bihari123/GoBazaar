package controllers

import (
	"GoBazaar/structs"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) structs.WelcomeMsg {
	msg := structs.WelcomeMsg{
		Msg: "Welcome to GoBazaar! Your own desi shop",
	}

	fmt.Println(msg)
	return msg
}
