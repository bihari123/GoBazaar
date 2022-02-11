package main

import (
	"GoBazaar/database"
	"GoBazaar/handlers"

	"github.com/gin-gonic/gin"
)

const (
	Home     = "/"
	User     = "/user"
	Merchant = "/merchant"
	Register = "/register"
	Login    = "/login"
	Logout   = "/logout"
	Upload   = "/upload"
	Purchase = "/purchase"
	Update   = "/update"
	Stock    = "/stock"
	Cart     = "/cart"
	Id       = "/:id"
	Search   = "/search/:letter1/:letter2/:letter3"
)

func init() {
	database.RegisterDatabase()
}
func main() {
	router := gin.Default()

	router.GET(Home, handlers.HomeHandler)

	router.POST(User+Register, handlers.RegisterUser)
	router.GET(User+Cart+Id, handlers.ShowUserCart)
	router.POST(User+Login, handlers.LoginUser)
	router.GET(User+Logout, handlers.LogoutUser)
	router.GET(User+Search, handlers.SearchUser)
	router.POST(User+Purchase, handlers.PurchaseUser)

	router.POST(Merchant+Register, handlers.RegisterMerchant)
	router.POST(Merchant+Login, handlers.LoginMerchant)
	router.POST(Merchant+Logout, handlers.LogoutMerchant)
	router.POST(Merchant+Upload, handlers.UploadMerchant)
	router.POST(Merchant+Update, handlers.UpdateMerchant)
	router.GET(Merchant+Stock+Id, handlers.UpdateMerchant)

	router.Run("localhost:8080")
}
