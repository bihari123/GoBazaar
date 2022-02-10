package handlers

import (
	"GoBazaar/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	controllers.UserRegister(c)
}

func LoginUser(c *gin.Context) {
	controllers.UserLogin(c)
}

func LogoutUser(c *gin.Context) {
	controllers.UserLogout(c)
}

func PurchaseUser(c *gin.Context) {
	controllers.UserPurchase(c)
}
func ShowUserCart(c *gin.Context) {
	controllers.UserCart(c)
}

func SearchUser(c *gin.Context) {
	controllers.UserSearch(c)
}
