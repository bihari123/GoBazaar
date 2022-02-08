package controllers

import (
	"GoBazaar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// code to insert into the database
	c.IndentedJSON(http.StatusCreated, newUser)
}

func UserLogin(c *gin.Context) {
	var userCred models.Cred
	if err := c.BindJSON(&userCred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	// code to find the usesr in database

	// if the user is found
	// c.IndentedJSON(http.StatusCreated, userCred.UserID)

	// if the user is not found
	// c.IndentedJSON(http.StatusNotFound, gin.H("message":"enter valid credentials"))

}

func UserLogout(c *gin.Context) {

}

func UserCart(c *gin.Context) {
	user_id := c.Param("id")

	// get the user cart

	// if the user exist
	//c.IndentedJSON(http.StatusOK,user_chart)

	// if the user doesn't exist
	//c.IdentedJSON(http.StatusNotFound,gin.H("message":"user not found"))
}
