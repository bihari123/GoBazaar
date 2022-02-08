package handlers

import (
	"GoBazaar/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMerchant(c *gin.Context) {
	controllers.MerchantRegister(c)
}

func LoginMerchant(c *gin.Context) {
	controllers.MerchantLogin(c)
}

func LogoutMerchant(c *gin.Context) {
	controllers.MerchantLogout(c)
}

func UploadMerchant(c *gin.Context) {
	controllers.MerchantUpload(c)
}

func UpdateMerchant(c *gin.Context) {
	controllers.MerchantUpdate(c)
}
