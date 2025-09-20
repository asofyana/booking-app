package handler

import (
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	BookingService services.BookingServiceInterface
}

func NewHomeHandler(bookingService services.BookingServiceInterface) *HomeHandler {
	return &HomeHandler{
		BookingService: bookingService,
	}
}

func (s *HomeHandler) ShowHomePage(c *gin.Context) {
	user := services.GetUserSession(c)
	bookings, _ := s.BookingService.GetAllIncomingBookingDashboard()

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page", "User": user, "bookings": bookings})
}
