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
	result, err := s.DB.Query("select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno " +
		"from booking a, lookup c where a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY' and a.start_date > current_timestamp order by a.end_date")

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) InsertBooking(booking entity.Booking) error {
	result, err := s.DB.Exec("INSERT INTO booking (title,start_date,end_date,participant_count,status,activity_code,organizer,pic,pic_contactno,created_by) VALUES (?,?,?,?,?,?,?,?,?,?)",
		booking.Title, booking.StartDate, booking.EndDate, booking.ParticipantCount, booking.Status, booking.Activity, booking.Organizer, booking.Pic, booking.PicContactNo, booking.CreatedBy)

	if err != nil {
		log.Println("Error inserting booking: ", err.Error())
		return err
	}

	// Insert booking room
	bookingId, _ := result.LastInsertId()
	for _, room := range booking.Rooms {
		_, err := s.DB.Exec("insert into booking_room(booking_id, room_id) values(?,?)", bookingId, room.RoomId)
		if err != nil {
			log.Println("Error inserting booking_room: ", err.Error())
			return err
		}
	}

	return nil
}

func (s *BookingRepository) GetOverlapBookingCount(booking entity.Booking) int {
	total := 0
	for _, room := range booking.Rooms {
		var count int
		err := s.DB.QueryRow("select count(1) from booking a, booking_room b where a.booking_id = b.booking_id and b.room_id = ? and start_date < ? and end_date > ?", room.RoomId, booking.EndDate, booking.StartDate).Scan(&count)
		if err != nil {
			return 0
		}
		total += count
	}
	return total
}

func (s *BookingRepository) GetIncomingBookingByUsername(username string) ([]entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno "+
		"from booking a, lookup c where a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY' "+
		"and a.created_by = ? and a.start_date > current_timestamp order by a.end_date", username)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var bookings []entity.Booking

	for result.Next() {
		var booking entity.Booking
		err := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo)
		if err != nil {
			log.Println("Error scanning booking: ", err.Error())
			return nil, err
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *BookingRepository) GetBookingById(bookingId int) (entity.Booking, error) {
	result, err := s.DB.Query("select a.booking_id, a.title, a.start_date, a.end_date, a.participant_count, a.status, a.activity_code, c.lookup_text as activity_text, a.organizer, a.pic, a.pic_contactno "+
		"from booking a, lookup c where a.booking_id = ? and a.activity_code = c.lookup_code and c.lookup_type='ACTIVITY'", bookingId)

	if err != nil {
		log.Println("Error query booking: ", err.Error())
		return entity.Booking{}, err
	}
	defer result.Close()

	var booking entity.Booking

	if result.Next() {
		err2 := result.Scan(&booking.BookingId, &booking.Title, &booking.StartDate, &booking.EndDate, &booking.ParticipantCount, &booking.Status, &booking.Activity, &booking.ActivityText, &booking.Organizer, &booking.Pic, &booking.PicContactNo)
		if err2 != nil {
			log.Println("Error scanning booking: ", err2.Error())
			return entity.Booking{}, err2
		}
		rooms, _ := getRoomByBookingId(s.DB, booking.BookingId)
		booking.Rooms = rooms
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

func getRoomByBookingId(DB *sql.DB, bookingId int) ([]entity.Room, error) {
	result, err := DB.Query("select b.room_id, b.room_name, b.room_description from booking_room a, room b where a.booking_id = ? and a.room_id = b.room_id", bookingId)

	if err != nil {
		return nil, err
	}
	defer result.Close()

	var rooms []entity.Room

	for result.Next() {
		var room entity.Room
		err := result.Scan(&room.RoomId, &room.Name, &room.Description)
		if err != nil {
			log.Println("Error scanning room: ", err.Error())
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
