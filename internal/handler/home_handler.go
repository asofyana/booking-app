package handler

import (
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	BookingService services.BookingServiceInterface
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (s *HomeHandler) ShowHomePage(c *gin.Context) {

	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page", "User": user})
}
