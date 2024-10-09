package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/paghapour/golang-clean-web-api/api/handlers"
	"github.com/paghapour/golang-clean-web-api/api/middlewares"
	"github.com/paghapour/golang-clean-web-api/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}