package controllers

import (
	"GoBazaar/database"
	"GoBazaar/models"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, UserList)
	} else {
		newUser.Credentials.Pass = "Pass123"
	}

	result, err := database.Db.Exec("INSERT INTO user (first_name, last_name,email,contact,city,wallet_balance) VALUES (?, ?, ?, ?, ?, ?)", newUser.FirstName, newUser.LastName, newUser.Email, newUser.Contact, newUser.City, newUser.WalletBalance)
	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error adding into database user"})
		return
	}
	id, err := result.LastInsertId()

	newUser.Credentials.ID = int(id)
	result, err = database.Db.Exec("INSERT INTO userCreds (id,pass) VALUES (?, ?)", newUser.Credentials.ID, newUser.Credentials.Pass)
	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error adding into database userCred"})
		return
	}

	// code to insert into the database
	c.IndentedJSON(http.StatusCreated, newUser)
}

func UserLogin(c *gin.Context) {
	var userCred models.UserCred
	if err := c.BindJSON(&userCred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	// code to find the usesr in database

	row := database.Db.QueryRow("SELECT * FROM userCreds WHERE id = ? AND pass = \"?\"", userCred.ID, userCred.Pass)
	fmt.Println(row)
	if err := row.Scan(&userSession.Credentials.ID, &userSession.Credentials.Pass); err != nil {
		if err == sql.ErrNoRows {
			// if the user is not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
	}

	row = database.Db.QueryRow("SELECT * FROM user WHERE id = ?", userCred.ID)

	if err := row.Scan(&userSession.Credentials.ID, &userSession.FirstName, &userSession.LastName, &userSession.Email, &userSession.Contact, &userSession.City, &userSession.WalletBalance); err != nil {
		if err == sql.ErrNoRows {
			// if the user is not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
	}

	c.IndentedJSON(http.StatusFound, gin.H{"message": "login successful.\nWelcome " + userSession.FirstName})
	return

}

func UserLogout(c *gin.Context) {
	if userSession.Credentials.ID != 0 {
		userSession = models.User{}
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

	//var product models.Product
	//var merchant models.Merchant
	//var discount float64
	/*
		if session.Credentials.ID != "" {
			if err := c.BindJSON(&product); err != nil {
				c.IndentedJSON(http.StatusBadRequest, err)
			}

			for _, val := range ProductList {
				if val.ProductID == product.ProductID {
					for _, m := range MerchantList {
						if m.Credentials.ID == product.MerchantID {
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
	*/
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

		if strings.Contains(val.ProductDescription, letter1) && strings.Contains(val.ProductDescription, letter2) && strings.Contains(val.ProductDescription, letter3) {
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
