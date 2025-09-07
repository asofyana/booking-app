package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
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
}

func (s *BookingService) SaveBooking(c *gin.Context) (entity.Booking, error) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	roomId, _ := strconv.Atoi(c.PostForm("room"))
	startDate := c.PostForm("start-date")
	startTime := c.PostForm("start-time")
	endDate := c.PostForm("end-date")
	endTime := c.PostForm("end-time")
	participantCount, _ := strconv.Atoi(c.PostForm("number-of-participant"))
	activityCode := c.PostForm("activity")

	user := GetUserSession(c)

	booking := entity.Booking{
		Title:            title,
		Description:      description,
		RoomId:           roomId,
		ParticipantCount: participantCount,
		CreatedBy:        user.UserName,
		Status:           "Pending",
		Activity:         activityCode,
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

	overlapCount := s.BookingRepository.GetOverlapBookingCount(booking)
	if overlapCount > 0 {
		return booking, errors.New("date overlap with other booking")
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
