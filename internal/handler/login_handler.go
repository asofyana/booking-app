package handler

import (
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService   services.LoginServiceInterface
	BookingService services.BookingServiceInterface
}

func NewLoginHandler(loginService services.LoginServiceInterface, bookingService services.BookingServiceInterface) *LoginHandler {
	return &LoginHandler{
		LoginService:   loginService,
		BookingService: bookingService,
	}
}

func (s *LoginHandler) ShowLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login",
		},
	)
}

func (s *LoginHandler) ProcessLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := s.LoginService.ProcessLogin(c, username, password)

	if err == nil {
		bookings, _ := s.BookingService.GetAllIncomingBookingDashboard()

		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Home Page", "User": user, "bookings": bookings, "LastIndex": len(bookings) - 1})

		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"title": "Login", "message": "Invalid username or password"})

}

func (s *LoginHandler) ProcessLogout(c *gin.Context) {
	services.InvalidateSession(c)
	c.Redirect(http.StatusTemporaryRedirect, "/")
	//c.HTML(http.StatusOK, "login.html", nil)
}
