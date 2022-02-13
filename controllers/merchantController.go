package controllers

import (
	"GoBazaar/database"
	"GoBazaar/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MerchantRegister(c *gin.Context) {

	var newMerchant models.Merchant
	if err := c.BindJSON(&newMerchant); err != nil {
		return
	} else {
		//newMerchant.Credentials.ID = uuid.New().String()
		newMerchant.Credentials.Pass = "Pass123"

	}

	result, err := database.Db.Exec("INSERT INTO merchant (company_name,email,merchant_address,discount_offered)VALUES(?,?,?,?)", newMerchant.CompanyName, newMerchant.Email, newMerchant.MerchantAddress, newMerchant.DiscountOffered)

	if err != nil {

		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error adding into database merchant"})
		return
	}
	id, err := result.LastInsertId()

	newMerchant.Credentials.ID = int(id)

	result, err = database.Db.Exec("INSERT INTO merchCreds (id,pass) VALUES (?, ?)", newMerchant.Credentials.ID, newMerchant.Credentials.Pass)
	if err != nil {
		c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error adding into database merchantCred"})
		return
	}

	// code to insert into the database
	c.IndentedJSON(http.StatusCreated, newMerchant)
}

func MerchantLogin(c *gin.Context) {
	var merchantCred models.MerchCred
	if err := c.BindJSON(&merchantCred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	// code to find the usesr in database
	row := database.Db.QueryRow("SELECT * FROM merchCreds WHERE id = ? AND pass = \"?\"", merchantCred.ID, merchantCred.Pass)
	fmt.Println(row)
	if err := row.Scan(&merchantSession.Credentials.ID, &merchantSession.Credentials.Pass); err != nil {
		if err == sql.ErrNoRows {
			// if the user is not found
			fmt.Println("\n\n\n")
			fmt.Println(err)
			fmt.Println("\n\n\n")
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Merchant not found "})
			return
		}
	}

	row = database.Db.QueryRow("SELECT * FROM merchant WHERE id = ?", merchantCred.ID)

	if err := row.Scan(&merchantSession.Credentials.ID, &merchantSession.CompanyName, &merchantSession.Email, &merchantSession.MerchantAddress, &merchantSession.DiscountOffered); err != nil {
		if err == sql.ErrNoRows {
			// if the merchant is not found
			fmt.Println("\n\n\n")
			fmt.Println(err)
			fmt.Println("\n\n\n")

			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "merchant not found"})
			return
		}
	}
	c.IndentedJSON(http.StatusFound, gin.H{"message": "login successful.\nWelcome " + merchantSession.CompanyName})
	return

}

func MerchantLogout(c *gin.Context) {
	if merchantSession.Credentials.ID != 0 {
		merchantSession = models.Merchant{}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Session deleted. User Logged out"})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Session empty. Plz Log in first"})
}

func MerchantUpload(c *gin.Context) {
	var product models.Product

	if err := c.BindJSON(&product); err != nil {
		fmt.Println("\n\n\n")
		fmt.Println(err)
		fmt.Println("\n\n\n")

		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	if merchantSession.CompanyName != "" {
		product.MerchantID = merchantSession.Credentials.ID
		_, err := database.Db.Exec("INSERT INTO product (merchantID,namel,product_description,price,stock) VALUES (?, ?, ?, ?, ?) ", product.MerchantID, product.Name, product.ProductDescription, product.Price, product.Stock)

		if err != nil {

			c.IndentedJSON(http.StatusExpectationFailed, gin.H{"message": "error adding into database product"})
			return
		}

		merchantSession.Products = append(merchantSession.Products, product)

		c.IndentedJSON(http.StatusCreated, gin.H{"message": "product uploaded. Dhanda Kaayam rhe!!!"})
		return

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No merchant logged in"})
		return
	}

}

func MerchantUpdate(c *gin.Context) {
	var product models.Product
	var merchant models.Merchant
	var found bool = false
	if err := c.BindJSON(&product); err != nil {
		return
	}
	updatedPrice := product.Price

	for _, val := range MerchantList {
		if val.Credentials.ID == product.MerchantID {
			merchant = val
			found = true
		}
	}

	if found {
		for index, val := range merchant.Products {
			if val.ProductID == product.ProductID {
				merchant.Products[index].Price = updatedPrice
				c.IndentedJSON(http.StatusFound, merchant)
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "merchant not found"})

}
