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
	adminRouter := router.Group("/admin", services.AdminAuth)
	adminRouter.GET("/user-search", handler.UserSearch)
	adminRouter.POST("/user-search", handler.UserSearchPost)
	adminRouter.GET("/user-detail", handler.UserDetail)
	adminRouter.POST("/user-detail", handler.UserDetailPost)
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

func RegisterRouteUser(router *gin.Engine, handler *UserHandler) {
	userRouter := router.Group("/user", services.Auth)
	userRouter.GET("/change-password", handler.ShowChangePasswordPage)
	userRouter.POST("/change-password", handler.ProcessChangePassword)
}
