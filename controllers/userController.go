package controllers

import (
	"GoBazaar/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserRegister(c *gin.Context) {

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, UserList)
	} else {
		newUser.Credentials.UserID = uuid.New().String()
		newUser.Credentials.UserPass = "Pass123"
		UserList = append(UserList, newUser)
	}

	// code to insert into the database
	c.IndentedJSON(http.StatusCreated, UserList)
}

func UserLogin(c *gin.Context) {
	var userCred models.Cred
	if err := c.BindJSON(&userCred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	// code to find the usesr in database
	for _, val := range UserList {
		if val.Credentials.UserID == userCred.UserID && val.Credentials.UserPass == userCred.UserPass {
			// if the user is found
			c.IndentedJSON(http.StatusFound, val)
			return

		}
	}

	// if the user is not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})

}

func UserLogout(c *gin.Context) {

}

func UserCart(c *gin.Context) {
	//user_id := c.Param("id")

	// get the user cart

	// if the user exist
	//c.IndentedJSON(http.StatusOK,user_chart)

	// if the user doesn't exist
	//c.IdentedJSON(http.StatusNotFound,gin.H("message":"user not found"))
}

func UserPurchase(c *gin.Context) {
	// know how much discount is being offered by the merchant on that product
	// deduct the amount from the user wallet
	// reduce the stock of the product by one
	//

}
