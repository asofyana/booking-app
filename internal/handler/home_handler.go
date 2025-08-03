package handler

import (
	"booking-app/internal/services"
	"fmt"
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
	fmt.Println("zzzzzz: ", user.Role)

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page", "User": user})
}
