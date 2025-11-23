package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
	"booking-app/internal/utils"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingService struct {
	BookingRepository repository.BookingRepositoryInterface
}

func NewBookingService(bookingRepository repository.BookingRepositoryInterface) *BookingService {
	return &BookingService{
		BookingRepository: bookingRepository,
	}
}

type BookingServiceInterface interface {
	SaveBooking(c *gin.Context) (entity.Booking, error)
	GetIncomingBookingByUsername(c *gin.Context) ([]entity.Booking, error)
	GetAllIncomingBooking(c *gin.Context) ([]entity.Booking, error)
	GetBookingById(c *gin.Context, bookingId int) (entity.Booking, error)
	UpdateBookingStatus(c *gin.Context, bookingId int, status string) error
	SearchBooking(entity.Booking) ([]entity.Booking, error)
	GetAllIncomingBookingDashboard() ([]entity.Booking, error)
	RejectBooking(c *gin.Context, bookingId int, rejectReason string) error
}

func (s *BookingService) SaveBooking(c *gin.Context) (entity.Booking, error) {
	title := c.PostForm("title")
	startDate := c.PostForm("start-date")
	startTime := c.PostForm("start-time")
	endDate := c.PostForm("end-date")
	endTime := c.PostForm("end-time")
	participantCount, _ := strconv.Atoi(c.PostForm("number-of-participant"))
	activityCode := c.PostForm("activity")
	organizer := c.PostForm("organizer")
	pic := c.PostForm("pic")
	pic_contactno := c.PostForm("pic_contactno")

	user := GetUserSession(c)

	booking := entity.Booking{
		Title:            title,
		ParticipantCount: participantCount,
		CreatedBy:        user.UserName,
		Status:           "Pending",
		Activity:         activityCode,
		Organizer:        organizer,
		Pic:              pic,
		PicContactNo:     pic_contactno,
	}

	layoutFormat := "2006-01-02 15:04"
	startDateTime, err := time.Parse(layoutFormat, startDate+" "+startTime)
	if err != nil {
		log.Println("Error Parsing startDate: ", err.Error())
		return booking, errors.New("invalid start date")
	}

	booking.StartDate = startDateTime

	endDateTime, err := time.Parse(layoutFormat, endDate+" "+endTime)
	if err != nil {
		log.Println("Error Parsing endDate: ", err.Error())
		return booking, errors.New("invalid end date")
	}

	booking.EndDate = endDateTime

	if startDateTime.After(endDateTime) || startDateTime.Equal(endDateTime) {
		return booking, errors.New("invalid date")
	}

	now := time.Now()
	if startDateTime.Before(now) {
		return booking, errors.New("invalid date")
	}

	// Get booking room
	selectedRooms := c.PostFormArray("rooms")
	var rooms []entity.Room
	for _, roomId := range selectedRooms {
		var room entity.Room
		room.RoomId, _ = strconv.Atoi(roomId)
		rooms = append(rooms, room)
	}
	booking.Rooms = rooms

	if len(rooms) == 0 {
		return booking, errors.New(utils.Translate("booking_err_select_room", nil))
	}

	overlapCount := s.BookingRepository.GetOverlapBookingCount(booking)
	if overlapCount == -1 {
		return booking, errors.New("Error")
	}
	if overlapCount > 0 {
		return booking, errors.New(utils.Translate("booking_err_overlap", nil))
	}

	err2 := s.BookingRepository.InsertBooking(booking)
	if err2 != nil {
		log.Println("Error Insert Booking: ", err2.Error())
		return booking, errors.New("error save booking")
	}

	return booking, nil
}

func (s *BookingService) GetIncomingBookingByUsername(c *gin.Context) ([]entity.Booking, error) {
	user := GetUserSession(c)
	return s.BookingRepository.GetIncomingBookingByUsername(user.UserName)
}

func (s *BookingService) GetAllIncomingBooking(c *gin.Context) ([]entity.Booking, error) {
	return s.BookingRepository.GetAllIncomingBooking()
}

func (s *BookingService) GetBookingById(c *gin.Context, bookingId int) (entity.Booking, error) {
	return s.BookingRepository.GetBookingById(bookingId)
}

func (s *BookingService) UpdateBookingStatus(c *gin.Context, bookingId int, status string) error {
	return s.BookingRepository.UpdateBookingStatus(bookingId, status)
}

func (s *BookingService) SearchBooking(booking entity.Booking) ([]entity.Booking, error) {
	return s.BookingRepository.SearchBooking(booking)
}

func (s *BookingService) GetAllIncomingBookingDashboard() ([]entity.Booking, error) {
	return s.BookingRepository.GetAllIncomingBookingDashboard()
}

func (s *BookingService) RejectBooking(c *gin.Context, bookingId int, rejectReason string) error {
	return s.BookingRepository.RejectBooking(bookingId, rejectReason)
}
