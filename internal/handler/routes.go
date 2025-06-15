package handler

import (
	"booking-app/internal/services"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	authRouter := router.Group("", services.Auth)
	authRouter.GET("/home", ShowHomePage)

}

func RegisterRouteBooking(router *gin.Engine, handler BookingHandler) {
	bookingRouter := router.Group("/booking", services.BookingAuth)
	bookingRouter.GET("/create", handler.BookingNew)
	bookingRouter.POST("/create", handler.BookingNewPost)
}

func RegisterRouteLogin(app *gin.Engine, handler *LoginHandler) {
	app.GET("/", handler.ShowLoginPage)
	app.GET("/login", handler.ShowLoginPage)
	app.POST("/login", handler.ProcessLogin)
	app.GET("/logout", handler.ProcessLogout)
}
