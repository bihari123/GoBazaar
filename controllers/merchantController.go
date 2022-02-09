package controllers

import (
	"GoBazaar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MerchantRegister(c *gin.Context) {

	var newMerchant models.Merchant
	if err := c.BindJSON(&newMerchant); err != nil {
		return
	} else {
		MerchantList = append(MerchantList, newMerchant)
	}

	// code to insert into the database
	c.IndentedJSON(http.StatusCreated, MerchantList)
}

func MerchantLogin(c *gin.Context) {
	var merchantCred models.Cred
	if err := c.BindJSON(&merchantCred); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}

	// code to find the usesr in database

	// if the user is found
	// c.IndentedJSON(http.StatusCreated, userCred.UserID)

	// if the user is not found
	// c.IndentedJSON(http.StatusNotFound, gin.H("message":"enter valid credentials"))

}

func MerchantLogout(c *gin.Context) {

}

func MerchantUpload(c *gin.Context) {
	var product models.Product
	var merchant models.Merchant
	var found bool = false
	if err := c.BindJSON(&product); err != nil {
		return
	}

	for _, val := range MerchantList {
		if val.ID == product.MerchantID {
			merchant = val
			found = true
		}
	}

	if found {
		merchant.Products = append(merchant.Products, product)
		c.IndentedJSON(http.StatusCreated, merchant)
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "merchant not found"})

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
		if val.ID == product.MerchantID {
			merchant = val
			found = true
		}
	}

	if found {
		for index, val := range merchant.Products {
			if val.ID == product.ID {
				merchant.Products[index].Price = updatedPrice
				c.IndentedJSON(http.StatusCreated, merchant)
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "merchant not found"})

}
