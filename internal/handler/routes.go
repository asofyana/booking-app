package handler

import (
	"booking-app/internal/services"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, handler *HomeHandler) {
	authRouter := router.Group("", services.Auth)
	authRouter.GET("/home", handler.ShowHomePage)
}

func RegisterRouteAdmin(router *gin.Engine, handler *AdminHandler) {
	adminRouter := router.Group("/admin", services.Auth)
	adminRouter.GET("", handler.ShowAdminPage)
	adminRouter.POST("/users", handler.CreateUser)
	adminRouter.PUT("/users", handler.UpdateUser)
	adminRouter.DELETE("/users/:id", handler.DeleteUser)
}

func RegisterRouteBooking(router *gin.Engine, handler *BookingHandler) {
	bookingRouter := router.Group("/booking", services.BookingAuth)
	bookingRouter.GET("/home", handler.BookingHome)
	bookingRouter.GET("/create", handler.BookingNew)
	bookingRouter.POST("/create", handler.BookingNewPost)
	bookingRouter.GET("/approval", handler.BookingApproval)
	bookingRouter.POST("/approval", handler.BookingApprovalPost)
}

func RegisterRouteLogin(app *gin.Engine, handler *LoginHandler) {
	app.GET("/", handler.ShowLoginPage)
	app.GET("/login", handler.ShowLoginPage)
	app.POST("/login", handler.ProcessLogin)
	app.GET("/logout", handler.ProcessLogout)
}
