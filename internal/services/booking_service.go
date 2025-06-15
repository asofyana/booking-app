package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
	"errors"
	"fmt"
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

	user := GetUserSession(c)
	fmt.Println("Save Booking:", user)

	booking := entity.Booking{
		Title:            title,
		Description:      description,
		RoomId:           roomId,
		ParticipantCount: participantCount,
		CreatedBy:        user.UserName,
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
