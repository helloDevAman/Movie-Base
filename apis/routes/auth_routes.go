package routes

import (
	"github.com/gin-gonic/gin"
)

func initAuthRoutes(router *gin.Engine) {

	// router.Handle("POST", "/auth/login")

	// router.POST("/auth/verify-mobile", func(c *gin.Context) {
	// 	var data map[string]interface{}
	// 	if err := c.BindJSON(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	// 		return
	// 	}

	// 	if mobile, ok := data["mobile"]; ok {
	// 		// Process the mobile number (e.g., send OTP, validate, etc.)
	// 		c.JSON(http.StatusOK, gin.H{"message": "Mobile number received", "mobile": mobile})
	// 	} else {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number is required"})
	// 	}
	// })

	// router.POST("/auth/verify-otp", func(c *gin.Context) {
	// 	var data map[string]interface{}
	// 	if err := c.BindJSON(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	// 		return
	// 	}

	// 	if otp, ok := data["otp"]; ok {
	// 		// Validate the OTP (this is just a placeholder, implement your own logic)
	// 		if otp == "123456" {
	// 			c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
	// 		} else {
	// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
	// 		}
	// 	} else {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP is required"})
	// 	}
	// })
}
