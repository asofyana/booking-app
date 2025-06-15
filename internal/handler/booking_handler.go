package handler

import (
	"booking-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	RoomService    services.RoomServiceInterface
	BookingService services.BookingServiceInterface
}

func NewBookingHandler(roomService services.RoomServiceInterface,
	bookingService services.BookingServiceInterface) *BookingHandler {
	return &BookingHandler{
		RoomService:    roomService,
		BookingService: bookingService,
	}
}

func (s *BookingHandler) BookingNew(c *gin.Context) {

	rooms, _ := s.RoomService.GetAllActiveRoom(c)

	c.HTML(http.StatusOK, "booking-new.html", gin.H{
		"title": "New Booking", "Rooms": rooms})

}

func (s *BookingHandler) BookingNewPost(c *gin.Context) {

	rooms, _ := s.RoomService.GetAllActiveRoom(c)

	booking, err := s.BookingService.SaveBooking(c)
	if err != nil {
		c.HTML(http.StatusOK, "booking-new.html", gin.H{
			"title": "New Booking", "Rooms": rooms, "booking": booking, "message": err.Error(), "alert": "alert-danger"})
		return
	}

	c.HTML(http.StatusOK, "booking-new.html", gin.H{
		"title": "New Booking", "Rooms": rooms, "booking": booking, "message": "Success", "alert": "alert-success"})

}
