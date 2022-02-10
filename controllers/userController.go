package controllers

import (
	"GoBazaar/models"
	"net/http"
	"strings"

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
			session = val
			c.IndentedJSON(http.StatusFound, gin.H{"message": "User logged in. Session created"})
			return

		}
	}

	// if the user is not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})

}

func UserLogout(c *gin.Context) {
	if session.Credentials.UserID != "" {
		session = models.User{}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Session deleted. User Logged out"})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Session empty. Plz Log in first"})
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
	// know how much discount is being offered by the merchant on that product, deduct the discount from the MRP
	// deduct the amount from the user wallet
	// reduce the stock of the product by one

	var product models.Product
	//var merchant models.Merchant
	var discount float64

	if session.Credentials.UserID != "" {
		if err := c.BindJSON(&product); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		for _, val := range ProductList {
			if val.ID == product.ID {
				for _, m := range MerchantList {
					if m.Credentials.UserID == product.MerchantID {
						discount = m.DiscountOffered
						netPrice := product.Price - discount
						if session.WalletBalance > netPrice {
							session.WalletBalance = session.WalletBalance - netPrice
							c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Thanks for the purchase. Plz visit again"})
							return
						}

						c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not enough balance. Plz recharge your wallet"})
						return

					}
				}
			}
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "product not found"})
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Session Emplty.Plz log in first"})

}

//make the user enter three letters
//return the list of products whose name or description has those letters

func UserSearch(c *gin.Context) {
	letter1 := c.Param("letter1")
	letter2 := c.Param("letter2")
	letter3 := c.Param("letter3")

	var searchResult []models.Product

	for _, val := range ProductList {
		if strings.Contains(val.Name, letter1) && strings.Contains(val.Name, letter2) && strings.Contains(val.Name, letter3) {
			searchResult = append(searchResult, val)
		}

		if strings.Contains(val.Description, letter1) && strings.Contains(val.Description, letter2) && strings.Contains(val.Description, letter3) {
			searchResult = append(searchResult, val)
		}

		if strings.Contains(val.Name, letter1) || strings.Contains(val.Name, letter2) || strings.Contains(val.Name, letter3) {
			searchResult = append(searchResult, val)
		}
	}

	if len(searchResult) > 0 {
		c.IndentedJSON(http.StatusFound, searchResult)
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no result found"})
	}

}
