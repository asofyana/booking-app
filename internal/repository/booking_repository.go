package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"log"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		DB: db,
	}
}

type BookingRepositoryInterface interface {
	GetAllBooking() ([]entity.Booking, error)
	InsertBooking(booking entity.Booking) error
	GetOverlapBookingCount(booking entity.Booking) int
}

func (s *BookingRepository) GetAllBooking() ([]entity.Booking, error) {
	return nil, nil
}

func (s *BookingRepository) InsertBooking(booking entity.Booking) error {
	_, err := s.DB.Exec("INSERT INTO booking (title,room_id,booking_description,start_date,end_date,participant_count,created_by) VALUES (?,?,?,?,?,?,?)",
		booking.Title, booking.RoomId, booking.Description, booking.StartDate, booking.EndDate, booking.ParticipantCount, booking.CreatedBy)

	if err != nil {
		log.Println("Error inserting person: ", err.Error())
		return err
	}

	return nil
}

func (s *BookingRepository) GetOverlapBookingCount(booking entity.Booking) int {
	var count int
	err := s.DB.QueryRow("select count(1) from booking where room_id = ? and start_date < ? and end_date > ?", booking.RoomId, booking.EndDate, booking.StartDate).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}
