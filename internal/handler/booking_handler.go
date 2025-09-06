package handler

import (
	"booking-app/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	RoomService    services.RoomServiceInterface
	BookingService services.BookingServiceInterface
	LookupService  services.LookupServiceInterface
}

func NewBookingHandler(roomService services.RoomServiceInterface,
	bookingService services.BookingServiceInterface, lookupService services.LookupServiceInterface) *BookingHandler {
	return &BookingHandler{
		RoomService:    roomService,
		BookingService: bookingService,
		LookupService:  lookupService,
	}
}

func (s *BookingHandler) BookingNew(c *gin.Context) {

	rooms, _ := s.RoomService.GetAllActiveRoom(c)
	activities, _ := s.LookupService.GetAllActiveLookupByType(c, "ACTIVITY")
	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "booking-new.html", gin.H{
		"title": "New Booking", "User": user, "Rooms": rooms, "Activities": activities})

}

func (s *BookingHandler) BookingNewPost(c *gin.Context) {

	rooms, _ := s.RoomService.GetAllActiveRoom(c)
	activities, _ := s.LookupService.GetAllActiveLookupByType(c, "ACTIVITY")
	user := services.GetUserSession(c)

	booking, err := s.BookingService.SaveBooking(c)
	if err != nil {
		c.HTML(http.StatusOK, "booking-new.html", gin.H{
			"title": "New Booking", "User": user, "Rooms": rooms, "Activities": activities, "booking": booking, "message": err.Error(), "alert": "alert-danger"})
		return
	}

	c.HTML(http.StatusOK, "booking-new.html", gin.H{
		"title": "New Booking", "User": user, "Rooms": rooms, "Activities": activities, "booking": booking, "message": "Success", "alert": "alert-success"})

}

func (s *BookingHandler) BookingHome(c *gin.Context) {

	userIncomingBooking, _ := s.BookingService.GetIncomingBookingByUsername(c)
	allIncomingBooking, _ := s.BookingService.GetAllIncomingBooking(c)
	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "booking-home.html", gin.H{
		"title": "New Booking", "User": user, "UserIncomingBooking": userIncomingBooking, "AllIncomingBooking": allIncomingBooking})

}

func (s *BookingHandler) BookingApproval(c *gin.Context) {

	bookingIdStr, _ := c.GetQuery("id")
	bookingIdInt, _ := strconv.Atoi(bookingIdStr)
	booking, _ := s.BookingService.GetBookingById(c, bookingIdInt)
	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "booking-approval.html", gin.H{
		"title": "Booking Approval", "User": user, "booking": booking})

}

func (s *BookingHandler) BookingApprovalPost(c *gin.Context) {

	bookingIdStr := c.Request.FormValue("bookingId")
	bookingIdInt, _ := strconv.Atoi(bookingIdStr)
	booking, err := s.BookingService.GetBookingById(c, bookingIdInt)
	user := services.GetUserSession(c)
	message := ""
	alert := "alert-danger"

	if err != nil {
		fmt.Println("Error: ", err.Error())
		message = "Invalid Booking"
	}

	if user.Role != "admin" {
		message = "Invalid User Role"
	}

	if message == "" {
		action := c.Request.FormValue("btnAction")
		if action == "Approve" {
			booking.Status = "Approved"
		} else if action == "Reject" {
			booking.Status = "Rejected"
		}
		err := s.BookingService.UpdateBookingStatus(c, bookingIdInt, booking.Status)
		if err == nil {
			message = "Success"
			alert = "alert-success"
		} else {
			message = "Failed to update booking"
		}
	}

	c.HTML(http.StatusOK, "booking-approval.html", gin.H{
		"title": "Booking Approval", "User": user, "booking": booking, "message": message, "alert": alert})

}

func (s *BookingHandler) BookingView(c *gin.Context) {

	bookingIdStr, _ := c.GetQuery("id")
	bookingIdInt, _ := strconv.Atoi(bookingIdStr)
	booking, _ := s.BookingService.GetBookingById(c, bookingIdInt)
	user := services.GetUserSession(c)

	c.HTML(http.StatusOK, "booking-view.html", gin.H{
		"title": "Booking Approval", "User": user, "booking": booking})

}

func (s *BookingHandler) BookingViewPost(c *gin.Context) {

	bookingIdStr := c.Request.FormValue("bookingId")
	bookingIdInt, _ := strconv.Atoi(bookingIdStr)
	booking, err := s.BookingService.GetBookingById(c, bookingIdInt)
	user := services.GetUserSession(c)
	message := ""
	alert := "alert-danger"

	if err != nil {
		fmt.Println("Error: ", err.Error())
		message = "Invalid Booking"
	}

	if message == "" {
		action := c.Request.FormValue("btnAction")
		if action == "Cancel" {
			booking.Status = "Approved"
			err := s.BookingService.UpdateBookingStatus(c, bookingIdInt, "Canceled")
			if err == nil {
				message = "Success"
				alert = "alert-success"
			} else {
				message = "Failed to update booking"
			}

		}
	}

	c.HTML(http.StatusOK, "booking-view.html", gin.H{
		"title": "Booking Approval", "User": user, "booking": booking, "message": message, "alert": alert})

}
