package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"errors"
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
	GetAllIncomingBooking() ([]entity.Booking, error)
	InsertBooking(booking entity.Booking) error
	GetOverlapBookingCount(booking entity.Booking) int
	GetIncomingBookingByUsername(username string) ([]entity.Booking, error)
	GetBookingById(bookingId int) (entity.Booking, error)
	UpdateBookingStatus(bookingId int, status string) error
}

func (s *BookingRepository) GetAllIncomingBooking() ([]entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.room_id, b.room_name as room_name, a.booking_description, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text " +
		"from booking a, room b, lookup c where a.room_id = b.room_id and a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY' and a.start_date > current_timestamp order by a.end_date")

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.RoomId, &booking.RoomName, &booking.Description, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) InsertBooking(booking entity.Booking) error {
	_, err := s.DB.Exec("INSERT INTO booking (title,room_id,booking_description,start_date,end_date,participant_count,status, activity_code,created_by) VALUES (?,?,?,?,?,?,?,?,?)",
		booking.Title, booking.RoomId, booking.Description, booking.StartDate, booking.EndDate, booking.ParticipantCount, booking.Status, booking.Activity, booking.CreatedBy)

	if err != nil {
		log.Println("Error inserting booking: ", err.Error())
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

func (s *BookingRepository) GetIncomingBookingByUsername(username string) ([]entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.room_id, b.room_name as room_name, a.booking_description, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text "+
		"from booking a, room b, lookup c where a.room_id = b.room_id and a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY' "+
		"and a.created_by = ? and a.start_date > current_timestamp order by a.end_date", username)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.RoomId, &booking.RoomName, &booking.Description, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) GetBookingById(bookingId int) (entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.room_id, b.room_name as room_name, a.booking_description, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text "+
		"from booking a, room b, lookup c where a.room_id = b.room_id and a.booking_id = ? and a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY'", bookingId)

	if err != nil {
		log.Println("Error query booking: ", err.Error())
		return entity.Booking{}, err
	}
	defer result.Close()

	var booking entity.Booking

	if result.Next() {
		err2 := result.Scan(&booking.BookingId, &booking.Title, &booking.RoomId, &booking.RoomName, &booking.Description, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText)
		if err2 != nil {
			log.Println("Error scanning booking: ", err2.Error())
			return entity.Booking{}, err2
		}
		return booking, nil
	}
	return entity.Booking{}, errors.New("invalid bookingId")
}

func (s *BookingRepository) UpdateBookingStatus(bookingId int, status string) error {
	_, err := s.DB.Exec("UPDATE booking set status = ? where booking_id = ?",
		status, bookingId)

	if err != nil {
		log.Println("Error updating booking status: ", err.Error())
		return err
	}

	return nil

}
