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

	var product models.Product
	var merchant models.Merchant
	if userSession.FirstName != "" {
		if err := c.BindJSON(&product); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
		}

		row := database.Db.QueryRow("SELECT * FROM product WHERE id = ?", product.ProductID)

		if err := row.Scan(&product.ProductID, &product.MerchantID, &product.Name, &product.ProductDescription, &product.Price, &product.Stock); err != nil {
			if err == sql.ErrNoRows {
				// if the user is not found
				fmt.Println("\n\n\n")
				fmt.Println(err)
				fmt.Println("\n\n\n")
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
				return
			}
		}

		row = database.Db.QueryRow("SELECT * FROM merchant WHERE id = ?", product.MerchantID)

		if err := row.Scan(&merchant.Credentials.ID, &merchant.CompanyName, &merchant.Email, &merchant.MerchantAddress, &merchant.DiscountOffered); err != nil {
			if err == sql.ErrNoRows {
				// if the user is not found
				fmt.Println("\n\n\n")
				fmt.Println(err)
				fmt.Println("\n\n\n")
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "merchant not found"})
				return
			}
		}
		discount := merchant.DiscountOffered
		netPrice := product.Price - discount

		walletBalace := userSession.WalletBalance

		if walletBalace > netPrice {
			userSession.WalletBalance = walletBalace - netPrice
			_, err := database.Db.Exec("UPDATE user SET wallet_balance = ? WHERE id = ?", userSession.WalletBalance, userSession.Credentials.ID)
			product.Stock -= 1
			if err != nil {
				c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error updating the user wallet balance"})
				return
			}
			_, err = database.Db.Exec("UPDATE product SET stock= ? WHERE id = ?", product.Stock, product.ProductID)

			if err != nil {
				c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error updating the product stock"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Item Purchased"})
			return
		}
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Not Enough Balance. Plz recharge your wallet"})
		return

	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "there is no active session for user"})
	return

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
